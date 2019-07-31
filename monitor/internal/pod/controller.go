package podmonitor

import (
	"context"
	errs "errors"
	"time"

	"k8s.io/client-go/tools/record"

	"go.aporeto.io/trireme-lib/common"
	"go.aporeto.io/trireme-lib/monitor/config"
	"go.aporeto.io/trireme-lib/monitor/extractors"
	"go.aporeto.io/trireme-lib/policy"
	"go.uber.org/zap"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var (
	// ErrHandlePUStartEventFailed is the error sent back if a start event fails
	ErrHandlePUStartEventFailed = errs.New("Aporeto Enforcer start event failed")

	// ErrNetnsExtractionMissing is the error when we are missing a PID or netns path after successful metadata extraction
	ErrNetnsExtractionMissing = errs.New("Aporeto Enforcer missed to extract PID or netns path")

	// ErrHandlePUStopEventFailed is the error sent back if a stop event fails
	ErrHandlePUStopEventFailed = errs.New("Aporeto Enforcer stop event failed")

	// ErrHandlePUDestroyEventFailed is the error sent back if a create event fails
	ErrHandlePUDestroyEventFailed = errs.New("Aporeto Enforcer destroy event failed")
)

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, handler *config.ProcessorConfig, metadataExtractor extractors.PodMetadataExtractor, netclsProgrammer extractors.PodNetclsProgrammer, nodeName string, enableHostPods bool, deleteCh chan<- DeleteEvent, deleteReconcileCh chan<- struct{}) *ReconcilePod {
	return &ReconcilePod{
		client:            mgr.GetClient(),
		scheme:            mgr.GetScheme(),
		recorder:          mgr.GetRecorder("trireme-pod-controller"),
		handler:           handler,
		metadataExtractor: metadataExtractor,
		netclsProgrammer:  netclsProgrammer,
		nodeName:          nodeName,
		enableHostPods:    enableHostPods,
		deleteCh:          deleteCh,
		deleteReconcileCh: deleteReconcileCh,

		// TODO: should move into configuration
		handlePUEventTimeout:   60 * time.Second,
		metadataExtractTimeout: 10 * time.Second,
		netclsProgramTimeout:   10 * time.Second,
	}
}

// addController adds a new Controller to mgr with r as the reconcile.Reconciler
func addController(mgr manager.Manager, r *ReconcilePod, eventsCh <-chan event.GenericEvent) error {
	// Create a new controller
	c, err := controller.New("trireme-pod-controller", mgr, controller.Options{
		Reconciler: r,
		// TODO: should move into configuration
		MaxConcurrentReconciles: 4,
	})
	if err != nil {
		return err
	}

	// we use this mapper in both of our event sources
	mapper := &WatchPodMapper{
		client:         mgr.GetClient(),
		nodeName:       r.nodeName,
		enableHostPods: r.enableHostPods,
	}

	// use the our watch pod mapper which filters pods before we reconcile
	if err := c.Watch(
		&source.Kind{Type: &corev1.Pod{}},
		&handler.EnqueueRequestsFromMapFunc{ToRequests: mapper},
	); err != nil {
		return err
	}

	// we pass in a custom channel for events generated by resync
	return c.Watch(
		&source.Channel{Source: eventsCh},
		&handler.EnqueueRequestsFromMapFunc{ToRequests: mapper},
	)
}

var _ reconcile.Reconciler = &ReconcilePod{}

// DeleteEvent is used to send delete events to our event loop which will watch
// them for real deletion in the Kubernetes API. Once an object is gone, we will
// send down destroy events to trireme.
type DeleteEvent struct {
	NativeID string
	Key      client.ObjectKey
}

// ReconcilePod reconciles a Pod object
type ReconcilePod struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client            client.Client
	scheme            *runtime.Scheme
	recorder          record.EventRecorder
	handler           *config.ProcessorConfig
	metadataExtractor extractors.PodMetadataExtractor
	netclsProgrammer  extractors.PodNetclsProgrammer
	nodeName          string
	enableHostPods    bool
	deleteCh          chan<- DeleteEvent
	deleteReconcileCh chan<- struct{}

	metadataExtractTimeout time.Duration
	handlePUEventTimeout   time.Duration
	netclsProgramTimeout   time.Duration
}

// Reconcile reads that state of the cluster for a pod object
func (r *ReconcilePod) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	ctx := context.Background()
	nn := request.NamespacedName.String()

	// Fetch the corresponding pod object.
	pod := &corev1.Pod{}
	if err := r.client.Get(ctx, request.NamespacedName, pod); err != nil {
		if errors.IsNotFound(err) {
			r.deleteReconcileCh <- struct{}{}
			return reconcile.Result{}, nil
		}
		// Otherwise, we retry.
		return reconcile.Result{}, err
	}

	puID := string(pod.GetUID())

	// abort immediately if this is a HostNetwork pod, but we don't want to activate them
	// NOTE: is already done in the mapper, however, this additional check does not hurt
	if pod.Spec.HostNetwork && !r.enableHostPods {
		zap.L().Debug("Pod is a HostNetwork pod, but enableHostPods is false", zap.String("puID", puID), zap.String("namespacedName", nn))
		return reconcile.Result{}, nil
	}

	// it looks like we can miss events for all sorts of unknown reasons
	// if we reconcile though and the pod exists, we definitely know though
	// that it must go away at some point, so always register it with the delete controller
	r.deleteCh <- DeleteEvent{
		NativeID: puID,
		Key:      request.NamespacedName,
	}

	// try to find out if any of the containers have been started yet
	// this is static information on the pod, we don't need to care of the phase for determining that
	// NOTE: This is important because InitContainers are started during the PodPending phase which is
	//       what we need to rely on for activation as early as possible
	var started bool
	for _, status := range pod.Status.InitContainerStatuses {
		if status.State.Running != nil {
			started = true
			break
		}
	}
	if !started {
		for _, status := range pod.Status.ContainerStatuses {
			if status.State.Running != nil {
				started = true
				break
			}
		}
	}

	switch pod.Status.Phase {
	case corev1.PodPending:
		fallthrough
	case corev1.PodRunning:
		zap.L().Debug("PodPending / PodRunning", zap.String("puID", puID), zap.String("namespacedName", nn), zap.Bool("anyContainerStarted", started))

		// now try to do the metadata extraction
		extractCtx, extractCancel := context.WithTimeout(ctx, r.metadataExtractTimeout)
		defer extractCancel()
		puRuntime, err := r.metadataExtractor(extractCtx, r.client, r.scheme, pod, started)
		if err != nil {
			zap.L().Error("failed to extract metadata", zap.String("puID", puID), zap.String("namespacedName", nn), zap.Error(err))
			r.recorder.Eventf(pod, "Warning", "PUExtractMetadata", "PU '%s' failed to extract metadata: %s", puID, err.Error())
			return reconcile.Result{}, err
		}

		// now create/update the PU
		// every HandlePUEvent call gets done in this context
		handlePUCtx, handlePUCancel := context.WithTimeout(ctx, r.handlePUEventTimeout)
		defer handlePUCancel()
		if err := r.handler.Policy.HandlePUEvent(
			handlePUCtx,
			puID,
			common.EventUpdate,
			puRuntime,
		); err != nil {
			zap.L().Error("failed to handle update event", zap.String("puID", puID), zap.String("namespacedName", nn), zap.Error(err))
			r.recorder.Eventf(pod, "Warning", "PUUpdate", "failed to handle update event for PU '%s': %s", puID, err.Error())
			// return reconcile.Result{}, err
		} else {
			r.recorder.Eventf(pod, "Normal", "PUUpdate", "PU '%s' updated successfully", puID)
		}

		// NOTE: a pod that is terminating, is going to reconcile as well in the PodRunning phase,
		// however, it will have the deletion timestamp set which is an indicator for us that it is
		// shutting down. It means for us, that we don't have to start anything anymore. We can safely stop
		// the PU when the phase is PodSucceeded/PodFailed. However, we sent an update event above and included
		// some new tags from the metadata extractor.
		if pod.DeletionTimestamp != nil {
			return reconcile.Result{}, nil
		}

		if started {
			// if the metadata extractor is missing the PID or nspath, we need to try again
			// we need it for starting the PU. However, only require this if we are not in host network mode.
			// NOTE: this can happen for example if the containers are not in a running state on their own
			if !pod.Spec.HostNetwork && len(puRuntime.NSPath()) == 0 && puRuntime.Pid() == 0 {
				zap.L().Error("Kubernetes thinks a container is running, however, we failed to extract a PID or NSPath with the metadata extractor. Requeueing...", zap.String("puID", puID), zap.String("namespacedName", nn))
				r.recorder.Eventf(pod, "Warning", "PUStart", "PU '%s' failed to extract netns", puID)
				return reconcile.Result{}, ErrNetnsExtractionMissing
			}

			// now start the PU
			// every HandlePUEvent call gets done in this context
			handlePUStartCtx, handlePUStartCancel := context.WithTimeout(ctx, r.handlePUEventTimeout)
			defer handlePUStartCancel()
			if err := r.handler.Policy.HandlePUEvent(
				handlePUStartCtx,
				puID,
				common.EventStart,
				puRuntime,
			); err != nil {
				if policy.IsErrPUAlreadyActivated(err) {
					// abort early if this PU has already been activated before
					zap.L().Debug("PU has already been activated", zap.String("puID", puID), zap.String("namespacedName", nn), zap.Error(err))
				} else {
					zap.L().Error("failed to handle start event", zap.String("puID", puID), zap.String("namespacedName", nn), zap.Error(err))
					r.recorder.Eventf(pod, "Warning", "PUStart", "PU '%s' failed to start: %s", puID, err.Error())
				}
			} else {
				r.recorder.Eventf(pod, "Normal", "PUStart", "PU '%s' started successfully", puID)
			}

			// if this is a host network pod, we need to program the net_cls cgroup
			if pod.Spec.HostNetwork {
				netclsProgramCtx, netclsProgramCancel := context.WithTimeout(ctx, r.netclsProgramTimeout)
				defer netclsProgramCancel()
				if err := r.netclsProgrammer(netclsProgramCtx, pod, puRuntime); err != nil {
					if extractors.IsErrNetclsAlreadyProgrammed(err) {
						zap.L().Debug("net_cls cgroup has already been programmed previously", zap.String("puID", puID), zap.String("namespacedName", nn), zap.Error(err))
					} else if extractors.IsErrNoHostNetworkPod(err) {
						zap.L().Error("net_cls cgroup programmer told us that this is no host network pod.", zap.String("puID", puID), zap.String("namespacedName", nn), zap.Error(err))
					} else {
						zap.L().Error("failed to program net_cls cgroup of pod", zap.String("puID", puID), zap.String("namespacedName", nn), zap.Error(err))
						r.recorder.Eventf(pod, "Warning", "PUStart", "Host Network PU '%s' failed to program its net_cls cgroups: %s", puID, err.Error())
						return reconcile.Result{}, err
					}
				} else {
					zap.L().Debug("net_cls cgroup has been successfully programmed for trireme", zap.String("puID", puID), zap.String("namespacedName", nn))
					r.recorder.Eventf(pod, "Normal", "PUStart", "Host Network PU '%s' has successfully programmed its net_cls cgroups", puID)
				}
			}
		}
		return reconcile.Result{}, nil

	case corev1.PodSucceeded:
		fallthrough
	case corev1.PodFailed:
		zap.L().Debug("PodSucceeded / PodFailed", zap.String("puID", puID), zap.String("namespacedName", nn))
		// do metadata extraction regardless of them being stopped
		//
		// there is the edge case that the enforcer is starting up and we encounter the pod for the first time
		// in stopped state, so we have to do metadata extraction here as well
		extractCtx, extractCancel := context.WithTimeout(ctx, r.metadataExtractTimeout)
		defer extractCancel()
		puRuntime, err := r.metadataExtractor(extractCtx, r.client, r.scheme, pod, started)
		if err != nil {
			zap.L().Error("failed to extract metadata", zap.String("puID", puID), zap.String("namespacedName", nn), zap.Error(err))
			r.recorder.Eventf(pod, "Warning", "PUExtractMetadata", "PU '%s' failed to extract metadata: %s", puID, err.Error())
			return reconcile.Result{}, err
		}

		// every HandlePUEvent call gets done in this context
		handlePUCtx, handlePUCancel := context.WithTimeout(ctx, r.handlePUEventTimeout)
		defer handlePUCancel()
		if err := r.handler.Policy.HandlePUEvent(
			handlePUCtx,
			puID,
			common.EventStop,
			puRuntime,
		); err != nil {
			zap.L().Error("failed to handle stop event", zap.String("puID", puID), zap.String("namespacedName", nn), zap.Error(err))
			r.recorder.Eventf(pod, "Warning", "PUStop", "PU '%s' failed to stop: %s", puID, err.Error())
		} else {
			r.recorder.Eventf(pod, "Normal", "PUStop", "PU '%s' has been successfully stopped", puID)
		}

		// we don't need to reconcile
		// sending the stop event is enough
		return reconcile.Result{}, nil

	case corev1.PodUnknown:
		zap.L().Error("pod is in unknown state", zap.String("puID", puID), zap.String("namespacedName", nn))

		// we don't need to retry, there is nothing *we* can do about it to fix this
		return reconcile.Result{}, nil
	default:
		zap.L().Error("unknown pod phase", zap.String("puID", puID), zap.String("namespacedName", nn), zap.String("podPhase", string(pod.Status.Phase)))

		// we don't need to retry, there is nothing *we* can do about it to fix this
		return reconcile.Result{}, nil
	}
}

// Code generated by MockGen. DO NOT EDIT.
// Source: controller/internal/enforcer/enforcer.go

package mockenforcer

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	constants "go.aporeto.io/trireme-lib/v11/controller/constants"
	ebpf "go.aporeto.io/trireme-lib/v11/controller/pkg/ebpf"
	fqconfig "go.aporeto.io/trireme-lib/v11/controller/pkg/fqconfig"
	packettracing "go.aporeto.io/trireme-lib/v11/controller/pkg/packettracing"
	secrets "go.aporeto.io/trireme-lib/v11/controller/pkg/secrets"
	runtime "go.aporeto.io/trireme-lib/v11/controller/runtime"
	policy "go.aporeto.io/trireme-lib/v11/policy"
)

// MockEnforcer is a mock of Enforcer interface
// nolint
type MockEnforcer struct {
	ctrl     *gomock.Controller
	recorder *MockEnforcerMockRecorder
}

// MockEnforcerMockRecorder is the mock recorder for MockEnforcer
// nolint
type MockEnforcerMockRecorder struct {
	mock *MockEnforcer
}

// NewMockEnforcer creates a new mock instance
// nolint
func NewMockEnforcer(ctrl *gomock.Controller) *MockEnforcer {
	mock := &MockEnforcer{ctrl: ctrl}
	mock.recorder = &MockEnforcerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockEnforcer) EXPECT() *MockEnforcerMockRecorder {
	return m.recorder
}

// Enforce mocks base method
// nolint
func (m *MockEnforcer) Enforce(contextID string, puInfo *policy.PUInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enforce", contextID, puInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enforce indicates an expected call of Enforce
// nolint
func (mr *MockEnforcerMockRecorder) Enforce(contextID, puInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enforce", reflect.TypeOf((*MockEnforcer)(nil).Enforce), contextID, puInfo)
}

// Unenforce mocks base method
// nolint
func (m *MockEnforcer) Unenforce(contextID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unenforce", contextID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unenforce indicates an expected call of Unenforce
// nolint
func (mr *MockEnforcerMockRecorder) Unenforce(contextID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unenforce", reflect.TypeOf((*MockEnforcer)(nil).Unenforce), contextID)
}

// GetFilterQueue mocks base method
// nolint
func (m *MockEnforcer) GetFilterQueue() *fqconfig.FilterQueue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilterQueue")
	ret0, _ := ret[0].(*fqconfig.FilterQueue)
	return ret0
}

// GetFilterQueue indicates an expected call of GetFilterQueue
// nolint
func (mr *MockEnforcerMockRecorder) GetFilterQueue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilterQueue", reflect.TypeOf((*MockEnforcer)(nil).GetFilterQueue))
}

// GetBPFObject mocks base method
// nolint
func (m *MockEnforcer) GetBPFObject() ebpf.BPFModule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBPFObject")
	ret0, _ := ret[0].(ebpf.BPFModule)
	return ret0
}

// GetBPFObject indicates an expected call of GetBPFObject
// nolint
func (mr *MockEnforcerMockRecorder) GetBPFObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBPFObject", reflect.TypeOf((*MockEnforcer)(nil).GetBPFObject))
}

// Run mocks base method
// nolint
func (m *MockEnforcer) Run(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
// nolint
func (mr *MockEnforcerMockRecorder) Run(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockEnforcer)(nil).Run), ctx)
}

// UpdateSecrets mocks base method
// nolint
func (m *MockEnforcer) UpdateSecrets(secrets secrets.Secrets) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecrets", secrets)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecrets indicates an expected call of UpdateSecrets
// nolint
func (mr *MockEnforcerMockRecorder) UpdateSecrets(secrets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecrets", reflect.TypeOf((*MockEnforcer)(nil).UpdateSecrets), secrets)
}

// SetTargetNetworks mocks base method
// nolint
func (m *MockEnforcer) SetTargetNetworks(cfg *runtime.Configuration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTargetNetworks", cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTargetNetworks indicates an expected call of SetTargetNetworks
// nolint
func (mr *MockEnforcerMockRecorder) SetTargetNetworks(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTargetNetworks", reflect.TypeOf((*MockEnforcer)(nil).SetTargetNetworks), cfg)
}

// SetLogLevel mocks base method
// nolint
func (m *MockEnforcer) SetLogLevel(level constants.LogLevel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLogLevel", level)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLogLevel indicates an expected call of SetLogLevel
// nolint
func (mr *MockEnforcerMockRecorder) SetLogLevel(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogLevel", reflect.TypeOf((*MockEnforcer)(nil).SetLogLevel), level)
}

// CleanUp mocks base method
// nolint
func (m *MockEnforcer) CleanUp() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanUp")
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUp indicates an expected call of CleanUp
// nolint
func (mr *MockEnforcerMockRecorder) CleanUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUp", reflect.TypeOf((*MockEnforcer)(nil).CleanUp))
}

// EnableDatapathPacketTracing mocks base method
// nolint
func (m *MockEnforcer) EnableDatapathPacketTracing(ctx context.Context, contextID string, direction packettracing.TracingDirection, interval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableDatapathPacketTracing", ctx, contextID, direction, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableDatapathPacketTracing indicates an expected call of EnableDatapathPacketTracing
// nolint
func (mr *MockEnforcerMockRecorder) EnableDatapathPacketTracing(ctx, contextID, direction, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDatapathPacketTracing", reflect.TypeOf((*MockEnforcer)(nil).EnableDatapathPacketTracing), ctx, contextID, direction, interval)
}

// EnableIPTablesPacketTracing mocks base method
// nolint
func (m *MockEnforcer) EnableIPTablesPacketTracing(ctx context.Context, contextID string, interval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableIPTablesPacketTracing", ctx, contextID, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableIPTablesPacketTracing indicates an expected call of EnableIPTablesPacketTracing
// nolint
func (mr *MockEnforcerMockRecorder) EnableIPTablesPacketTracing(ctx, contextID, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableIPTablesPacketTracing", reflect.TypeOf((*MockEnforcer)(nil).EnableIPTablesPacketTracing), ctx, contextID, interval)
}

// Ping mocks base method
// nolint
func (m *MockEnforcer) Ping(ctx context.Context, contextID string, pingConfig *policy.PingConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx, contextID, pingConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
// nolint
func (mr *MockEnforcerMockRecorder) Ping(ctx, contextID, pingConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockEnforcer)(nil).Ping), ctx, contextID, pingConfig)
}

// MockDebugInfo is a mock of DebugInfo interface
// nolint
type MockDebugInfo struct {
	ctrl     *gomock.Controller
	recorder *MockDebugInfoMockRecorder
}

// MockDebugInfoMockRecorder is the mock recorder for MockDebugInfo
// nolint
type MockDebugInfoMockRecorder struct {
	mock *MockDebugInfo
}

// NewMockDebugInfo creates a new mock instance
// nolint
func NewMockDebugInfo(ctrl *gomock.Controller) *MockDebugInfo {
	mock := &MockDebugInfo{ctrl: ctrl}
	mock.recorder = &MockDebugInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (_m *MockDebugInfo) EXPECT() *MockDebugInfoMockRecorder {
	return _m.recorder
}

// EnableDatapathPacketTracing mocks base method
// nolint
func (m *MockDebugInfo) EnableDatapathPacketTracing(ctx context.Context, contextID string, direction packettracing.TracingDirection, interval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableDatapathPacketTracing", ctx, contextID, direction, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableDatapathPacketTracing indicates an expected call of EnableDatapathPacketTracing
// nolint
func (mr *MockDebugInfoMockRecorder) EnableDatapathPacketTracing(ctx, contextID, direction, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDatapathPacketTracing", reflect.TypeOf((*MockDebugInfo)(nil).EnableDatapathPacketTracing), ctx, contextID, direction, interval)
}

// EnableIPTablesPacketTracing mocks base method
// nolint
func (m *MockDebugInfo) EnableIPTablesPacketTracing(ctx context.Context, contextID string, interval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableIPTablesPacketTracing", ctx, contextID, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableIPTablesPacketTracing indicates an expected call of EnableIPTablesPacketTracing
// nolint
func (mr *MockDebugInfoMockRecorder) EnableIPTablesPacketTracing(ctx, contextID, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableIPTablesPacketTracing", reflect.TypeOf((*MockDebugInfo)(nil).EnableIPTablesPacketTracing), ctx, contextID, interval)
}

// Ping mocks base method
// nolint
func (m *MockDebugInfo) Ping(ctx context.Context, contextID string, pingConfig *policy.PingConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx, contextID, pingConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
// nolint
func (mr *MockDebugInfoMockRecorder) Ping(ctx, contextID, pingConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDebugInfo)(nil).Ping), ctx, contextID, pingConfig)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: controller/interfaces.go

package mockcontroller

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	packettracing "go.aporeto.io/trireme-lib/v11/controller/pkg/packettracing"
	secrets "go.aporeto.io/trireme-lib/v11/controller/pkg/secrets"
	runtime "go.aporeto.io/trireme-lib/v11/controller/runtime"
	policy "go.aporeto.io/trireme-lib/v11/policy"
)

// MockTriremeController is a mock of TriremeController interface
// nolint
type MockTriremeController struct {
	ctrl     *gomock.Controller
	recorder *MockTriremeControllerMockRecorder
}

// MockTriremeControllerMockRecorder is the mock recorder for MockTriremeController
// nolint
type MockTriremeControllerMockRecorder struct {
	mock *MockTriremeController
}

// NewMockTriremeController creates a new mock instance
// nolint
func NewMockTriremeController(ctrl *gomock.Controller) *MockTriremeController {
	mock := &MockTriremeController{ctrl: ctrl}
	mock.recorder = &MockTriremeControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
// nolint
func (m *MockTriremeController) EXPECT() *MockTriremeControllerMockRecorder {
	return m.recorder
}

// Run mocks base method
// nolint
func (m *MockTriremeController) Run(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
// nolint
func (mr *MockTriremeControllerMockRecorder) Run(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockTriremeController)(nil).Run), ctx)
}

// CleanUp mocks base method
// nolint
func (m *MockTriremeController) CleanUp() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanUp")
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUp indicates an expected call of CleanUp
// nolint
func (mr *MockTriremeControllerMockRecorder) CleanUp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUp", reflect.TypeOf((*MockTriremeController)(nil).CleanUp))
}

// Enforce mocks base method
// nolint
func (m *MockTriremeController) Enforce(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enforce", ctx, puID, policy, runtime)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enforce indicates an expected call of Enforce
// nolint
func (mr *MockTriremeControllerMockRecorder) Enforce(ctx, puID, policy, runtime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enforce", reflect.TypeOf((*MockTriremeController)(nil).Enforce), ctx, puID, policy, runtime)
}

// UnEnforce mocks base method
// nolint
func (m *MockTriremeController) UnEnforce(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnEnforce", ctx, puID, policy, runtime)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnEnforce indicates an expected call of UnEnforce
// nolint
func (mr *MockTriremeControllerMockRecorder) UnEnforce(ctx, puID, policy, runtime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnEnforce", reflect.TypeOf((*MockTriremeController)(nil).UnEnforce), ctx, puID, policy, runtime)
}

// UpdatePolicy mocks base method
// nolint
func (m *MockTriremeController) UpdatePolicy(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePolicy", ctx, puID, policy, runtime)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePolicy indicates an expected call of UpdatePolicy
// nolint
func (mr *MockTriremeControllerMockRecorder) UpdatePolicy(ctx, puID, policy, runtime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePolicy", reflect.TypeOf((*MockTriremeController)(nil).UpdatePolicy), ctx, puID, policy, runtime)
}

// UpdateSecrets mocks base method
// nolint
func (m *MockTriremeController) UpdateSecrets(secrets secrets.Secrets) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecrets", secrets)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecrets indicates an expected call of UpdateSecrets
// nolint
func (mr *MockTriremeControllerMockRecorder) UpdateSecrets(secrets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecrets", reflect.TypeOf((*MockTriremeController)(nil).UpdateSecrets), secrets)
}

// UpdateConfiguration mocks base method
// nolint
func (m *MockTriremeController) UpdateConfiguration(cfg *runtime.Configuration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfiguration", cfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfiguration indicates an expected call of UpdateConfiguration
// nolint
func (mr *MockTriremeControllerMockRecorder) UpdateConfiguration(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfiguration", reflect.TypeOf((*MockTriremeController)(nil).UpdateConfiguration), cfg)
}

// EnableDatapathPacketTracing mocks base method
// nolint
func (m *MockTriremeController) EnableDatapathPacketTracing(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, direction packettracing.TracingDirection, interval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableDatapathPacketTracing", ctx, puID, policy, runtime, direction, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableDatapathPacketTracing indicates an expected call of EnableDatapathPacketTracing
// nolint
func (mr *MockTriremeControllerMockRecorder) EnableDatapathPacketTracing(ctx, puID, policy, runtime, direction, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDatapathPacketTracing", reflect.TypeOf((*MockTriremeController)(nil).EnableDatapathPacketTracing), ctx, puID, policy, runtime, direction, interval)
}

// EnableIPTablesPacketTracing mocks base method
// nolint
func (m *MockTriremeController) EnableIPTablesPacketTracing(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, interval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableIPTablesPacketTracing", ctx, puID, policy, runtime, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableIPTablesPacketTracing indicates an expected call of EnableIPTablesPacketTracing
// nolint
func (mr *MockTriremeControllerMockRecorder) EnableIPTablesPacketTracing(ctx, puID, policy, runtime, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableIPTablesPacketTracing", reflect.TypeOf((*MockTriremeController)(nil).EnableIPTablesPacketTracing), ctx, puID, policy, runtime, interval)
}

// Ping mocks base method
// nolint
func (m *MockTriremeController) Ping(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, pingConfig *policy.PingConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx, puID, policy, runtime, pingConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
// nolint
func (mr *MockTriremeControllerMockRecorder) Ping(ctx, puID, policy, runtime, pingConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockTriremeController)(nil).Ping), ctx, puID, policy, runtime, pingConfig)
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
func (m *MockDebugInfo) EnableDatapathPacketTracing(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, direction packettracing.TracingDirection, interval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableDatapathPacketTracing", ctx, puID, policy, runtime, direction, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableDatapathPacketTracing indicates an expected call of EnableDatapathPacketTracing
// nolint
func (mr *MockDebugInfoMockRecorder) EnableDatapathPacketTracing(ctx, puID, policy, runtime, direction, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableDatapathPacketTracing", reflect.TypeOf((*MockDebugInfo)(nil).EnableDatapathPacketTracing), ctx, puID, policy, runtime, direction, interval)
}

// EnableIPTablesPacketTracing mocks base method
// nolint
func (m *MockDebugInfo) EnableIPTablesPacketTracing(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, interval time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableIPTablesPacketTracing", ctx, puID, policy, runtime, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableIPTablesPacketTracing indicates an expected call of EnableIPTablesPacketTracing
// nolint
func (mr *MockDebugInfoMockRecorder) EnableIPTablesPacketTracing(ctx, puID, policy, runtime, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableIPTablesPacketTracing", reflect.TypeOf((*MockDebugInfo)(nil).EnableIPTablesPacketTracing), ctx, puID, policy, runtime, interval)
}

// Ping mocks base method
// nolint
func (m *MockDebugInfo) Ping(ctx context.Context, puID string, policy *policy.PUPolicy, runtime *policy.PURuntime, pingConfig *policy.PingConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx, puID, policy, runtime, pingConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
// nolint
func (mr *MockDebugInfoMockRecorder) Ping(ctx, puID, policy, runtime, pingConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDebugInfo)(nil).Ping), ctx, puID, policy, runtime, pingConfig)
}

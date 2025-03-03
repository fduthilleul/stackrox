// Code generated by MockGen. DO NOT EDIT.
// Source: types.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/types.go -source types.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	views "github.com/stackrox/rox/central/views"
	common "github.com/stackrox/rox/central/views/common"
	imagecve "github.com/stackrox/rox/central/views/imagecve"
	v1 "github.com/stackrox/rox/generated/api/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockCveCore is a mock of CveCore interface.
type MockCveCore struct {
	ctrl     *gomock.Controller
	recorder *MockCveCoreMockRecorder
	isgomock struct{}
}

// MockCveCoreMockRecorder is the mock recorder for MockCveCore.
type MockCveCoreMockRecorder struct {
	mock *MockCveCore
}

// NewMockCveCore creates a new mock instance.
func NewMockCveCore(ctrl *gomock.Controller) *MockCveCore {
	mock := &MockCveCore{ctrl: ctrl}
	mock.recorder = &MockCveCoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCveCore) EXPECT() *MockCveCoreMockRecorder {
	return m.recorder
}

// GetAffectedImageCount mocks base method.
func (m *MockCveCore) GetAffectedImageCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAffectedImageCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetAffectedImageCount indicates an expected call of GetAffectedImageCount.
func (mr *MockCveCoreMockRecorder) GetAffectedImageCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAffectedImageCount", reflect.TypeOf((*MockCveCore)(nil).GetAffectedImageCount))
}

// GetCVE mocks base method.
func (m *MockCveCore) GetCVE() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCVE")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCVE indicates an expected call of GetCVE.
func (mr *MockCveCoreMockRecorder) GetCVE() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCVE", reflect.TypeOf((*MockCveCore)(nil).GetCVE))
}

// GetCVEIDs mocks base method.
func (m *MockCveCore) GetCVEIDs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCVEIDs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetCVEIDs indicates an expected call of GetCVEIDs.
func (mr *MockCveCoreMockRecorder) GetCVEIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCVEIDs", reflect.TypeOf((*MockCveCore)(nil).GetCVEIDs))
}

// GetFirstDiscoveredInSystem mocks base method.
func (m *MockCveCore) GetFirstDiscoveredInSystem() *time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstDiscoveredInSystem")
	ret0, _ := ret[0].(*time.Time)
	return ret0
}

// GetFirstDiscoveredInSystem indicates an expected call of GetFirstDiscoveredInSystem.
func (mr *MockCveCoreMockRecorder) GetFirstDiscoveredInSystem() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstDiscoveredInSystem", reflect.TypeOf((*MockCveCore)(nil).GetFirstDiscoveredInSystem))
}

// GetImagesBySeverity mocks base method.
func (m *MockCveCore) GetImagesBySeverity() common.ResourceCountByCVESeverity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImagesBySeverity")
	ret0, _ := ret[0].(common.ResourceCountByCVESeverity)
	return ret0
}

// GetImagesBySeverity indicates an expected call of GetImagesBySeverity.
func (mr *MockCveCoreMockRecorder) GetImagesBySeverity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImagesBySeverity", reflect.TypeOf((*MockCveCore)(nil).GetImagesBySeverity))
}

// GetPublishDate mocks base method.
func (m *MockCveCore) GetPublishDate() *time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublishDate")
	ret0, _ := ret[0].(*time.Time)
	return ret0
}

// GetPublishDate indicates an expected call of GetPublishDate.
func (mr *MockCveCoreMockRecorder) GetPublishDate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublishDate", reflect.TypeOf((*MockCveCore)(nil).GetPublishDate))
}

// GetTopCVSS mocks base method.
func (m *MockCveCore) GetTopCVSS() float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopCVSS")
	ret0, _ := ret[0].(float32)
	return ret0
}

// GetTopCVSS indicates an expected call of GetTopCVSS.
func (mr *MockCveCoreMockRecorder) GetTopCVSS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopCVSS", reflect.TypeOf((*MockCveCore)(nil).GetTopCVSS))
}

// GetTopNVDCVSS mocks base method.
func (m *MockCveCore) GetTopNVDCVSS() float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopNVDCVSS")
	ret0, _ := ret[0].(float32)
	return ret0
}

// GetTopNVDCVSS indicates an expected call of GetTopNVDCVSS.
func (mr *MockCveCoreMockRecorder) GetTopNVDCVSS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopNVDCVSS", reflect.TypeOf((*MockCveCore)(nil).GetTopNVDCVSS))
}

// MockCveView is a mock of CveView interface.
type MockCveView struct {
	ctrl     *gomock.Controller
	recorder *MockCveViewMockRecorder
	isgomock struct{}
}

// MockCveViewMockRecorder is the mock recorder for MockCveView.
type MockCveViewMockRecorder struct {
	mock *MockCveView
}

// NewMockCveView creates a new mock instance.
func NewMockCveView(ctrl *gomock.Controller) *MockCveView {
	mock := &MockCveView{ctrl: ctrl}
	mock.recorder = &MockCveViewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCveView) EXPECT() *MockCveViewMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockCveView) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockCveViewMockRecorder) Count(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCveView)(nil).Count), ctx, q)
}

// CountBySeverity mocks base method.
func (m *MockCveView) CountBySeverity(ctx context.Context, q *v1.Query) (common.ResourceCountByCVESeverity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountBySeverity", ctx, q)
	ret0, _ := ret[0].(common.ResourceCountByCVESeverity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountBySeverity indicates an expected call of CountBySeverity.
func (mr *MockCveViewMockRecorder) CountBySeverity(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountBySeverity", reflect.TypeOf((*MockCveView)(nil).CountBySeverity), ctx, q)
}

// Get mocks base method.
func (m *MockCveView) Get(ctx context.Context, q *v1.Query, options views.ReadOptions) ([]imagecve.CveCore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, q, options)
	ret0, _ := ret[0].([]imagecve.CveCore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCveViewMockRecorder) Get(ctx, q, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCveView)(nil).Get), ctx, q, options)
}

// GetDeploymentIDs mocks base method.
func (m *MockCveView) GetDeploymentIDs(ctx context.Context, q *v1.Query) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentIDs", ctx, q)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentIDs indicates an expected call of GetDeploymentIDs.
func (mr *MockCveViewMockRecorder) GetDeploymentIDs(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentIDs", reflect.TypeOf((*MockCveView)(nil).GetDeploymentIDs), ctx, q)
}

// GetImageIDs mocks base method.
func (m *MockCveView) GetImageIDs(ctx context.Context, q *v1.Query) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageIDs", ctx, q)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageIDs indicates an expected call of GetImageIDs.
func (mr *MockCveViewMockRecorder) GetImageIDs(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageIDs", reflect.TypeOf((*MockCveView)(nil).GetImageIDs), ctx, q)
}

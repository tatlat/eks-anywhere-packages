// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/bundle/manager.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere-packages/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// DownloadBundle mocks base method.
func (m *MockManager) DownloadBundle(ctx context.Context, ref string) (*v1alpha1.PackageBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadBundle", ctx, ref)
	ret0, _ := ret[0].(*v1alpha1.PackageBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadBundle indicates an expected call of DownloadBundle.
func (mr *MockManagerMockRecorder) DownloadBundle(ctx, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadBundle", reflect.TypeOf((*MockManager)(nil).DownloadBundle), ctx, ref)
}

// LatestBundle mocks base method.
func (m *MockManager) LatestBundle(ctx context.Context, baseRef string) (*v1alpha1.PackageBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestBundle", ctx, baseRef)
	ret0, _ := ret[0].(*v1alpha1.PackageBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestBundle indicates an expected call of LatestBundle.
func (mr *MockManagerMockRecorder) LatestBundle(ctx, baseRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestBundle", reflect.TypeOf((*MockManager)(nil).LatestBundle), ctx, baseRef)
}

// ProcessLatestBundle mocks base method.
func (m *MockManager) ProcessLatestBundle(ctx context.Context, bundle *v1alpha1.PackageBundle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessLatestBundle", ctx, bundle)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessLatestBundle indicates an expected call of ProcessLatestBundle.
func (mr *MockManagerMockRecorder) ProcessLatestBundle(ctx, bundle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessLatestBundle", reflect.TypeOf((*MockManager)(nil).ProcessLatestBundle), ctx, bundle)
}

// SortBundlesDescending mocks base method.
func (m *MockManager) SortBundlesDescending(bundles []v1alpha1.PackageBundle) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SortBundlesDescending", bundles)
}

// SortBundlesDescending indicates an expected call of SortBundlesDescending.
func (mr *MockManagerMockRecorder) SortBundlesDescending(bundles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SortBundlesDescending", reflect.TypeOf((*MockManager)(nil).SortBundlesDescending), bundles)
}

// Update mocks base method.
func (m *MockManager) ProcessBundle(ctx context.Context, newBundle *v1alpha1.PackageBundle) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessBundle", ctx, newBundle)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockManagerMockRecorder) Update(ctx, newBundle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessBundle", reflect.TypeOf((*MockManager)(nil).ProcessBundle), ctx, newBundle)
}

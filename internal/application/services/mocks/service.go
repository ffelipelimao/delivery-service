// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	domain "github.com/ffelipelimao/delivery-service/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockService) Create(ctx context.Context, register domain.Register) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, register)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(ctx, register interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), ctx, register)
}

// Get mocks base method.
func (m *MockService) Get(ctx context.Context, ID string) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, ID)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServiceMockRecorder) Get(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockService)(nil).Get), ctx, ID)
}

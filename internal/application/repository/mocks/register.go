// Code generated by MockGen. DO NOT EDIT.
// Source: ./register.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	domain "github.com/ffelipelimao/delivery-service/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockRegisterRepository is a mock of RegisterRepository interface.
type MockRegisterRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterRepositoryMockRecorder
}

// MockRegisterRepositoryMockRecorder is the mock recorder for MockRegisterRepository.
type MockRegisterRepositoryMockRecorder struct {
	mock *MockRegisterRepository
}

// NewMockRegisterRepository creates a new mock instance.
func NewMockRegisterRepository(ctrl *gomock.Controller) *MockRegisterRepository {
	mock := &MockRegisterRepository{ctrl: ctrl}
	mock.recorder = &MockRegisterRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegisterRepository) EXPECT() *MockRegisterRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRegisterRepository) Get(ID string) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ID)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRegisterRepositoryMockRecorder) Get(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRegisterRepository)(nil).Get), ID)
}

// Insert mocks base method.
func (m *MockRegisterRepository) Insert(register domain.Register) (domain.Register, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", register)
	ret0, _ := ret[0].(domain.Register)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockRegisterRepositoryMockRecorder) Insert(register interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockRegisterRepository)(nil).Insert), register)
}

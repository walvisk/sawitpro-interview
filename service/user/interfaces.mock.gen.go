// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/user/interfaces.go

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	generated "github.com/SawitProRecruitment/UserService/generated"
	repository "github.com/SawitProRecruitment/UserService/repository"
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

// FindUserByID mocks base method.
func (m *MockService) FindUserByID(c context.Context, id int64) (*repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", c, id)
	ret0, _ := ret[0].(*repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockServiceMockRecorder) FindUserByID(c, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockService)(nil).FindUserByID), c, id)
}

// FindUserByPhone mocks base method.
func (m *MockService) FindUserByPhone(c context.Context, phone string) (*repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByPhone", c, phone)
	ret0, _ := ret[0].(*repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByPhone indicates an expected call of FindUserByPhone.
func (mr *MockServiceMockRecorder) FindUserByPhone(c, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByPhone", reflect.TypeOf((*MockService)(nil).FindUserByPhone), c, phone)
}

// RegisterUser mocks base method.
func (m *MockService) RegisterUser(c context.Context, params generated.CreateUserJSONRequestBody) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", c, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockServiceMockRecorder) RegisterUser(c, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockService)(nil).RegisterUser), c, params)
}

// UpdateUser mocks base method.
func (m *MockService) UpdateUser(c context.Context, u *repository.User, phone, fullName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", c, u, phone, fullName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockServiceMockRecorder) UpdateUser(c, u, phone, fullName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockService)(nil).UpdateUser), c, u, phone, fullName)
}

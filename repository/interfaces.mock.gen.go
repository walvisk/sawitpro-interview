// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockRepositoryInterface) CreateUser(c context.Context, u *User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", c, u)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryInterfaceMockRecorder) CreateUser(c, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateUser), c, u)
}

// CreateUserLog mocks base method.
func (m *MockRepositoryInterface) CreateUserLog(c context.Context, u *User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserLog", c, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserLog indicates an expected call of CreateUserLog.
func (mr *MockRepositoryInterfaceMockRecorder) CreateUserLog(c, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserLog", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateUserLog), c, u)
}

// FindUserByID mocks base method.
func (m *MockRepositoryInterface) FindUserByID(c context.Context, id int64) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", c, id)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID.
func (mr *MockRepositoryInterfaceMockRecorder) FindUserByID(c, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockRepositoryInterface)(nil).FindUserByID), c, id)
}

// FindUserByPhoneAndCountryCode mocks base method.
func (m *MockRepositoryInterface) FindUserByPhoneAndCountryCode(c context.Context, phone, countryCode string) (*User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByPhoneAndCountryCode", c, phone, countryCode)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByPhoneAndCountryCode indicates an expected call of FindUserByPhoneAndCountryCode.
func (mr *MockRepositoryInterfaceMockRecorder) FindUserByPhoneAndCountryCode(c, phone, countryCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByPhoneAndCountryCode", reflect.TypeOf((*MockRepositoryInterface)(nil).FindUserByPhoneAndCountryCode), c, phone, countryCode)
}

// UpdateUser mocks base method.
func (m *MockRepositoryInterface) UpdateUser(c context.Context, u *User, fullName, phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", c, u, fullName, phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateUser(c, u, fullName, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateUser), c, u, fullName, phone)
}

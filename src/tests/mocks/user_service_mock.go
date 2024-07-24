// Code generated by MockGen. DO NOT EDIT.
// Source: src/model/service/user_interface.go
//
// Generated by this command:
//
//	mockgen --source=src/model/service/user_interface.go --destination=src/test/mocks/user_service_mock.go --package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	rest_err "go-with-docker-and-swagger/src/configuration/rest_err"
	model "go-with-docker-and-swagger/src/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserDomainService is a mock of UserDomainService interface.
type MockUserDomainService struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainServiceMockRecorder
}

// MockUserDomainServiceMockRecorder is the mock recorder for MockUserDomainService.
type MockUserDomainServiceMockRecorder struct {
	mock *MockUserDomainService
}

// NewMockUserDomainService creates a new mock instance.
func NewMockUserDomainService(ctrl *gomock.Controller) *MockUserDomainService {
	mock := &MockUserDomainService{ctrl: ctrl}
	mock.recorder = &MockUserDomainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDomainService) EXPECT() *MockUserDomainServiceMockRecorder {
	return m.recorder
}

// CreateUserServices mocks base method.
func (m *MockUserDomainService) CreateUserServices(arg0 model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserServices", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// CreateUserServices indicates an expected call of CreateUserServices.
func (mr *MockUserDomainServiceMockRecorder) CreateUserServices(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserServices", reflect.TypeOf((*MockUserDomainService)(nil).CreateUserServices), arg0)
}

// DeleteUserServices mocks base method.
func (m *MockUserDomainService) DeleteUserServices(arg0 string) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserServices", arg0)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// DeleteUserServices indicates an expected call of DeleteUserServices.
func (mr *MockUserDomainServiceMockRecorder) DeleteUserServices(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserServices", reflect.TypeOf((*MockUserDomainService)(nil).DeleteUserServices), arg0)
}

// FindUserByEmailServices mocks base method.
func (m *MockUserDomainService) FindUserByEmailServices(arg0 string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmailServices", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserByEmailServices indicates an expected call of FindUserByEmailServices.
func (mr *MockUserDomainServiceMockRecorder) FindUserByEmailServices(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmailServices", reflect.TypeOf((*MockUserDomainService)(nil).FindUserByEmailServices), arg0)
}

// FindUserByIDServices mocks base method.
func (m *MockUserDomainService) FindUserByIDServices(arg0 string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByIDServices", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserByIDServices indicates an expected call of FindUserByIDServices.
func (mr *MockUserDomainServiceMockRecorder) FindUserByIDServices(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByIDServices", reflect.TypeOf((*MockUserDomainService)(nil).FindUserByIDServices), arg0)
}

// LoginUserServices mocks base method.
func (m *MockUserDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUserServices", userDomain)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(*rest_err.RestErr)
	return ret0, ret1, ret2
}

// LoginUserServices indicates an expected call of LoginUserServices.
func (mr *MockUserDomainServiceMockRecorder) LoginUserServices(userDomain any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUserServices", reflect.TypeOf((*MockUserDomainService)(nil).LoginUserServices), userDomain)
}

// UpdateUserServices mocks base method.
func (m *MockUserDomainService) UpdateUserServices(arg0 string, arg1 model.UserDomainInterface) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserServices", arg0, arg1)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// UpdateUserServices indicates an expected call of UpdateUserServices.
func (mr *MockUserDomainServiceMockRecorder) UpdateUserServices(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserServices", reflect.TypeOf((*MockUserDomainService)(nil).UpdateUserServices), arg0, arg1)
}

// findUserByEmailAndPasswordServices mocks base method.
func (m *MockUserDomainService) findUserByEmailAndPasswordServices(arg0, arg1 string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "findUserByEmailAndPasswordServices", arg0, arg1)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// findUserByEmailAndPasswordServices indicates an expected call of findUserByEmailAndPasswordServices.
func (mr *MockUserDomainServiceMockRecorder) findUserByEmailAndPasswordServices(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "findUserByEmailAndPasswordServices", reflect.TypeOf((*MockUserDomainService)(nil).findUserByEmailAndPasswordServices), arg0, arg1)
}
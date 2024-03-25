package mocks

import (
	"reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/response"
	"github.com/minand-mohan/library-app-api/api/users/dto"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserHandler.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserHandler creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserService) CreateUser(arg0 *dto.UserRequestBody) (*response.HTTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*response.HTTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), arg0)
}

func (m *MockUserService) FindAllUsers(arg0 *dto.UserQueryParams) (*response.HTTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllUsers", arg0)
	ret0, _ := ret[0].(*response.HTTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserServiceMockRecorder) FindAllUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllUsers", reflect.TypeOf((*MockUserService)(nil).FindAllUsers), arg0)
}

func (m *MockUserService) FindByUserId(arg0 uuid.UUID) (*response.HTTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserId", arg0)
	ret0, _ := ret[0].(*response.HTTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserServiceMockRecorder) FindByUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserId", reflect.TypeOf((*MockUserService)(nil).FindByUserId), arg0)
}

func (m *MockUserService) UpdateByUserId(arg0 uuid.UUID, arg1 *dto.UserRequestBody) (*response.HTTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByUserId", arg0, arg1)
	ret0, _ := ret[0].(*response.HTTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserServiceMockRecorder) UpdateByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByUserId", reflect.TypeOf((*MockUserService)(nil).UpdateByUserId), arg0, arg1)
}

func (m *MockUserService) DeleteByUserId(arg0 uuid.UUID) (*response.HTTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUserId", arg0)
	ret0, _ := ret[0].(*response.HTTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserServiceMockRecorder) DeleteByUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUserId", reflect.TypeOf((*MockUserService)(nil).DeleteByUserId), arg0)
}

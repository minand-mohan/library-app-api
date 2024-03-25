package validator_tests

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/minand-mohan/library-app-api/api/users/dto"
)

// MockUserValidator is a mock type for the UserValidator type
type MockUserValidator struct {
	ctrl     *gomock.Controller
	recorder *MockUserValidatorMockRecorder
}

type MockUserValidatorMockRecorder struct {
	mock *MockUserValidator
}

func NewMockUserValidator(ctrl *gomock.Controller) *MockUserValidator {
	mock := &MockUserValidator{ctrl: ctrl}
	mock.recorder = &MockUserValidatorMockRecorder{mock}
	return mock
}

func (m *MockUserValidator) EXPECT() *MockUserValidatorMockRecorder {
	return m.recorder
}

func (m *MockUserValidator) ValidateUser(arg0 *dto.UserRequestBody) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUserValidatorMockRecorder) ValidateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUser", reflect.TypeOf((*MockUserValidator)(nil).ValidateUser), arg0)
}

func (m *MockUserValidator) ValidateUserQueryParams(arg0 *dto.UserQueryParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUserQueryParams", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUserValidatorMockRecorder) ValidateUserQueryParams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUserQueryParams", reflect.TypeOf((*MockUserValidator)(nil).ValidateUserQueryParams), arg0)
}

package mocks

import (
	"reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/minand-mohan/library-app-api/api/users/dto"
	"github.com/minand-mohan/library-app-api/database/models"
)

// MockUserRepository is a mock of UserService interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

func (m *MockUserRepository) CreateUser(arg0 *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUserRepositoryMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), arg0)
}

func (m *MockUserRepository) FindByEmailOrUsernameOrPhone(arg0 string, arg1 string, arg2 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmailOrUsernameOrPhone", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockRecorder) FindByEmailOrUsernameOrPhone(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmailOrUsernameOrPhone", reflect.TypeOf((*MockUserRepository)(nil).FindByEmailOrUsernameOrPhone), arg0, arg1, arg2)
}

func (m *MockUserRepository) FindAllUsers(arg0 *dto.UserQueryParams) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllUsers", arg0)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockRecorder) FindAllUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllUsers", reflect.TypeOf((*MockUserRepository)(nil).FindAllUsers), arg0)
}

func (m *MockUserRepository) FindByUserId(arg0 uuid.UUID) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserId", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockRecorder) FindByUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserId", reflect.TypeOf((*MockUserRepository)(nil).FindByUserId), arg0)
}

func (m *MockUserRepository) UpdateByUserId(arg0 uuid.UUID, arg1 *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByUserId", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserRepositoryMockRecorder) UpdateByUserId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByUserId", reflect.TypeOf((*MockUserRepository)(nil).UpdateByUserId), arg0, arg1)
}

func (m *MockUserRepository) DeleteByUserId(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUserId", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUserRepositoryMockRecorder) DeleteByUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUserId", reflect.TypeOf((*MockUserRepository)(nil).DeleteByUserId), arg0)
}

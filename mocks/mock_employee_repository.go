// Code generated by MockGen. DO NOT EDIT.
// Source: store/store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/itsmaheshkariya/gin-gorm-rest/models"
)

// MockUserRepository is a mock of UserRepository interface.
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

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), user)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), id)
}

// GetUserByID mocks base method.
func (m *MockUserRepository) GetUserByID(id string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepositoryMockRecorder) GetUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepository)(nil).GetUserByID), id)
}

// GetUsers mocks base method.
func (m *MockUserRepository) GetUsers() []models.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]models.User)
	return ret0
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserRepositoryMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserRepository)(nil).GetUsers))
}

// UpdateUser mocks base method.
func (m *MockUserRepository) UpdateUser(id string, user *models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", id, user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), id, user)
}

// MockEmployeeRepository is a mock of EmployeeRepository interface.
type MockEmployeeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeRepositoryMockRecorder
}

// MockEmployeeRepositoryMockRecorder is the mock recorder for MockEmployeeRepository.
type MockEmployeeRepositoryMockRecorder struct {
	mock *MockEmployeeRepository
}

// NewMockEmployeeRepository creates a new mock instance.
func NewMockEmployeeRepository(ctrl *gomock.Controller) *MockEmployeeRepository {
	mock := &MockEmployeeRepository{ctrl: ctrl}
	mock.recorder = &MockEmployeeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeRepository) EXPECT() *MockEmployeeRepositoryMockRecorder {
	return m.recorder
}

// CreateEmployee mocks base method.
func (m *MockEmployeeRepository) CreateEmployee(employee *models.Employee) (*models.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployee", employee)
	ret0, _ := ret[0].(*models.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmployee indicates an expected call of CreateEmployee.
func (mr *MockEmployeeRepositoryMockRecorder) CreateEmployee(employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployee", reflect.TypeOf((*MockEmployeeRepository)(nil).CreateEmployee), employee)
}

// DeleteEmployee mocks base method.
func (m *MockEmployeeRepository) DeleteEmployee(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmployee", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmployee indicates an expected call of DeleteEmployee.
func (mr *MockEmployeeRepositoryMockRecorder) DeleteEmployee(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmployee", reflect.TypeOf((*MockEmployeeRepository)(nil).DeleteEmployee), id)
}

// GetEmployeeByID mocks base method.
func (m *MockEmployeeRepository) GetEmployeeByID(id string) (*models.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeByID", id)
	ret0, _ := ret[0].(*models.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployeeByID indicates an expected call of GetEmployeeByID.
func (mr *MockEmployeeRepositoryMockRecorder) GetEmployeeByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeByID", reflect.TypeOf((*MockEmployeeRepository)(nil).GetEmployeeByID), id)
}

// GetEmployees mocks base method.
func (m *MockEmployeeRepository) GetEmployees(page, pageSize int) ([]models.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployees", page, pageSize)
	ret0, _ := ret[0].([]models.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployees indicates an expected call of GetEmployees.
func (mr *MockEmployeeRepositoryMockRecorder) GetEmployees(page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployees", reflect.TypeOf((*MockEmployeeRepository)(nil).GetEmployees), page, pageSize)
}

// UpdateEmployee mocks base method.
func (m *MockEmployeeRepository) UpdateEmployee(id string, employee *models.Employee) (*models.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmployee", id, employee)
	ret0, _ := ret[0].(*models.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEmployee indicates an expected call of UpdateEmployee.
func (mr *MockEmployeeRepositoryMockRecorder) UpdateEmployee(id, employee interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmployee", reflect.TypeOf((*MockEmployeeRepository)(nil).UpdateEmployee), id, employee)
}

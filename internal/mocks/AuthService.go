// Code generated by mockery v2.28.2. DO NOT EDIT.

package mocks

import (
	dto "github.com/saadozone/gin-gorm-rest/internal/dto"
	mock "github.com/stretchr/testify/mock"

	model "github.com/saadozone/gin-gorm-rest/internal/model"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// Attempt provides a mock function with given fields: input
func (_m *AuthService) Attempt(input *dto.LoginRequestBody) (*model.User, error) {
	ret := _m.Called(input)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*dto.LoginRequestBody) (*model.User, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(*dto.LoginRequestBody) *model.User); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*dto.LoginRequestBody) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForgotPass provides a mock function with given fields: input
func (_m *AuthService) ForgotPass(input *dto.ForgotPasswordRequestBody) (*model.PasswordReset, error) {
	ret := _m.Called(input)

	var r0 *model.PasswordReset
	var r1 error
	if rf, ok := ret.Get(0).(func(*dto.ForgotPasswordRequestBody) (*model.PasswordReset, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(*dto.ForgotPasswordRequestBody) *model.PasswordReset); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PasswordReset)
		}
	}

	if rf, ok := ret.Get(1).(func(*dto.ForgotPasswordRequestBody) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetPass provides a mock function with given fields: input
func (_m *AuthService) ResetPass(input *dto.ResetPasswordRequestBody) (*model.PasswordReset, error) {
	ret := _m.Called(input)

	var r0 *model.PasswordReset
	var r1 error
	if rf, ok := ret.Get(0).(func(*dto.ResetPasswordRequestBody) (*model.PasswordReset, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(*dto.ResetPasswordRequestBody) *model.PasswordReset); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PasswordReset)
		}
	}

	if rf, ok := ret.Get(1).(func(*dto.ResetPasswordRequestBody) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthService interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthService(t mockConstructorTestingTNewAuthService) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
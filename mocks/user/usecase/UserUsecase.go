// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import usecase "github.com/stereoit/eventival/pkg/user/usecase"

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// FindByID provides a mock function with given fields: id
func (_m *UserUsecase) FindByID(id string) (*usecase.User, error) {
	ret := _m.Called(id)

	var r0 *usecase.User
	if rf, ok := ret.Get(0).(func(string) *usecase.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usecase.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllUsers provides a mock function with given fields:
func (_m *UserUsecase) ListAllUsers() ([]*usecase.User, error) {
	ret := _m.Called()

	var r0 []*usecase.User
	if rf, ok := ret.Get(0).(func() []*usecase.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*usecase.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: email
func (_m *UserUsecase) RegisterUser(email string) (string, error) {
	ret := _m.Called(email)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0
func (_m *UserUsecase) UpdateUser(_a0 *usecase.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*usecase.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
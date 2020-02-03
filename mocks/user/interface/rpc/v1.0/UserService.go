// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import protocol "github.com/stereoit/eventival/pkg/user/interface/rpc/v1.0/protocol"

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// ListUser provides a mock function with given fields: ctx, in
func (_m *UserService) ListUser(ctx context.Context, in *protocol.ListUserRequestType) (*protocol.ListUserResponseType, error) {
	ret := _m.Called(ctx, in)

	var r0 *protocol.ListUserResponseType
	if rf, ok := ret.Get(0).(func(context.Context, *protocol.ListUserRequestType) *protocol.ListUserResponseType); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protocol.ListUserResponseType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protocol.ListUserRequestType) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, in
func (_m *UserService) RegisterUser(ctx context.Context, in *protocol.RegisterUserRequestType) (*protocol.RegisterUserResponseType, error) {
	ret := _m.Called(ctx, in)

	var r0 *protocol.RegisterUserResponseType
	if rf, ok := ret.Get(0).(func(context.Context, *protocol.RegisterUserRequestType) *protocol.RegisterUserResponseType); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protocol.RegisterUserResponseType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protocol.RegisterUserRequestType) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
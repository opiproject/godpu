/* SPDX-License-Identifier: Apache-2.0
   Copyright (c) 2023 Dell Inc, or its subsidiaries.
*/
// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	_go "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// SviServiceClient is an autogenerated mock type for the SviServiceClient type
type SviServiceClient struct {
	mock.Mock
}

type SviServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *SviServiceClient) EXPECT() *SviServiceClient_Expecter {
	return &SviServiceClient_Expecter{mock: &_m.Mock}
}

// CreateSvi provides a mock function with given fields: ctx, in, opts
func (_m *SviServiceClient) CreateSvi(ctx context.Context, in *_go.CreateSviRequest, opts ...grpc.CallOption) (*_go.Svi, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.Svi
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.CreateSviRequest, ...grpc.CallOption) (*_go.Svi, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.CreateSviRequest, ...grpc.CallOption) *_go.Svi); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.Svi)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.CreateSviRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SviServiceClient_CreateSvi_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSvi'
type SviServiceClient_CreateSvi_Call struct {
	*mock.Call
}

// CreateSvi is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.CreateSviRequest
//   - opts ...grpc.CallOption
func (_e *SviServiceClient_Expecter) CreateSvi(ctx interface{}, in interface{}, opts ...interface{}) *SviServiceClient_CreateSvi_Call {
	return &SviServiceClient_CreateSvi_Call{Call: _e.mock.On("CreateSvi",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *SviServiceClient_CreateSvi_Call) Run(run func(ctx context.Context, in *_go.CreateSviRequest, opts ...grpc.CallOption)) *SviServiceClient_CreateSvi_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.CreateSviRequest), variadicArgs...)
	})
	return _c
}

func (_c *SviServiceClient_CreateSvi_Call) Return(_a0 *_go.Svi, _a1 error) *SviServiceClient_CreateSvi_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SviServiceClient_CreateSvi_Call) RunAndReturn(run func(context.Context, *_go.CreateSviRequest, ...grpc.CallOption) (*_go.Svi, error)) *SviServiceClient_CreateSvi_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSvi provides a mock function with given fields: ctx, in, opts
func (_m *SviServiceClient) DeleteSvi(ctx context.Context, in *_go.DeleteSviRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *emptypb.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.DeleteSviRequest, ...grpc.CallOption) (*emptypb.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.DeleteSviRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.DeleteSviRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SviServiceClient_DeleteSvi_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSvi'
type SviServiceClient_DeleteSvi_Call struct {
	*mock.Call
}

// DeleteSvi is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.DeleteSviRequest
//   - opts ...grpc.CallOption
func (_e *SviServiceClient_Expecter) DeleteSvi(ctx interface{}, in interface{}, opts ...interface{}) *SviServiceClient_DeleteSvi_Call {
	return &SviServiceClient_DeleteSvi_Call{Call: _e.mock.On("DeleteSvi",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *SviServiceClient_DeleteSvi_Call) Run(run func(ctx context.Context, in *_go.DeleteSviRequest, opts ...grpc.CallOption)) *SviServiceClient_DeleteSvi_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.DeleteSviRequest), variadicArgs...)
	})
	return _c
}

func (_c *SviServiceClient_DeleteSvi_Call) Return(_a0 *emptypb.Empty, _a1 error) *SviServiceClient_DeleteSvi_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SviServiceClient_DeleteSvi_Call) RunAndReturn(run func(context.Context, *_go.DeleteSviRequest, ...grpc.CallOption) (*emptypb.Empty, error)) *SviServiceClient_DeleteSvi_Call {
	_c.Call.Return(run)
	return _c
}

// GetSvi provides a mock function with given fields: ctx, in, opts
func (_m *SviServiceClient) GetSvi(ctx context.Context, in *_go.GetSviRequest, opts ...grpc.CallOption) (*_go.Svi, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.Svi
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.GetSviRequest, ...grpc.CallOption) (*_go.Svi, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.GetSviRequest, ...grpc.CallOption) *_go.Svi); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.Svi)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.GetSviRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SviServiceClient_GetSvi_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSvi'
type SviServiceClient_GetSvi_Call struct {
	*mock.Call
}

// GetSvi is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.GetSviRequest
//   - opts ...grpc.CallOption
func (_e *SviServiceClient_Expecter) GetSvi(ctx interface{}, in interface{}, opts ...interface{}) *SviServiceClient_GetSvi_Call {
	return &SviServiceClient_GetSvi_Call{Call: _e.mock.On("GetSvi",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *SviServiceClient_GetSvi_Call) Run(run func(ctx context.Context, in *_go.GetSviRequest, opts ...grpc.CallOption)) *SviServiceClient_GetSvi_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.GetSviRequest), variadicArgs...)
	})
	return _c
}

func (_c *SviServiceClient_GetSvi_Call) Return(_a0 *_go.Svi, _a1 error) *SviServiceClient_GetSvi_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SviServiceClient_GetSvi_Call) RunAndReturn(run func(context.Context, *_go.GetSviRequest, ...grpc.CallOption) (*_go.Svi, error)) *SviServiceClient_GetSvi_Call {
	_c.Call.Return(run)
	return _c
}

// ListSvis provides a mock function with given fields: ctx, in, opts
func (_m *SviServiceClient) ListSvis(ctx context.Context, in *_go.ListSvisRequest, opts ...grpc.CallOption) (*_go.ListSvisResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.ListSvisResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.ListSvisRequest, ...grpc.CallOption) (*_go.ListSvisResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.ListSvisRequest, ...grpc.CallOption) *_go.ListSvisResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.ListSvisResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.ListSvisRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SviServiceClient_ListSvis_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSvis'
type SviServiceClient_ListSvis_Call struct {
	*mock.Call
}

// ListSvis is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.ListSvisRequest
//   - opts ...grpc.CallOption
func (_e *SviServiceClient_Expecter) ListSvis(ctx interface{}, in interface{}, opts ...interface{}) *SviServiceClient_ListSvis_Call {
	return &SviServiceClient_ListSvis_Call{Call: _e.mock.On("ListSvis",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *SviServiceClient_ListSvis_Call) Run(run func(ctx context.Context, in *_go.ListSvisRequest, opts ...grpc.CallOption)) *SviServiceClient_ListSvis_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.ListSvisRequest), variadicArgs...)
	})
	return _c
}

func (_c *SviServiceClient_ListSvis_Call) Return(_a0 *_go.ListSvisResponse, _a1 error) *SviServiceClient_ListSvis_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SviServiceClient_ListSvis_Call) RunAndReturn(run func(context.Context, *_go.ListSvisRequest, ...grpc.CallOption) (*_go.ListSvisResponse, error)) *SviServiceClient_ListSvis_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSvi provides a mock function with given fields: ctx, in, opts
func (_m *SviServiceClient) UpdateSvi(ctx context.Context, in *_go.UpdateSviRequest, opts ...grpc.CallOption) (*_go.Svi, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.Svi
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.UpdateSviRequest, ...grpc.CallOption) (*_go.Svi, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.UpdateSviRequest, ...grpc.CallOption) *_go.Svi); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.Svi)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.UpdateSviRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SviServiceClient_UpdateSvi_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSvi'
type SviServiceClient_UpdateSvi_Call struct {
	*mock.Call
}

// UpdateSvi is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.UpdateSviRequest
//   - opts ...grpc.CallOption
func (_e *SviServiceClient_Expecter) UpdateSvi(ctx interface{}, in interface{}, opts ...interface{}) *SviServiceClient_UpdateSvi_Call {
	return &SviServiceClient_UpdateSvi_Call{Call: _e.mock.On("UpdateSvi",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *SviServiceClient_UpdateSvi_Call) Run(run func(ctx context.Context, in *_go.UpdateSviRequest, opts ...grpc.CallOption)) *SviServiceClient_UpdateSvi_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.UpdateSviRequest), variadicArgs...)
	})
	return _c
}

func (_c *SviServiceClient_UpdateSvi_Call) Return(_a0 *_go.Svi, _a1 error) *SviServiceClient_UpdateSvi_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SviServiceClient_UpdateSvi_Call) RunAndReturn(run func(context.Context, *_go.UpdateSviRequest, ...grpc.CallOption) (*_go.Svi, error)) *SviServiceClient_UpdateSvi_Call {
	_c.Call.Return(run)
	return _c
}

// NewSviServiceClient creates a new instance of SviServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSviServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *SviServiceClient {
	mock := &SviServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

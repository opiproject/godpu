/* SPDX-License-Identifier: Apache-2.0
   Copyright (c) 2023 Dell Inc, or its subsidiaries.
*/
// Code generated by mockery v2.33.1. DO NOT EDIT.

package mocks

import (
	context "context"

	_go "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// NullVolumeServiceClient is an autogenerated mock type for the NullVolumeServiceClient type
type NullVolumeServiceClient struct {
	mock.Mock
}

type NullVolumeServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *NullVolumeServiceClient) EXPECT() *NullVolumeServiceClient_Expecter {
	return &NullVolumeServiceClient_Expecter{mock: &_m.Mock}
}

// CreateNullVolume provides a mock function with given fields: ctx, in, opts
func (_m *NullVolumeServiceClient) CreateNullVolume(ctx context.Context, in *_go.CreateNullVolumeRequest, opts ...grpc.CallOption) (*_go.NullVolume, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.NullVolume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.CreateNullVolumeRequest, ...grpc.CallOption) (*_go.NullVolume, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.CreateNullVolumeRequest, ...grpc.CallOption) *_go.NullVolume); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.NullVolume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.CreateNullVolumeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NullVolumeServiceClient_CreateNullVolume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateNullVolume'
type NullVolumeServiceClient_CreateNullVolume_Call struct {
	*mock.Call
}

// CreateNullVolume is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.CreateNullVolumeRequest
//   - opts ...grpc.CallOption
func (_e *NullVolumeServiceClient_Expecter) CreateNullVolume(ctx interface{}, in interface{}, opts ...interface{}) *NullVolumeServiceClient_CreateNullVolume_Call {
	return &NullVolumeServiceClient_CreateNullVolume_Call{Call: _e.mock.On("CreateNullVolume",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *NullVolumeServiceClient_CreateNullVolume_Call) Run(run func(ctx context.Context, in *_go.CreateNullVolumeRequest, opts ...grpc.CallOption)) *NullVolumeServiceClient_CreateNullVolume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.CreateNullVolumeRequest), variadicArgs...)
	})
	return _c
}

func (_c *NullVolumeServiceClient_CreateNullVolume_Call) Return(_a0 *_go.NullVolume, _a1 error) *NullVolumeServiceClient_CreateNullVolume_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NullVolumeServiceClient_CreateNullVolume_Call) RunAndReturn(run func(context.Context, *_go.CreateNullVolumeRequest, ...grpc.CallOption) (*_go.NullVolume, error)) *NullVolumeServiceClient_CreateNullVolume_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteNullVolume provides a mock function with given fields: ctx, in, opts
func (_m *NullVolumeServiceClient) DeleteNullVolume(ctx context.Context, in *_go.DeleteNullVolumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
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
	if rf, ok := ret.Get(0).(func(context.Context, *_go.DeleteNullVolumeRequest, ...grpc.CallOption) (*emptypb.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.DeleteNullVolumeRequest, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.DeleteNullVolumeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NullVolumeServiceClient_DeleteNullVolume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteNullVolume'
type NullVolumeServiceClient_DeleteNullVolume_Call struct {
	*mock.Call
}

// DeleteNullVolume is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.DeleteNullVolumeRequest
//   - opts ...grpc.CallOption
func (_e *NullVolumeServiceClient_Expecter) DeleteNullVolume(ctx interface{}, in interface{}, opts ...interface{}) *NullVolumeServiceClient_DeleteNullVolume_Call {
	return &NullVolumeServiceClient_DeleteNullVolume_Call{Call: _e.mock.On("DeleteNullVolume",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *NullVolumeServiceClient_DeleteNullVolume_Call) Run(run func(ctx context.Context, in *_go.DeleteNullVolumeRequest, opts ...grpc.CallOption)) *NullVolumeServiceClient_DeleteNullVolume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.DeleteNullVolumeRequest), variadicArgs...)
	})
	return _c
}

func (_c *NullVolumeServiceClient_DeleteNullVolume_Call) Return(_a0 *emptypb.Empty, _a1 error) *NullVolumeServiceClient_DeleteNullVolume_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NullVolumeServiceClient_DeleteNullVolume_Call) RunAndReturn(run func(context.Context, *_go.DeleteNullVolumeRequest, ...grpc.CallOption) (*emptypb.Empty, error)) *NullVolumeServiceClient_DeleteNullVolume_Call {
	_c.Call.Return(run)
	return _c
}

// GetNullVolume provides a mock function with given fields: ctx, in, opts
func (_m *NullVolumeServiceClient) GetNullVolume(ctx context.Context, in *_go.GetNullVolumeRequest, opts ...grpc.CallOption) (*_go.NullVolume, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.NullVolume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.GetNullVolumeRequest, ...grpc.CallOption) (*_go.NullVolume, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.GetNullVolumeRequest, ...grpc.CallOption) *_go.NullVolume); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.NullVolume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.GetNullVolumeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NullVolumeServiceClient_GetNullVolume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNullVolume'
type NullVolumeServiceClient_GetNullVolume_Call struct {
	*mock.Call
}

// GetNullVolume is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.GetNullVolumeRequest
//   - opts ...grpc.CallOption
func (_e *NullVolumeServiceClient_Expecter) GetNullVolume(ctx interface{}, in interface{}, opts ...interface{}) *NullVolumeServiceClient_GetNullVolume_Call {
	return &NullVolumeServiceClient_GetNullVolume_Call{Call: _e.mock.On("GetNullVolume",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *NullVolumeServiceClient_GetNullVolume_Call) Run(run func(ctx context.Context, in *_go.GetNullVolumeRequest, opts ...grpc.CallOption)) *NullVolumeServiceClient_GetNullVolume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.GetNullVolumeRequest), variadicArgs...)
	})
	return _c
}

func (_c *NullVolumeServiceClient_GetNullVolume_Call) Return(_a0 *_go.NullVolume, _a1 error) *NullVolumeServiceClient_GetNullVolume_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NullVolumeServiceClient_GetNullVolume_Call) RunAndReturn(run func(context.Context, *_go.GetNullVolumeRequest, ...grpc.CallOption) (*_go.NullVolume, error)) *NullVolumeServiceClient_GetNullVolume_Call {
	_c.Call.Return(run)
	return _c
}

// ListNullVolumes provides a mock function with given fields: ctx, in, opts
func (_m *NullVolumeServiceClient) ListNullVolumes(ctx context.Context, in *_go.ListNullVolumesRequest, opts ...grpc.CallOption) (*_go.ListNullVolumesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.ListNullVolumesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.ListNullVolumesRequest, ...grpc.CallOption) (*_go.ListNullVolumesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.ListNullVolumesRequest, ...grpc.CallOption) *_go.ListNullVolumesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.ListNullVolumesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.ListNullVolumesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NullVolumeServiceClient_ListNullVolumes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListNullVolumes'
type NullVolumeServiceClient_ListNullVolumes_Call struct {
	*mock.Call
}

// ListNullVolumes is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.ListNullVolumesRequest
//   - opts ...grpc.CallOption
func (_e *NullVolumeServiceClient_Expecter) ListNullVolumes(ctx interface{}, in interface{}, opts ...interface{}) *NullVolumeServiceClient_ListNullVolumes_Call {
	return &NullVolumeServiceClient_ListNullVolumes_Call{Call: _e.mock.On("ListNullVolumes",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *NullVolumeServiceClient_ListNullVolumes_Call) Run(run func(ctx context.Context, in *_go.ListNullVolumesRequest, opts ...grpc.CallOption)) *NullVolumeServiceClient_ListNullVolumes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.ListNullVolumesRequest), variadicArgs...)
	})
	return _c
}

func (_c *NullVolumeServiceClient_ListNullVolumes_Call) Return(_a0 *_go.ListNullVolumesResponse, _a1 error) *NullVolumeServiceClient_ListNullVolumes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NullVolumeServiceClient_ListNullVolumes_Call) RunAndReturn(run func(context.Context, *_go.ListNullVolumesRequest, ...grpc.CallOption) (*_go.ListNullVolumesResponse, error)) *NullVolumeServiceClient_ListNullVolumes_Call {
	_c.Call.Return(run)
	return _c
}

// StatsNullVolume provides a mock function with given fields: ctx, in, opts
func (_m *NullVolumeServiceClient) StatsNullVolume(ctx context.Context, in *_go.StatsNullVolumeRequest, opts ...grpc.CallOption) (*_go.StatsNullVolumeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.StatsNullVolumeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.StatsNullVolumeRequest, ...grpc.CallOption) (*_go.StatsNullVolumeResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.StatsNullVolumeRequest, ...grpc.CallOption) *_go.StatsNullVolumeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.StatsNullVolumeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.StatsNullVolumeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NullVolumeServiceClient_StatsNullVolume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StatsNullVolume'
type NullVolumeServiceClient_StatsNullVolume_Call struct {
	*mock.Call
}

// StatsNullVolume is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.StatsNullVolumeRequest
//   - opts ...grpc.CallOption
func (_e *NullVolumeServiceClient_Expecter) StatsNullVolume(ctx interface{}, in interface{}, opts ...interface{}) *NullVolumeServiceClient_StatsNullVolume_Call {
	return &NullVolumeServiceClient_StatsNullVolume_Call{Call: _e.mock.On("StatsNullVolume",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *NullVolumeServiceClient_StatsNullVolume_Call) Run(run func(ctx context.Context, in *_go.StatsNullVolumeRequest, opts ...grpc.CallOption)) *NullVolumeServiceClient_StatsNullVolume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.StatsNullVolumeRequest), variadicArgs...)
	})
	return _c
}

func (_c *NullVolumeServiceClient_StatsNullVolume_Call) Return(_a0 *_go.StatsNullVolumeResponse, _a1 error) *NullVolumeServiceClient_StatsNullVolume_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NullVolumeServiceClient_StatsNullVolume_Call) RunAndReturn(run func(context.Context, *_go.StatsNullVolumeRequest, ...grpc.CallOption) (*_go.StatsNullVolumeResponse, error)) *NullVolumeServiceClient_StatsNullVolume_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateNullVolume provides a mock function with given fields: ctx, in, opts
func (_m *NullVolumeServiceClient) UpdateNullVolume(ctx context.Context, in *_go.UpdateNullVolumeRequest, opts ...grpc.CallOption) (*_go.NullVolume, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.NullVolume
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.UpdateNullVolumeRequest, ...grpc.CallOption) (*_go.NullVolume, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.UpdateNullVolumeRequest, ...grpc.CallOption) *_go.NullVolume); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.NullVolume)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.UpdateNullVolumeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NullVolumeServiceClient_UpdateNullVolume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateNullVolume'
type NullVolumeServiceClient_UpdateNullVolume_Call struct {
	*mock.Call
}

// UpdateNullVolume is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.UpdateNullVolumeRequest
//   - opts ...grpc.CallOption
func (_e *NullVolumeServiceClient_Expecter) UpdateNullVolume(ctx interface{}, in interface{}, opts ...interface{}) *NullVolumeServiceClient_UpdateNullVolume_Call {
	return &NullVolumeServiceClient_UpdateNullVolume_Call{Call: _e.mock.On("UpdateNullVolume",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *NullVolumeServiceClient_UpdateNullVolume_Call) Run(run func(ctx context.Context, in *_go.UpdateNullVolumeRequest, opts ...grpc.CallOption)) *NullVolumeServiceClient_UpdateNullVolume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.UpdateNullVolumeRequest), variadicArgs...)
	})
	return _c
}

func (_c *NullVolumeServiceClient_UpdateNullVolume_Call) Return(_a0 *_go.NullVolume, _a1 error) *NullVolumeServiceClient_UpdateNullVolume_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NullVolumeServiceClient_UpdateNullVolume_Call) RunAndReturn(run func(context.Context, *_go.UpdateNullVolumeRequest, ...grpc.CallOption) (*_go.NullVolume, error)) *NullVolumeServiceClient_UpdateNullVolume_Call {
	_c.Call.Return(run)
	return _c
}

// NewNullVolumeServiceClient creates a new instance of NullVolumeServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNullVolumeServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *NullVolumeServiceClient {
	mock := &NullVolumeServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

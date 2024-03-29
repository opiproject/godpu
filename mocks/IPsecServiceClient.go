/* SPDX-License-Identifier: Apache-2.0
   Copyright (c) 2023 Dell Inc, or its subsidiaries.
*/

// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	_go "github.com/opiproject/opi-api/security/v1/gen/go"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// IPsecServiceClient is an autogenerated mock type for the IPsecServiceClient type
type IPsecServiceClient struct {
	mock.Mock
}

type IPsecServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *IPsecServiceClient) EXPECT() *IPsecServiceClient_Expecter {
	return &IPsecServiceClient_Expecter{mock: &_m.Mock}
}

// IPsecInitiate provides a mock function with given fields: ctx, in, opts
func (_m *IPsecServiceClient) IPsecInitiate(ctx context.Context, in *_go.IPsecInitiateRequest, opts ...grpc.CallOption) (*_go.IPsecInitiateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.IPsecInitiateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecInitiateRequest, ...grpc.CallOption) (*_go.IPsecInitiateResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecInitiateRequest, ...grpc.CallOption) *_go.IPsecInitiateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.IPsecInitiateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.IPsecInitiateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPsecServiceClient_IPsecInitiate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPsecInitiate'
type IPsecServiceClient_IPsecInitiate_Call struct {
	*mock.Call
}

// IPsecInitiate is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.IPsecInitiateRequest
//   - opts ...grpc.CallOption
func (_e *IPsecServiceClient_Expecter) IPsecInitiate(ctx interface{}, in interface{}, opts ...interface{}) *IPsecServiceClient_IPsecInitiate_Call {
	return &IPsecServiceClient_IPsecInitiate_Call{Call: _e.mock.On("IPsecInitiate",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *IPsecServiceClient_IPsecInitiate_Call) Run(run func(ctx context.Context, in *_go.IPsecInitiateRequest, opts ...grpc.CallOption)) *IPsecServiceClient_IPsecInitiate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.IPsecInitiateRequest), variadicArgs...)
	})
	return _c
}

func (_c *IPsecServiceClient_IPsecInitiate_Call) Return(_a0 *_go.IPsecInitiateResponse, _a1 error) *IPsecServiceClient_IPsecInitiate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IPsecServiceClient_IPsecInitiate_Call) RunAndReturn(run func(context.Context, *_go.IPsecInitiateRequest, ...grpc.CallOption) (*_go.IPsecInitiateResponse, error)) *IPsecServiceClient_IPsecInitiate_Call {
	_c.Call.Return(run)
	return _c
}

// IPsecListCerts provides a mock function with given fields: ctx, in, opts
func (_m *IPsecServiceClient) IPsecListCerts(ctx context.Context, in *_go.IPsecListCertsRequest, opts ...grpc.CallOption) (*_go.IPsecListCertsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.IPsecListCertsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecListCertsRequest, ...grpc.CallOption) (*_go.IPsecListCertsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecListCertsRequest, ...grpc.CallOption) *_go.IPsecListCertsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.IPsecListCertsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.IPsecListCertsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPsecServiceClient_IPsecListCerts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPsecListCerts'
type IPsecServiceClient_IPsecListCerts_Call struct {
	*mock.Call
}

// IPsecListCerts is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.IPsecListCertsRequest
//   - opts ...grpc.CallOption
func (_e *IPsecServiceClient_Expecter) IPsecListCerts(ctx interface{}, in interface{}, opts ...interface{}) *IPsecServiceClient_IPsecListCerts_Call {
	return &IPsecServiceClient_IPsecListCerts_Call{Call: _e.mock.On("IPsecListCerts",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *IPsecServiceClient_IPsecListCerts_Call) Run(run func(ctx context.Context, in *_go.IPsecListCertsRequest, opts ...grpc.CallOption)) *IPsecServiceClient_IPsecListCerts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.IPsecListCertsRequest), variadicArgs...)
	})
	return _c
}

func (_c *IPsecServiceClient_IPsecListCerts_Call) Return(_a0 *_go.IPsecListCertsResponse, _a1 error) *IPsecServiceClient_IPsecListCerts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IPsecServiceClient_IPsecListCerts_Call) RunAndReturn(run func(context.Context, *_go.IPsecListCertsRequest, ...grpc.CallOption) (*_go.IPsecListCertsResponse, error)) *IPsecServiceClient_IPsecListCerts_Call {
	_c.Call.Return(run)
	return _c
}

// IPsecListConns provides a mock function with given fields: ctx, in, opts
func (_m *IPsecServiceClient) IPsecListConns(ctx context.Context, in *_go.IPsecListConnsRequest, opts ...grpc.CallOption) (*_go.IPsecListConnsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.IPsecListConnsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecListConnsRequest, ...grpc.CallOption) (*_go.IPsecListConnsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecListConnsRequest, ...grpc.CallOption) *_go.IPsecListConnsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.IPsecListConnsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.IPsecListConnsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPsecServiceClient_IPsecListConns_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPsecListConns'
type IPsecServiceClient_IPsecListConns_Call struct {
	*mock.Call
}

// IPsecListConns is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.IPsecListConnsRequest
//   - opts ...grpc.CallOption
func (_e *IPsecServiceClient_Expecter) IPsecListConns(ctx interface{}, in interface{}, opts ...interface{}) *IPsecServiceClient_IPsecListConns_Call {
	return &IPsecServiceClient_IPsecListConns_Call{Call: _e.mock.On("IPsecListConns",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *IPsecServiceClient_IPsecListConns_Call) Run(run func(ctx context.Context, in *_go.IPsecListConnsRequest, opts ...grpc.CallOption)) *IPsecServiceClient_IPsecListConns_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.IPsecListConnsRequest), variadicArgs...)
	})
	return _c
}

func (_c *IPsecServiceClient_IPsecListConns_Call) Return(_a0 *_go.IPsecListConnsResponse, _a1 error) *IPsecServiceClient_IPsecListConns_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IPsecServiceClient_IPsecListConns_Call) RunAndReturn(run func(context.Context, *_go.IPsecListConnsRequest, ...grpc.CallOption) (*_go.IPsecListConnsResponse, error)) *IPsecServiceClient_IPsecListConns_Call {
	_c.Call.Return(run)
	return _c
}

// IPsecListSas provides a mock function with given fields: ctx, in, opts
func (_m *IPsecServiceClient) IPsecListSas(ctx context.Context, in *_go.IPsecListSasRequest, opts ...grpc.CallOption) (*_go.IPsecListSasResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.IPsecListSasResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecListSasRequest, ...grpc.CallOption) (*_go.IPsecListSasResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecListSasRequest, ...grpc.CallOption) *_go.IPsecListSasResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.IPsecListSasResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.IPsecListSasRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPsecServiceClient_IPsecListSas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPsecListSas'
type IPsecServiceClient_IPsecListSas_Call struct {
	*mock.Call
}

// IPsecListSas is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.IPsecListSasRequest
//   - opts ...grpc.CallOption
func (_e *IPsecServiceClient_Expecter) IPsecListSas(ctx interface{}, in interface{}, opts ...interface{}) *IPsecServiceClient_IPsecListSas_Call {
	return &IPsecServiceClient_IPsecListSas_Call{Call: _e.mock.On("IPsecListSas",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *IPsecServiceClient_IPsecListSas_Call) Run(run func(ctx context.Context, in *_go.IPsecListSasRequest, opts ...grpc.CallOption)) *IPsecServiceClient_IPsecListSas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.IPsecListSasRequest), variadicArgs...)
	})
	return _c
}

func (_c *IPsecServiceClient_IPsecListSas_Call) Return(_a0 *_go.IPsecListSasResponse, _a1 error) *IPsecServiceClient_IPsecListSas_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IPsecServiceClient_IPsecListSas_Call) RunAndReturn(run func(context.Context, *_go.IPsecListSasRequest, ...grpc.CallOption) (*_go.IPsecListSasResponse, error)) *IPsecServiceClient_IPsecListSas_Call {
	_c.Call.Return(run)
	return _c
}

// IPsecLoadConn provides a mock function with given fields: ctx, in, opts
func (_m *IPsecServiceClient) IPsecLoadConn(ctx context.Context, in *_go.IPsecLoadConnRequest, opts ...grpc.CallOption) (*_go.IPsecLoadConnResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.IPsecLoadConnResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecLoadConnRequest, ...grpc.CallOption) (*_go.IPsecLoadConnResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecLoadConnRequest, ...grpc.CallOption) *_go.IPsecLoadConnResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.IPsecLoadConnResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.IPsecLoadConnRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPsecServiceClient_IPsecLoadConn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPsecLoadConn'
type IPsecServiceClient_IPsecLoadConn_Call struct {
	*mock.Call
}

// IPsecLoadConn is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.IPsecLoadConnRequest
//   - opts ...grpc.CallOption
func (_e *IPsecServiceClient_Expecter) IPsecLoadConn(ctx interface{}, in interface{}, opts ...interface{}) *IPsecServiceClient_IPsecLoadConn_Call {
	return &IPsecServiceClient_IPsecLoadConn_Call{Call: _e.mock.On("IPsecLoadConn",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *IPsecServiceClient_IPsecLoadConn_Call) Run(run func(ctx context.Context, in *_go.IPsecLoadConnRequest, opts ...grpc.CallOption)) *IPsecServiceClient_IPsecLoadConn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.IPsecLoadConnRequest), variadicArgs...)
	})
	return _c
}

func (_c *IPsecServiceClient_IPsecLoadConn_Call) Return(_a0 *_go.IPsecLoadConnResponse, _a1 error) *IPsecServiceClient_IPsecLoadConn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IPsecServiceClient_IPsecLoadConn_Call) RunAndReturn(run func(context.Context, *_go.IPsecLoadConnRequest, ...grpc.CallOption) (*_go.IPsecLoadConnResponse, error)) *IPsecServiceClient_IPsecLoadConn_Call {
	_c.Call.Return(run)
	return _c
}

// IPsecRekey provides a mock function with given fields: ctx, in, opts
func (_m *IPsecServiceClient) IPsecRekey(ctx context.Context, in *_go.IPsecRekeyRequest, opts ...grpc.CallOption) (*_go.IPsecRekeyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.IPsecRekeyResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecRekeyRequest, ...grpc.CallOption) (*_go.IPsecRekeyResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecRekeyRequest, ...grpc.CallOption) *_go.IPsecRekeyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.IPsecRekeyResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.IPsecRekeyRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPsecServiceClient_IPsecRekey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPsecRekey'
type IPsecServiceClient_IPsecRekey_Call struct {
	*mock.Call
}

// IPsecRekey is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.IPsecRekeyRequest
//   - opts ...grpc.CallOption
func (_e *IPsecServiceClient_Expecter) IPsecRekey(ctx interface{}, in interface{}, opts ...interface{}) *IPsecServiceClient_IPsecRekey_Call {
	return &IPsecServiceClient_IPsecRekey_Call{Call: _e.mock.On("IPsecRekey",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *IPsecServiceClient_IPsecRekey_Call) Run(run func(ctx context.Context, in *_go.IPsecRekeyRequest, opts ...grpc.CallOption)) *IPsecServiceClient_IPsecRekey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.IPsecRekeyRequest), variadicArgs...)
	})
	return _c
}

func (_c *IPsecServiceClient_IPsecRekey_Call) Return(_a0 *_go.IPsecRekeyResponse, _a1 error) *IPsecServiceClient_IPsecRekey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IPsecServiceClient_IPsecRekey_Call) RunAndReturn(run func(context.Context, *_go.IPsecRekeyRequest, ...grpc.CallOption) (*_go.IPsecRekeyResponse, error)) *IPsecServiceClient_IPsecRekey_Call {
	_c.Call.Return(run)
	return _c
}

// IPsecStats provides a mock function with given fields: ctx, in, opts
func (_m *IPsecServiceClient) IPsecStats(ctx context.Context, in *_go.IPsecStatsRequest, opts ...grpc.CallOption) (*_go.IPsecStatsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.IPsecStatsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecStatsRequest, ...grpc.CallOption) (*_go.IPsecStatsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecStatsRequest, ...grpc.CallOption) *_go.IPsecStatsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.IPsecStatsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.IPsecStatsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPsecServiceClient_IPsecStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPsecStats'
type IPsecServiceClient_IPsecStats_Call struct {
	*mock.Call
}

// IPsecStats is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.IPsecStatsRequest
//   - opts ...grpc.CallOption
func (_e *IPsecServiceClient_Expecter) IPsecStats(ctx interface{}, in interface{}, opts ...interface{}) *IPsecServiceClient_IPsecStats_Call {
	return &IPsecServiceClient_IPsecStats_Call{Call: _e.mock.On("IPsecStats",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *IPsecServiceClient_IPsecStats_Call) Run(run func(ctx context.Context, in *_go.IPsecStatsRequest, opts ...grpc.CallOption)) *IPsecServiceClient_IPsecStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.IPsecStatsRequest), variadicArgs...)
	})
	return _c
}

func (_c *IPsecServiceClient_IPsecStats_Call) Return(_a0 *_go.IPsecStatsResponse, _a1 error) *IPsecServiceClient_IPsecStats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IPsecServiceClient_IPsecStats_Call) RunAndReturn(run func(context.Context, *_go.IPsecStatsRequest, ...grpc.CallOption) (*_go.IPsecStatsResponse, error)) *IPsecServiceClient_IPsecStats_Call {
	_c.Call.Return(run)
	return _c
}

// IPsecTerminate provides a mock function with given fields: ctx, in, opts
func (_m *IPsecServiceClient) IPsecTerminate(ctx context.Context, in *_go.IPsecTerminateRequest, opts ...grpc.CallOption) (*_go.IPsecTerminateResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.IPsecTerminateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecTerminateRequest, ...grpc.CallOption) (*_go.IPsecTerminateResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecTerminateRequest, ...grpc.CallOption) *_go.IPsecTerminateResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.IPsecTerminateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.IPsecTerminateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPsecServiceClient_IPsecTerminate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPsecTerminate'
type IPsecServiceClient_IPsecTerminate_Call struct {
	*mock.Call
}

// IPsecTerminate is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.IPsecTerminateRequest
//   - opts ...grpc.CallOption
func (_e *IPsecServiceClient_Expecter) IPsecTerminate(ctx interface{}, in interface{}, opts ...interface{}) *IPsecServiceClient_IPsecTerminate_Call {
	return &IPsecServiceClient_IPsecTerminate_Call{Call: _e.mock.On("IPsecTerminate",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *IPsecServiceClient_IPsecTerminate_Call) Run(run func(ctx context.Context, in *_go.IPsecTerminateRequest, opts ...grpc.CallOption)) *IPsecServiceClient_IPsecTerminate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.IPsecTerminateRequest), variadicArgs...)
	})
	return _c
}

func (_c *IPsecServiceClient_IPsecTerminate_Call) Return(_a0 *_go.IPsecTerminateResponse, _a1 error) *IPsecServiceClient_IPsecTerminate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IPsecServiceClient_IPsecTerminate_Call) RunAndReturn(run func(context.Context, *_go.IPsecTerminateRequest, ...grpc.CallOption) (*_go.IPsecTerminateResponse, error)) *IPsecServiceClient_IPsecTerminate_Call {
	_c.Call.Return(run)
	return _c
}

// IPsecUnloadConn provides a mock function with given fields: ctx, in, opts
func (_m *IPsecServiceClient) IPsecUnloadConn(ctx context.Context, in *_go.IPsecUnloadConnRequest, opts ...grpc.CallOption) (*_go.IPsecUnloadConnResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.IPsecUnloadConnResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecUnloadConnRequest, ...grpc.CallOption) (*_go.IPsecUnloadConnResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecUnloadConnRequest, ...grpc.CallOption) *_go.IPsecUnloadConnResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.IPsecUnloadConnResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.IPsecUnloadConnRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPsecServiceClient_IPsecUnloadConn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPsecUnloadConn'
type IPsecServiceClient_IPsecUnloadConn_Call struct {
	*mock.Call
}

// IPsecUnloadConn is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.IPsecUnloadConnRequest
//   - opts ...grpc.CallOption
func (_e *IPsecServiceClient_Expecter) IPsecUnloadConn(ctx interface{}, in interface{}, opts ...interface{}) *IPsecServiceClient_IPsecUnloadConn_Call {
	return &IPsecServiceClient_IPsecUnloadConn_Call{Call: _e.mock.On("IPsecUnloadConn",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *IPsecServiceClient_IPsecUnloadConn_Call) Run(run func(ctx context.Context, in *_go.IPsecUnloadConnRequest, opts ...grpc.CallOption)) *IPsecServiceClient_IPsecUnloadConn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.IPsecUnloadConnRequest), variadicArgs...)
	})
	return _c
}

func (_c *IPsecServiceClient_IPsecUnloadConn_Call) Return(_a0 *_go.IPsecUnloadConnResponse, _a1 error) *IPsecServiceClient_IPsecUnloadConn_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IPsecServiceClient_IPsecUnloadConn_Call) RunAndReturn(run func(context.Context, *_go.IPsecUnloadConnRequest, ...grpc.CallOption) (*_go.IPsecUnloadConnResponse, error)) *IPsecServiceClient_IPsecUnloadConn_Call {
	_c.Call.Return(run)
	return _c
}

// IPsecVersion provides a mock function with given fields: ctx, in, opts
func (_m *IPsecServiceClient) IPsecVersion(ctx context.Context, in *_go.IPsecVersionRequest, opts ...grpc.CallOption) (*_go.IPsecVersionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *_go.IPsecVersionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecVersionRequest, ...grpc.CallOption) (*_go.IPsecVersionResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_go.IPsecVersionRequest, ...grpc.CallOption) *_go.IPsecVersionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*_go.IPsecVersionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_go.IPsecVersionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPsecServiceClient_IPsecVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IPsecVersion'
type IPsecServiceClient_IPsecVersion_Call struct {
	*mock.Call
}

// IPsecVersion is a helper method to define mock.On call
//   - ctx context.Context
//   - in *_go.IPsecVersionRequest
//   - opts ...grpc.CallOption
func (_e *IPsecServiceClient_Expecter) IPsecVersion(ctx interface{}, in interface{}, opts ...interface{}) *IPsecServiceClient_IPsecVersion_Call {
	return &IPsecServiceClient_IPsecVersion_Call{Call: _e.mock.On("IPsecVersion",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *IPsecServiceClient_IPsecVersion_Call) Run(run func(ctx context.Context, in *_go.IPsecVersionRequest, opts ...grpc.CallOption)) *IPsecServiceClient_IPsecVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*_go.IPsecVersionRequest), variadicArgs...)
	})
	return _c
}

func (_c *IPsecServiceClient_IPsecVersion_Call) Return(_a0 *_go.IPsecVersionResponse, _a1 error) *IPsecServiceClient_IPsecVersion_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IPsecServiceClient_IPsecVersion_Call) RunAndReturn(run func(context.Context, *_go.IPsecVersionRequest, ...grpc.CallOption) (*_go.IPsecVersionResponse, error)) *IPsecServiceClient_IPsecVersion_Call {
	_c.Call.Return(run)
	return _c
}

// NewIPsecServiceClient creates a new instance of IPsecServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPsecServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *IPsecServiceClient {
	mock := &IPsecServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

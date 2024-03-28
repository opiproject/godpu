// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"context"
	"errors"
	"testing"

	"github.com/opiproject/godpu/mocks"
	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pc "github.com/opiproject/opi-api/network/opinetcommon/v1alpha1/gen/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestCreateLogicalBridge(t *testing.T) {
	testVni := uint32(500)
	wantIPPefix := &pc.IPPrefix{
		Addr: &pc.IPAddress{
			Af: pc.IpAf_IP_AF_INET,
			V4OrV6: &pc.IPAddress_V4Addr{
				V4Addr: 3232235776,
			},
		},
		Len: 24,
	}
	testLogicalBridge := &pb.LogicalBridge{
		Spec: &pb.LogicalBridgeSpec{
			VlanId:       100,
			Vni:          proto.Uint32(500),
			VtepIpPrefix: wantIPPefix,
		},
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.CreateLogicalBridgeRequest
		wantResponse     *pb.LogicalBridge
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest: &pb.CreateLogicalBridgeRequest{
				LogicalBridgeId: "lb1",
				LogicalBridge:   testLogicalBridge,
			},
			wantResponse:   proto.Clone(testLogicalBridge).(*pb.LogicalBridge),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateLogicalBridgeRequest{
				LogicalBridgeId: "lb1",
				LogicalBridge:   testLogicalBridge,
			},
			wantResponse:   nil,
			wantConnClosed: true,
		},
		"connector err": {
			giveConnectorErr: errors.New("Some conn error"),
			giveClientErr:    nil,
			wantErr:          errors.New("Some conn error"),
			wantRequest:      nil,
			wantConnClosed:   false,
			wantResponse:     nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			mockClient := mocks.NewLogicalBridgeServiceClient(t)
			if tt.wantRequest != nil {
				toReturn := proto.Clone(tt.wantResponse).(*pb.LogicalBridge)
				mockClient.EXPECT().CreateLogicalBridge(mock.Anything, tt.wantRequest).
					Return(toReturn, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewLogicalBridgeWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.LogicalBridgeServiceClient {
					return mockClient
				},
			)

			response, err := c.CreateLogicalBridge(
				context.Background(),
				"lb1", 100, &testVni, "192.168.1.0/24",
			)

			assert.Equal(t, tt.wantErr, err)
			assert.True(t, proto.Equal(response, tt.wantResponse))
			assert.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestDeleteLogicalBridge(t *testing.T) {
	name := "lb2"
	allowMissing := false
	testRequest := &pb.DeleteLogicalBridgeRequest{
		Name:         resourceIDToFullName("bridges", name),
		AllowMissing: allowMissing,
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.DeleteLogicalBridgeRequest
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteLogicalBridgeRequest),
			wantConnClosed:   true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteLogicalBridgeRequest),
			wantConnClosed:   true,
		},
		"connector err": {
			giveConnectorErr: errors.New("Some conn error"),
			giveClientErr:    nil,
			wantErr:          errors.New("Some conn error"),
			wantRequest:      nil,
			wantConnClosed:   false,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			mockClient := mocks.NewLogicalBridgeServiceClient(t)
			if tt.wantRequest != nil {
				mockClient.EXPECT().DeleteLogicalBridge(mock.Anything, tt.wantRequest).
					Return(&emptypb.Empty{}, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewLogicalBridgeWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.LogicalBridgeServiceClient {
					return mockClient
				},
			)

			_, err := c.DeleteLogicalBridge(context.Background(), name, allowMissing)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestGetLogicalBridge(t *testing.T) {
	name := "lb3"
	testRequest := &pb.GetLogicalBridgeRequest{
		Name: resourceIDToFullName("bridges", name),
	}
	testLogicalBridge := &pb.LogicalBridge{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.GetLogicalBridgeRequest
		wantConnClosed   bool
		wantResponse     *pb.LogicalBridge
	}{
		"GetLogicalBridge successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.GetLogicalBridgeRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testLogicalBridge).(*pb.LogicalBridge),
		},
		"GetLogicalBridge client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.GetLogicalBridgeRequest),
			wantConnClosed:   true,
			wantResponse:     nil,
		},
		"connector err": {
			giveConnectorErr: errors.New("Some conn error"),
			giveClientErr:    nil,
			wantErr:          errors.New("Some conn error"),
			wantRequest:      nil,
			wantConnClosed:   false,
			wantResponse:     nil,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			mockClient := mocks.NewLogicalBridgeServiceClient(t)
			if testName == "GetLogicalBridge successful call" {
				mockClient.EXPECT().GetLogicalBridge(mock.Anything, tt.wantRequest).
					Return(&pb.LogicalBridge{}, nil)
			}
			if testName == "GetLogicalBridge client err" {
				mockClient.EXPECT().GetLogicalBridge(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewLogicalBridgeWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.LogicalBridgeServiceClient {
					return mockClient
				},
			)

			response, err := c.GetLogicalBridge(context.Background(), name)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

func TestListLogicalBridges(t *testing.T) {
	pageSize := int32(10)
	pageToken := "def"

	testRequest := &pb.ListLogicalBridgesRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	}
	testLogicalBridgeList := &pb.ListLogicalBridgesResponse{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.ListLogicalBridgesRequest
		wantConnClosed   bool
		wantResponse     *pb.ListLogicalBridgesResponse
	}{
		"ListLogicalBridges successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.ListLogicalBridgesRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testLogicalBridgeList).(*pb.ListLogicalBridgesResponse),
		},
		"ListLogicalBridges client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.ListLogicalBridgesRequest),
			wantConnClosed:   true,
			wantResponse:     nil,
		},
		"connector err": {
			giveConnectorErr: errors.New("Some conn error"),
			giveClientErr:    nil,
			wantErr:          errors.New("Some conn error"),
			wantRequest:      nil,
			wantConnClosed:   false,
			wantResponse:     nil,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			mockClient := mocks.NewLogicalBridgeServiceClient(t)
			if testName == "ListLogicalBridges successful call" {
				mockClient.EXPECT().ListLogicalBridges(mock.Anything, tt.wantRequest).
					Return(&pb.ListLogicalBridgesResponse{}, nil)
			}
			if testName == "ListLogicalBridges client err" {
				mockClient.EXPECT().ListLogicalBridges(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewLogicalBridgeWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.LogicalBridgeServiceClient {
					return mockClient
				},
			)

			response, err := c.ListLogicalBridges(context.Background(), pageSize, pageToken)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

func TestUpdateLogicalBridge(t *testing.T) {
	name := "lb1"
	updateMask := []string{""}
	allowMissing := false

	testRequest := &pb.UpdateLogicalBridgeRequest{
		LogicalBridge: &pb.LogicalBridge{Name: resourceIDToFullName("bridges", name)},
		UpdateMask:    &fieldmaskpb.FieldMask{Paths: updateMask},
		AllowMissing:  allowMissing,
	}

	testLogicalBridgeList := &pb.LogicalBridge{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.UpdateLogicalBridgeRequest
		wantConnClosed   bool
		wantResponse     *pb.LogicalBridge
	}{
		"UpdateLogicalBridge successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.UpdateLogicalBridgeRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testLogicalBridgeList).(*pb.LogicalBridge),
		},
		"UpdateLogicalBridge client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.UpdateLogicalBridgeRequest),
			wantConnClosed:   true,
			wantResponse:     nil,
		},
		"connector err": {
			giveConnectorErr: errors.New("Some conn error"),
			giveClientErr:    nil,
			wantErr:          errors.New("Some conn error"),
			wantRequest:      nil,
			wantConnClosed:   false,
			wantResponse:     nil,
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			mockClient := mocks.NewLogicalBridgeServiceClient(t)
			if testName == "UpdateLogicalBridge successful call" {
				mockClient.EXPECT().UpdateLogicalBridge(mock.Anything, tt.wantRequest).
					Return(&pb.LogicalBridge{}, nil)
			}
			if testName == "UpdateLogicalBridge client err" {
				mockClient.EXPECT().UpdateLogicalBridge(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewLogicalBridgeWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.LogicalBridgeServiceClient {
					return mockClient
				},
			)

			response, err := c.UpdateLogicalBridge(context.Background(), name, updateMask)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

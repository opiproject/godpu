// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/opiproject/godpu/mocks"
	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBridgePort(t *testing.T) {
	macBytes, _ := net.ParseMAC("00:11:22:aa:bb:cc")
	testBridgePort := &pb.BridgePort{
		Spec: &pb.BridgePortSpec{
			MacAddress:     macBytes,
			Ptype:          pb.BridgePortType_BRIDGE_PORT_TYPE_ACCESS,
			LogicalBridges: []string{"lb1", "lb2"},
		},
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.CreateBridgePortRequest
		wantResponse     *pb.BridgePort
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest: &pb.CreateBridgePortRequest{
				BridgePortId: "bp1",
				BridgePort:   testBridgePort,
			},
			wantResponse:   proto.Clone(testBridgePort).(*pb.BridgePort),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateBridgePortRequest{
				BridgePortId: "bp1",
				BridgePort:   testBridgePort,
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
			mockClient := mocks.NewBridgePortServiceClient(t)
			if tt.wantRequest != nil {
				toReturn := proto.Clone(tt.wantResponse).(*pb.BridgePort)
				mockClient.EXPECT().CreateBridgePort(mock.Anything, tt.wantRequest).
					Return(toReturn, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewBridgePortWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.BridgePortServiceClient {
					return mockClient
				},
			)

			response, err := c.CreateBridgePort(
				context.Background(),
				"bp1", "00:11:22:aa:bb:cc", "access", []string{"lb1", "lb2"},
			)

			assert.Equal(t, tt.wantErr, err)
			assert.True(t, proto.Equal(response, tt.wantResponse))
			assert.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestDeleteBridgePort(t *testing.T) {
	name := "bp2"
	allowMissing := false
	testRequest := &pb.DeleteBridgePortRequest{
		Name:         resourceIDToFullName("ports", name),
		AllowMissing: allowMissing,
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.DeleteBridgePortRequest
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteBridgePortRequest),
			wantConnClosed:   true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteBridgePortRequest),
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
			mockClient := mocks.NewBridgePortServiceClient(t)
			if tt.wantRequest != nil {
				mockClient.EXPECT().DeleteBridgePort(mock.Anything, tt.wantRequest).
					Return(&emptypb.Empty{}, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewBridgePortWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.BridgePortServiceClient {
					return mockClient
				},
			)

			_, err := c.DeleteBridgePort(context.Background(), name, allowMissing)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestGetBridgePort(t *testing.T) {
	name := "bp3"
	testRequest := &pb.GetBridgePortRequest{
		Name: resourceIDToFullName("ports", name),
	}
	testBridgePort := &pb.BridgePort{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.GetBridgePortRequest
		wantConnClosed   bool
		wantResponse     *pb.BridgePort
	}{
		"GetBridgePort successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.GetBridgePortRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testBridgePort).(*pb.BridgePort),
		},
		"GetBridgePort client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.GetBridgePortRequest),
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
			mockClient := mocks.NewBridgePortServiceClient(t)
			if testName == "GetBridgePort successful call" {
				mockClient.EXPECT().GetBridgePort(mock.Anything, tt.wantRequest).
					Return(&pb.BridgePort{}, nil)
			}
			if testName == "GetBridgePort client err" {
				mockClient.EXPECT().GetBridgePort(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewBridgePortWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.BridgePortServiceClient {
					return mockClient
				},
			)

			response, err := c.GetBridgePort(context.Background(), name)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

func TestListBridgePorts(t *testing.T) {
	pageSize := int32(10)
	pageToken := "abc"

	testRequest := &pb.ListBridgePortsRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	}
	testBridgePortList := &pb.ListBridgePortsResponse{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.ListBridgePortsRequest
		wantConnClosed   bool
		wantResponse     *pb.ListBridgePortsResponse
	}{
		"ListBridgePorts successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.ListBridgePortsRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testBridgePortList).(*pb.ListBridgePortsResponse),
		},
		"ListBridgePorts client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.ListBridgePortsRequest),
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
			mockClient := mocks.NewBridgePortServiceClient(t)
			if testName == "ListBridgePorts successful call" {
				mockClient.EXPECT().ListBridgePorts(mock.Anything, tt.wantRequest).
					Return(&pb.ListBridgePortsResponse{}, nil)
			}
			if testName == "ListBridgePorts client err" {
				mockClient.EXPECT().ListBridgePorts(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewBridgePortWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.BridgePortServiceClient {
					return mockClient
				},
			)

			response, err := c.ListBridgePorts(context.Background(), pageSize, pageToken)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

func TestUpdateBridgePort(t *testing.T) {
	name := "bp1"
	updateMask := []string{""}
	allowMissing := false

	testRequest := &pb.UpdateBridgePortRequest{
		BridgePort:   &pb.BridgePort{Name: resourceIDToFullName("ports", name)},
		UpdateMask:   &fieldmaskpb.FieldMask{Paths: updateMask},
		AllowMissing: allowMissing,
	}

	testBridgePortList := &pb.BridgePort{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.UpdateBridgePortRequest
		wantConnClosed   bool
		wantResponse     *pb.BridgePort
	}{
		"UpdateBridgePort successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.UpdateBridgePortRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testBridgePortList).(*pb.BridgePort),
		},
		"UpdateBridgePort client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.UpdateBridgePortRequest),
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
			mockClient := mocks.NewBridgePortServiceClient(t)
			if testName == "UpdateBridgePort successful call" {
				mockClient.EXPECT().UpdateBridgePort(mock.Anything, tt.wantRequest).
					Return(&pb.BridgePort{}, nil)
			}
			if testName == "UpdateBridgePort client err" {
				mockClient.EXPECT().UpdateBridgePort(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewBridgePortWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.BridgePortServiceClient {
					return mockClient
				},
			)

			response, err := c.UpdateBridgePort(context.Background(), name, updateMask, allowMissing)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

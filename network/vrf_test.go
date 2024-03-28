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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestCreateVrf(t *testing.T) {
	loopback, err := parseIPAndPrefix("192.168.1.1/24")
	assert.NoError(t, err)
	var vni = uint32(100)
	vtep, err := parseIPAndPrefix("10.0.0.1/32")
	assert.NoError(t, err)

	testVrf := &pb.Vrf{
		Spec: &pb.VrfSpec{
			Vni:              &vni,
			LoopbackIpPrefix: loopback,
			VtepIpPrefix:     vtep,
		},
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.CreateVrfRequest
		wantResponse     *pb.Vrf
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest: &pb.CreateVrfRequest{
				VrfId: "Vrf1",
				Vrf:   testVrf,
			},
			wantResponse:   proto.Clone(testVrf).(*pb.Vrf),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateVrfRequest{
				VrfId: "Vrf1",
				Vrf:   testVrf,
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
			mockClient := mocks.NewVrfServiceClient(t)
			if tt.wantRequest != nil {
				toReturn := proto.Clone(tt.wantResponse).(*pb.Vrf)
				mockClient.EXPECT().CreateVrf(mock.Anything, tt.wantRequest).
					Return(toReturn, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewVRFWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.VrfServiceClient {
					return mockClient
				},
			)

			response, err := c.CreateVrf(context.Background(), "Vrf1", &vni, "192.168.1.1/24", "10.0.0.1/32")

			assert.Equal(t, tt.wantErr, err)
			assert.True(t, proto.Equal(response, tt.wantResponse))
			assert.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestDeleteVrf(t *testing.T) {
	name := "Vrf3"
	allowMissing := false
	testRequest := &pb.DeleteVrfRequest{
		Name:         resourceIDToFullName("vrfs", name),
		AllowMissing: allowMissing,
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.DeleteVrfRequest
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteVrfRequest),
			wantConnClosed:   true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteVrfRequest),
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
			mockClient := mocks.NewVrfServiceClient(t)
			if tt.wantRequest != nil {
				mockClient.EXPECT().DeleteVrf(mock.Anything, tt.wantRequest).
					Return(&emptypb.Empty{}, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewVRFWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.VrfServiceClient {
					return mockClient
				},
			)

			_, err := c.DeleteVrf(context.Background(), name, allowMissing)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestGetVrf(t *testing.T) {
	name := "Vrf1"
	testRequest := &pb.GetVrfRequest{
		Name: resourceIDToFullName("vrfs", name),
	}
	testVrf := &pb.Vrf{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.GetVrfRequest
		wantConnClosed   bool
		wantResponse     *pb.Vrf
	}{
		"GetVrf successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.GetVrfRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testVrf).(*pb.Vrf),
		},
		"GetVrf client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.GetVrfRequest),
			wantConnClosed:   true,
			wantResponse:     nil,
		},
		"GetVrf connector err": {
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
			mockClient := mocks.NewVrfServiceClient(t)
			if testName == "GetVrf successful call" {
				mockClient.EXPECT().GetVrf(mock.Anything, tt.wantRequest).
					Return(&pb.Vrf{}, nil)
			}
			if testName == "GetVrf client err" {
				mockClient.EXPECT().GetVrf(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewVRFWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.VrfServiceClient {
					return mockClient
				},
			)

			response, err := c.GetVrf(context.Background(), name)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

func TestListVrfs(t *testing.T) {
	pageSize := int32(10)
	pageToken := "jkl"

	testRequest := &pb.ListVrfsRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	}
	testVrfList := &pb.ListVrfsResponse{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.ListVrfsRequest
		wantConnClosed   bool
		wantResponse     *pb.ListVrfsResponse
	}{
		"ListVrfs successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.ListVrfsRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testVrfList).(*pb.ListVrfsResponse),
		},
		"ListVrfs client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.ListVrfsRequest),
			wantConnClosed:   true,
			wantResponse:     nil,
		},
		"ListVrfs connector err": {
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
			mockClient := mocks.NewVrfServiceClient(t)
			if testName == "ListVrfs successful call" {
				mockClient.EXPECT().ListVrfs(mock.Anything, tt.wantRequest).
					Return(&pb.ListVrfsResponse{}, nil)
			}
			if testName == "ListVrfs client err" {
				mockClient.EXPECT().ListVrfs(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewVRFWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.VrfServiceClient {
					return mockClient
				},
			)

			response, err := c.ListVrfs(context.Background(), pageSize, pageToken)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

func TestUpdateVrf(t *testing.T) {
	name := "Vrf1"
	updateMask := []string{""}
	allowMissing := false

	testRequest := &pb.UpdateVrfRequest{
		Vrf:          &pb.Vrf{Name: resourceIDToFullName("vrfs", name)},
		UpdateMask:   &fieldmaskpb.FieldMask{Paths: updateMask},
		AllowMissing: allowMissing,
	}

	testVrfList := &pb.Vrf{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.UpdateVrfRequest
		wantConnClosed   bool
		wantResponse     *pb.Vrf
	}{
		"UpdateVrf successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.UpdateVrfRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testVrfList).(*pb.Vrf),
		},
		"UpdateVrf client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.UpdateVrfRequest),
			wantConnClosed:   true,
			wantResponse:     nil,
		},
		"UpdateVrf connector err": {
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
			mockClient := mocks.NewVrfServiceClient(t)
			if testName == "UpdateVrf successful call" {
				mockClient.EXPECT().UpdateVrf(mock.Anything, tt.wantRequest).
					Return(&pb.Vrf{}, nil)
			}
			if testName == "UpdateVrf client err" {
				mockClient.EXPECT().UpdateVrf(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewVRFWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.VrfServiceClient {
					return mockClient
				},
			)

			response, err := c.UpdateVrf(context.Background(), name, updateMask, allowMissing)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

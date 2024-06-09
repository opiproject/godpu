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

	pc "github.com/opiproject/opi-api/network/opinetcommon/v1alpha1/gen/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestCreateSvi(t *testing.T) {
	wantGWPrefixes := []*pc.IPPrefix{
		{
			Addr: &pc.IPAddress{
				Af: pc.IpAf_IP_AF_INET,
				V4OrV6: &pc.IPAddress_V4Addr{
					V4Addr: 0xc0a80101},
			},
			Len: 32,
		},
	}
	macBytes, _ := net.ParseMAC("01:23:45:67:89:ab")

	testSvi := &pb.Svi{
		Spec: &pb.SviSpec{
			Vrf:           "//network.opiproject.org/vrfs/vrf1",
			LogicalBridge: "//network.opiproject.org/bridges/logical1",
			MacAddress:    macBytes,
			GwIpPrefix:    wantGWPrefixes,
			EnableBgp:     true,
			RemoteAs:      65000,
		},
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.CreateSviRequest
		wantResponse     *pb.Svi
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest: &pb.CreateSviRequest{
				SviId: "svi1",
				Svi:   testSvi,
			},
			wantResponse:   proto.Clone(testSvi).(*pb.Svi),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateSviRequest{
				SviId: "svi1",
				Svi:   testSvi,
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
			mockClient := mocks.NewSviServiceClient(t)
			if tt.wantRequest != nil {
				toReturn := proto.Clone(tt.wantResponse).(*pb.Svi)
				mockClient.EXPECT().CreateSvi(mock.Anything, tt.wantRequest).
					Return(toReturn, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewSVIWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.SviServiceClient {
					return mockClient
				},
			)

			response, err := c.CreateSvi(
				context.Background(),
				"svi1", "vrf1", "logical1", "01:23:45:67:89:ab", []string{"192.168.1.1/32"}, true, 65000,
			)

			assert.Equal(t, tt.wantErr, err)
			assert.True(t, proto.Equal(response, tt.wantResponse))
			assert.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestDeleteSvi(t *testing.T) {
	name := "svi2"
	allowMissing := false
	testRequest := &pb.DeleteSviRequest{
		Name:         resourceIDToFullName("svis", name),
		AllowMissing: allowMissing,
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.DeleteSviRequest
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteSviRequest),
			wantConnClosed:   true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteSviRequest),
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
			mockClient := mocks.NewSviServiceClient(t)
			if tt.wantRequest != nil {
				mockClient.EXPECT().DeleteSvi(mock.Anything, tt.wantRequest).
					Return(&emptypb.Empty{}, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewSVIWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.SviServiceClient {
					return mockClient
				},
			)

			_, err := c.DeleteSvi(context.Background(), name, allowMissing)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestGetSvi(t *testing.T) {
	name := "svi3"
	testRequest := &pb.GetSviRequest{
		Name: resourceIDToFullName("svis", name),
	}
	testSvi := &pb.Svi{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.GetSviRequest
		wantConnClosed   bool
		wantResponse     *pb.Svi
	}{
		"GetSvi successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.GetSviRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testSvi).(*pb.Svi),
		},
		"GetSvi client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.GetSviRequest),
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
			mockClient := mocks.NewSviServiceClient(t)
			if testName == "GetSvi successful call" {
				mockClient.EXPECT().GetSvi(mock.Anything, tt.wantRequest).
					Return(&pb.Svi{}, nil)
			}
			if testName == "GetSvi client err" {
				mockClient.EXPECT().GetSvi(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewSVIWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.SviServiceClient {
					return mockClient
				},
			)

			response, err := c.GetSvi(context.Background(), name)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

func TestListSvis(t *testing.T) {
	pageSize := int32(10)
	pageToken := "ghi"

	testRequest := &pb.ListSvisRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	}
	testSviList := &pb.ListSvisResponse{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.ListSvisRequest
		wantConnClosed   bool
		wantResponse     *pb.ListSvisResponse
	}{
		"ListSvi successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.ListSvisRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testSviList).(*pb.ListSvisResponse),
		},
		"ListSvi client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.ListSvisRequest),
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
			mockClient := mocks.NewSviServiceClient(t)
			if testName == "ListSvi successful call" {
				mockClient.EXPECT().ListSvis(mock.Anything, tt.wantRequest).
					Return(&pb.ListSvisResponse{}, nil)
			}
			if testName == "ListSvi client err" {
				mockClient.EXPECT().ListSvis(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewSVIWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.SviServiceClient {
					return mockClient
				},
			)

			response, err := c.ListSvis(context.Background(), pageSize, pageToken)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

func TestUpdateSvi(t *testing.T) {
	name := "svi4"
	updateMask := []string{""}
	allowMissing := false

	testRequest := &pb.UpdateSviRequest{
		Svi:          &pb.Svi{Name: resourceIDToFullName("svis", name)},
		UpdateMask:   &fieldmaskpb.FieldMask{Paths: updateMask},
		AllowMissing: allowMissing,
	}

	testSviList := &pb.Svi{}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.UpdateSviRequest
		wantConnClosed   bool
		wantResponse     *pb.Svi
	}{
		"UpdateSvi successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.UpdateSviRequest),
			wantConnClosed:   true,
			wantResponse:     proto.Clone(testSviList).(*pb.Svi),
		},
		"UpdateSvi client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.UpdateSviRequest),
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
			mockClient := mocks.NewSviServiceClient(t)
			if testName == "UpdateSvi successful call" {
				mockClient.EXPECT().UpdateSvi(mock.Anything, tt.wantRequest).
					Return(&pb.Svi{}, nil)
			}
			if testName == "UpdateSvi client err" {
				mockClient.EXPECT().UpdateSvi(mock.Anything, tt.wantRequest).
					Return(nil, tt.giveClientErr)
			}
			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewSVIWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.SviServiceClient {
					return mockClient
				},
			)

			response, err := c.UpdateSvi(context.Background(), name, updateMask, allowMissing)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantConnClosed, connClosed)
			assert.True(t, proto.Equal(response, tt.wantResponse))
		})
	}
}

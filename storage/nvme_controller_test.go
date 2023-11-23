// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/opiproject/godpu/mocks"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestCreateNvmeTCPController(t *testing.T) {
	controllerID := "nvmetcp0"
	subsystemName := "subsysTcp0Name"
	ipV4Addr := net.ParseIP("127.0.0.1")
	ipV6Addr := net.ParseIP("::")
	testIPV4Controller := &pb.NvmeController{
		Spec: &pb.NvmeControllerSpec{
			Trtype: pb.NvmeTransportType_NVME_TRANSPORT_TCP,
			Endpoint: &pb.NvmeControllerSpec_FabricsId{
				FabricsId: &pb.FabricsEndpoint{
					Traddr:  ipV4Addr.String(),
					Trsvcid: "4420",
					Adrfam:  pb.NvmeAddressFamily_NVME_ADRFAM_IPV4,
				},
			},
		},
	}
	testIPV6Controller := &pb.NvmeController{
		Spec: &pb.NvmeControllerSpec{
			Trtype: pb.NvmeTransportType_NVME_TRANSPORT_TCP,
			Endpoint: &pb.NvmeControllerSpec_FabricsId{
				FabricsId: &pb.FabricsEndpoint{
					Traddr:  ipV6Addr.String(),
					Trsvcid: "4420",
					Adrfam:  pb.NvmeAddressFamily_NVME_ADRFAM_IPV6,
				},
			},
		},
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		giveIP           net.IP
		wantErr          error
		wantRequest      *pb.CreateNvmeControllerRequest
		wantResponse     *pb.NvmeController
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			giveIP:           ipV4Addr,
			wantErr:          nil,
			wantRequest: &pb.CreateNvmeControllerRequest{
				Parent:           subsystemName,
				NvmeControllerId: controllerID,
				NvmeController:   proto.Clone(testIPV4Controller).(*pb.NvmeController),
			},
			wantResponse:   proto.Clone(testIPV4Controller).(*pb.NvmeController),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			giveIP:           ipV4Addr,
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateNvmeControllerRequest{
				Parent:           subsystemName,
				NvmeControllerId: controllerID,
				NvmeController:   proto.Clone(testIPV4Controller).(*pb.NvmeController),
			},
			wantResponse:   nil,
			wantConnClosed: true,
		},
		"connector err": {
			giveConnectorErr: errors.New("Some conn error"),
			giveClientErr:    nil,
			giveIP:           ipV4Addr,
			wantErr:          errors.New("Some conn error"),
			wantRequest:      nil,
			wantResponse:     nil,
			wantConnClosed:   false,
		},
		"ipv6 address": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			giveIP:           ipV6Addr,
			wantErr:          nil,
			wantRequest: &pb.CreateNvmeControllerRequest{
				Parent:           subsystemName,
				NvmeControllerId: controllerID,
				NvmeController:   proto.Clone(testIPV6Controller).(*pb.NvmeController),
			},
			wantResponse:   proto.Clone(testIPV6Controller).(*pb.NvmeController),
			wantConnClosed: true,
		},
		"invalid address": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			giveIP:           net.IP{},
			wantErr:          errors.New("invalid ip address format: <nil>"),
			wantRequest:      nil,
			wantResponse:     nil,
			wantConnClosed:   true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			mockClient := mocks.NewFrontendNvmeServiceClient(t)
			if tt.wantRequest != nil {
				toReturn := proto.Clone(tt.wantResponse).(*pb.NvmeController)
				mockClient.EXPECT().CreateNvmeController(ctx, tt.wantRequest).
					Return(toReturn, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.FrontendNvmeServiceClient {
					return mockClient
				},
			)

			response, err := c.CreateNvmeTCPController(
				ctx,
				controllerID,
				subsystemName,
				tt.giveIP,
				4420,
			)

			require.Equal(t, tt.wantErr, err)
			require.True(t, proto.Equal(response, tt.wantResponse))
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestDeleteNvmeController(t *testing.T) {
	testControllerName := "nvmetcp0Name"
	testRequest := &pb.DeleteNvmeControllerRequest{
		Name:         testControllerName,
		AllowMissing: true,
	}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.DeleteNvmeControllerRequest
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteNvmeControllerRequest),
			wantConnClosed:   true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteNvmeControllerRequest),
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
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			mockClient := mocks.NewFrontendNvmeServiceClient(t)
			if tt.wantRequest != nil {
				mockClient.EXPECT().DeleteNvmeController(ctx, tt.wantRequest).
					Return(&emptypb.Empty{}, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			)

			c, _ := NewWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.FrontendNvmeServiceClient {
					return mockClient
				},
			)

			err := c.DeleteNvmeController(ctx, testControllerName, true)

			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

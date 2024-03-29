// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package backend implements the go library for OPI backend storage
package backend

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/opiproject/godpu/mocks"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestCreateNvmeController(t *testing.T) {
	testControllerID := "remotenvme0"
	testController := &pb.NvmeRemoteController{
		Multipath: pb.NvmeMultipath_NVME_MULTIPATH_FAILOVER,
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.CreateNvmeRemoteControllerRequest
		wantResponse     *pb.NvmeRemoteController
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest: &pb.CreateNvmeRemoteControllerRequest{
				NvmeRemoteControllerId: testControllerID,
				NvmeRemoteController:   proto.Clone(testController).(*pb.NvmeRemoteController),
			},
			wantResponse:   proto.Clone(testController).(*pb.NvmeRemoteController),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateNvmeRemoteControllerRequest{
				NvmeRemoteControllerId: testControllerID,
				NvmeRemoteController:   proto.Clone(testController).(*pb.NvmeRemoteController),
			},
			wantResponse:   nil,
			wantConnClosed: true,
		},
		"connector err": {
			giveConnectorErr: errors.New("Some conn error"),
			giveClientErr:    nil,
			wantErr:          errors.New("Some conn error"),
			wantRequest:      nil,
			wantResponse:     nil,
			wantConnClosed:   false,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			mockClient := mocks.NewNvmeRemoteControllerServiceClient(t)
			if tt.wantRequest != nil {
				toReturn := proto.Clone(tt.wantResponse).(*pb.NvmeRemoteController)
				mockClient.EXPECT().CreateNvmeRemoteController(ctx, tt.wantRequest).
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
				func(grpc.ClientConnInterface) pb.NvmeRemoteControllerServiceClient {
					return mockClient
				},
			)

			response, err := c.CreateNvmeController(
				ctx,
				testControllerID,
				pb.NvmeMultipath_NVME_MULTIPATH_FAILOVER,
			)

			require.Equal(t, tt.wantErr, err)
			require.True(t, proto.Equal(response, tt.wantResponse))
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestDeleteNvmeController(t *testing.T) {
	testControllerName := "remotenvme0"
	testRequest := &pb.DeleteNvmeRemoteControllerRequest{
		Name:         testControllerName,
		AllowMissing: true,
	}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.DeleteNvmeRemoteControllerRequest
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteNvmeRemoteControllerRequest),
			wantConnClosed:   true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteNvmeRemoteControllerRequest),
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

			mockClient := mocks.NewNvmeRemoteControllerServiceClient(t)
			if tt.wantRequest != nil {
				mockClient.EXPECT().DeleteNvmeRemoteController(ctx, tt.wantRequest).
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
				func(grpc.ClientConnInterface) pb.NvmeRemoteControllerServiceClient {
					return mockClient
				},
			)

			err := c.DeleteNvmeController(ctx, testControllerName, true)

			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestGetNvmeController(t *testing.T) {
	testControllerName := "remotenvmeget"
	testRequest := &pb.GetNvmeRemoteControllerRequest{
		Name: testControllerName,
	}
	testController := &pb.NvmeRemoteController{
		Name:      testControllerName,
		Multipath: pb.NvmeMultipath_NVME_MULTIPATH_FAILOVER,
	}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.GetNvmeRemoteControllerRequest
		wantResponse     *pb.NvmeRemoteController
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.GetNvmeRemoteControllerRequest),
			wantResponse:     proto.Clone(testController).(*pb.NvmeRemoteController),
			wantConnClosed:   true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.GetNvmeRemoteControllerRequest),
			wantResponse:     nil,
			wantConnClosed:   true,
		},
		"connector err": {
			giveConnectorErr: errors.New("Some conn error"),
			giveClientErr:    nil,
			wantErr:          errors.New("Some conn error"),
			wantRequest:      nil,
			wantResponse:     nil,
			wantConnClosed:   false,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			mockClient := mocks.NewNvmeRemoteControllerServiceClient(t)
			if tt.wantRequest != nil {
				mockClient.EXPECT().GetNvmeRemoteController(ctx, tt.wantRequest).
					Return(proto.Clone(tt.wantResponse).(*pb.NvmeRemoteController), tt.giveClientErr)
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
				func(grpc.ClientConnInterface) pb.NvmeRemoteControllerServiceClient {
					return mockClient
				},
			)

			response, err := c.GetNvmeController(ctx, testControllerName)

			require.Equal(t, tt.wantErr, err)
			require.True(t, proto.Equal(tt.wantResponse, response))
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package frontend implements the go library for OPI frontend storage
package frontend

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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestCreateVirtioBlk(t *testing.T) {
	controllerID := "virtioblk0"
	volume := "vol0"
	testVirtioBlk := &pb.VirtioBlk{
		PcieId: &pb.PciEndpoint{
			PortId:           wrapperspb.Int32(0),
			PhysicalFunction: wrapperspb.Int32(1),
			VirtualFunction:  wrapperspb.Int32(2),
		},
		VolumeNameRef: volume,
		MaxIoQps:      3,
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.CreateVirtioBlkRequest
		wantResponse     *pb.VirtioBlk
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest: &pb.CreateVirtioBlkRequest{
				VirtioBlkId: controllerID,
				VirtioBlk:   proto.Clone(testVirtioBlk).(*pb.VirtioBlk),
			},
			wantResponse:   proto.Clone(testVirtioBlk).(*pb.VirtioBlk),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateVirtioBlkRequest{
				VirtioBlkId: controllerID,
				VirtioBlk:   proto.Clone(testVirtioBlk).(*pb.VirtioBlk),
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

			mockClient := mocks.NewFrontendVirtioBlkServiceClient(t)
			if tt.wantRequest != nil {
				toReturn := proto.Clone(tt.wantResponse).(*pb.VirtioBlk)
				mockClient.EXPECT().CreateVirtioBlk(ctx, tt.wantRequest).
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
				pb.NewFrontendNvmeServiceClient,
				func(grpc.ClientConnInterface) pb.FrontendVirtioBlkServiceClient {
					return mockClient
				},
			)

			response, err := c.CreateVirtioBlk(ctx, controllerID, volume, 0, 1, 2, 3)

			require.Equal(t, tt.wantErr, err)
			require.True(t, proto.Equal(response, tt.wantResponse))
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestDeleteVirtioBlk(t *testing.T) {
	testControllerName := "virtioBlk0Name"
	testRequest := &pb.DeleteVirtioBlkRequest{
		Name:         testControllerName,
		AllowMissing: true,
	}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.DeleteVirtioBlkRequest
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteVirtioBlkRequest),
			wantConnClosed:   true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteVirtioBlkRequest),
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

			mockClient := mocks.NewFrontendVirtioBlkServiceClient(t)
			if tt.wantRequest != nil {
				mockClient.EXPECT().DeleteVirtioBlk(ctx, tt.wantRequest).
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
				pb.NewFrontendNvmeServiceClient,
				func(grpc.ClientConnInterface) pb.FrontendVirtioBlkServiceClient {
					return mockClient
				},
			)

			err := c.DeleteVirtioBlk(ctx, testControllerName, true)

			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

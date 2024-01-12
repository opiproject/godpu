// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023-2024 Intel Corporation

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

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

func TestCreateNvmeNamespace(t *testing.T) {
	namespaceID := "namespace0"
	volume := "vol0"
	subsystem := "subsys0Name"
	testNamespace := &pb.NvmeNamespace{
		Spec: &pb.NvmeNamespaceSpec{
			VolumeNameRef: volume,
		},
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.CreateNvmeNamespaceRequest
		wantResponse     *pb.NvmeNamespace
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest: &pb.CreateNvmeNamespaceRequest{
				Parent:          subsystem,
				NvmeNamespaceId: namespaceID,
				NvmeNamespace:   proto.Clone(testNamespace).(*pb.NvmeNamespace),
			},
			wantResponse:   proto.Clone(testNamespace).(*pb.NvmeNamespace),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateNvmeNamespaceRequest{
				Parent:          subsystem,
				NvmeNamespaceId: namespaceID,
				NvmeNamespace:   proto.Clone(testNamespace).(*pb.NvmeNamespace),
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

			mockClient := mocks.NewFrontendNvmeServiceClient(t)
			if tt.wantRequest != nil {
				toReturn := proto.Clone(tt.wantResponse).(*pb.NvmeNamespace)
				mockClient.EXPECT().CreateNvmeNamespace(ctx, tt.wantRequest).
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
				pb.NewFrontendVirtioBlkServiceClient,
			)

			response, err := c.CreateNvmeNamespace(ctx, namespaceID, subsystem, volume)

			require.Equal(t, tt.wantErr, err)
			require.True(t, proto.Equal(response, tt.wantResponse))
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestDeleteNvmeNamespace(t *testing.T) {
	testNamespaceName := "name"
	testRequest := &pb.DeleteNvmeNamespaceRequest{
		Name:         testNamespaceName,
		AllowMissing: true,
	}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.DeleteNvmeNamespaceRequest
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteNvmeNamespaceRequest),
			wantConnClosed:   true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteNvmeNamespaceRequest),
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
				mockClient.EXPECT().DeleteNvmeNamespace(ctx, tt.wantRequest).
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
				pb.NewFrontendVirtioBlkServiceClient,
			)

			err := c.DeleteNvmeNamespace(ctx, testNamespaceName, true)

			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

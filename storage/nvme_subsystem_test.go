// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"
	"errors"
	"testing"

	"github.com/opiproject/godpu/mocks"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func TestCreateNvmeSubsystem(t *testing.T) {
	subsystemID := "subsys0"
	nqn := "nqn"
	hostnqn := "hostnqn"
	testSubsystem := &pb.NvmeSubsystem{
		Spec: &pb.NvmeSubsystemSpec{
			Nqn:     nqn,
			Hostnqn: hostnqn,
		},
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.CreateNvmeSubsystemRequest
		wantResponse     *pb.NvmeSubsystem
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest: &pb.CreateNvmeSubsystemRequest{
				NvmeSubsystemId: subsystemID,
				NvmeSubsystem:   proto.Clone(testSubsystem).(*pb.NvmeSubsystem),
			},
			wantResponse:   proto.Clone(testSubsystem).(*pb.NvmeSubsystem),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateNvmeSubsystemRequest{
				NvmeSubsystemId: subsystemID,
				NvmeSubsystem:   proto.Clone(testSubsystem).(*pb.NvmeSubsystem),
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
			mockClient := mocks.NewFrontendNvmeServiceClient(t)
			if tt.wantRequest != nil {
				toReturn := proto.Clone(tt.wantResponse).(*pb.NvmeSubsystem)
				mockClient.EXPECT().CreateNvmeSubsystem(mock.Anything, tt.wantRequest).
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

			response, err := c.CreateNvmeSubsystem(context.Background(), subsystemID, nqn, hostnqn)

			require.Equal(t, tt.wantErr, err)
			require.True(t, proto.Equal(response, tt.wantResponse))
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

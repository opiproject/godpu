// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package backend implements the go library for OPI backend storage
package backend

import (
	"context"
	"errors"
	"fmt"
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

func TestCreateNvmeTCPPath(t *testing.T) {
	testControllerName := "remotenvme0Name"
	testPathID := "remotepath0"
	testIPv4 := net.ParseIP("127.0.0.1")
	testNqn := "nqn.2019-06.io.spdk:0"
	testPath := &pb.NvmePath{
		Trtype: pb.NvmeTransportType_NVME_TRANSPORT_TYPE_TCP,
		Traddr: "127.0.0.1",
		Fabrics: &pb.FabricsPath{
			Trsvcid: 4420,
			Subnqn:  testNqn,
			Adrfam:  pb.NvmeAddressFamily_NVME_ADDRESS_FAMILY_IPV4,
		},
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		giveIP           net.IP
		wantErr          error
		wantRequest      *pb.CreateNvmePathRequest
		wantResponse     *pb.NvmePath
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			giveIP:           testIPv4,
			wantErr:          nil,
			wantRequest: &pb.CreateNvmePathRequest{
				Parent:     testControllerName,
				NvmePathId: testPathID,
				NvmePath:   proto.Clone(testPath).(*pb.NvmePath),
			},
			wantResponse:   proto.Clone(testPath).(*pb.NvmePath),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			giveIP:           testIPv4,
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateNvmePathRequest{
				Parent:     testControllerName,
				NvmePathId: testPathID,
				NvmePath:   proto.Clone(testPath).(*pb.NvmePath),
			},
			wantResponse:   nil,
			wantConnClosed: true,
		},
		"connector err": {
			giveConnectorErr: errors.New("Some conn error"),
			giveClientErr:    nil,
			giveIP:           testIPv4,
			wantErr:          errors.New("Some conn error"),
			wantRequest:      nil,
			wantResponse:     nil,
			wantConnClosed:   false,
		},
		"ipv6 address": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			giveIP:           net.ParseIP("2001:db8::68"),
			wantErr:          nil,
			wantRequest: &pb.CreateNvmePathRequest{
				Parent:     testControllerName,
				NvmePathId: testPathID,
				NvmePath: &pb.NvmePath{
					Trtype: pb.NvmeTransportType_NVME_TRANSPORT_TYPE_TCP,
					Traddr: "2001:db8::68",
					Fabrics: &pb.FabricsPath{
						Trsvcid: 4420,
						Subnqn:  testNqn,
						Adrfam:  pb.NvmeAddressFamily_NVME_ADDRESS_FAMILY_IPV6,
					},
				},
			},
			wantResponse:   &pb.NvmePath{},
			wantConnClosed: true,
		},
		"invalid ip": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			giveIP:           net.ParseIP("invalid ip"),
			wantErr:          fmt.Errorf("invalid ip address format: %v", "<nil>"),
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
				toReturn := proto.Clone(tt.wantResponse).(*pb.NvmePath)
				mockClient.EXPECT().CreateNvmePath(ctx, tt.wantRequest).
					Return(toReturn, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			).Maybe()

			c, _ := NewWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.NvmeRemoteControllerServiceClient {
					return mockClient
				},
			)

			response, err := c.CreateNvmeTCPPath(
				ctx,
				testPathID,
				testControllerName,
				tt.giveIP,
				4420,
				testNqn,
				"",
			)

			require.Equal(t, tt.wantErr, err)
			require.True(t, proto.Equal(response, tt.wantResponse))
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestCreateNvmePciePath(t *testing.T) {
	testControllerName := "remotenvme0Name"
	testPathID := "remotepath0"
	testBDF := "0000:64:00.0"
	testPath := &pb.NvmePath{
		Trtype: pb.NvmeTransportType_NVME_TRANSPORT_TYPE_PCIE,
		Traddr: testBDF,
	}

	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.CreateNvmePathRequest
		wantResponse     *pb.NvmePath
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest: &pb.CreateNvmePathRequest{
				Parent:     testControllerName,
				NvmePathId: testPathID,
				NvmePath:   proto.Clone(testPath).(*pb.NvmePath),
			},
			wantResponse:   proto.Clone(testPath).(*pb.NvmePath),
			wantConnClosed: true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest: &pb.CreateNvmePathRequest{
				Parent:     testControllerName,
				NvmePathId: testPathID,
				NvmePath:   proto.Clone(testPath).(*pb.NvmePath),
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
				toReturn := proto.Clone(tt.wantResponse).(*pb.NvmePath)
				mockClient.EXPECT().CreateNvmePath(ctx, tt.wantRequest).
					Return(toReturn, tt.giveClientErr)
			}

			connClosed := false
			mockConn := mocks.NewConnector(t)
			mockConn.EXPECT().NewConn().Return(
				&grpc.ClientConn{},
				func() { connClosed = true },
				tt.giveConnectorErr,
			).Maybe()

			c, _ := NewWithArgs(
				mockConn,
				func(grpc.ClientConnInterface) pb.NvmeRemoteControllerServiceClient {
					return mockClient
				},
			)

			response, err := c.CreateNvmePciePath(
				ctx,
				testPathID,
				testControllerName,
				testBDF,
			)

			require.Equal(t, tt.wantErr, err)
			require.True(t, proto.Equal(response, tt.wantResponse))
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

func TestDeleteNvmePath(t *testing.T) {
	testPathName := "path0"
	testRequest := &pb.DeleteNvmePathRequest{
		Name:         testPathName,
		AllowMissing: true,
	}
	tests := map[string]struct {
		giveClientErr    error
		giveConnectorErr error
		wantErr          error
		wantRequest      *pb.DeleteNvmePathRequest
		wantConnClosed   bool
	}{
		"successful call": {
			giveConnectorErr: nil,
			giveClientErr:    nil,
			wantErr:          nil,
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteNvmePathRequest),
			wantConnClosed:   true,
		},
		"client err": {
			giveConnectorErr: nil,
			giveClientErr:    errors.New("Some client error"),
			wantErr:          errors.New("Some client error"),
			wantRequest:      proto.Clone(testRequest).(*pb.DeleteNvmePathRequest),
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
				mockClient.EXPECT().DeleteNvmePath(ctx, tt.wantRequest).
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

			err := c.DeleteNvmePath(ctx, testPathName, true)

			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.wantConnClosed, connClosed)
		})
	}
}

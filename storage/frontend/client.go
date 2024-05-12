// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023-2024 Intel Corporation

// Package frontend implements the go library for OPI frontend storage
package frontend

import (
	grpcOpi "github.com/opiproject/godpu/grpc"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/grpc"
)

// CreateFrontendNvmeClient defines the function type used to retrieve FrontendNvmeServiceClient
type CreateFrontendNvmeClient func(cc grpc.ClientConnInterface) pb.FrontendNvmeServiceClient

// CreateFrontendVirtioBlkClient defines the function type used to retrieve FrontendVirtioBlkServiceClient
type CreateFrontendVirtioBlkClient func(cc grpc.ClientConnInterface) pb.FrontendVirtioBlkServiceClient

// Client is used for managing storage devices on OPI server
type Client struct {
	connector                     grpcOpi.Connector
	createFrontendNvmeClient      CreateFrontendNvmeClient
	createFrontendVirtioBlkClient CreateFrontendVirtioBlkClient
}

// New creates a new instance of Client
func New(addr string, tls string) (*Client, error) {
	connector, err := grpcOpi.New(addr, tls)
	if err != nil {
		return nil, err
	}

	return NewWithArgs(
		connector,
		pb.NewFrontendNvmeServiceClient,
		pb.NewFrontendVirtioBlkServiceClient,
	)
}

// NewWithArgs creates a new instance of Client with non-default members
func NewWithArgs(
	connector grpcOpi.Connector,
	createFrontendNvmeClient CreateFrontendNvmeClient,
	createFrontendVirtioBlkClient CreateFrontendVirtioBlkClient,
) (*Client, error) {
	return &Client{
		connector:                     connector,
		createFrontendNvmeClient:      createFrontendNvmeClient,
		createFrontendVirtioBlkClient: createFrontendVirtioBlkClient,
	}, nil
}

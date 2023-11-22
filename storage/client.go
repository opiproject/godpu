// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"time"

	grpcOpi "github.com/opiproject/godpu/grpc"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/grpc"
)

const defaultTimeout = 10 * time.Second

// CreateNvmeClient defines the function type used to retrieve FrontendNvmeServiceClient
type CreateNvmeClient func(cc grpc.ClientConnInterface) pb.FrontendNvmeServiceClient

// Client is used for managing storage devices on OPI server
type Client struct {
	connector    grpcOpi.Connector
	createClient CreateNvmeClient

	timeout time.Duration
}

// New creates a new instance of Client
func New(addr string) (*Client, error) {
	connector, err := grpcOpi.New(addr)
	if err != nil {
		return nil, err
	}

	return NewWithArgs(
		connector,
		pb.NewFrontendNvmeServiceClient,
	)
}

// NewWithArgs creates a new instance of Client with non-default members
func NewWithArgs(
	connector grpcOpi.Connector,
	createClient CreateNvmeClient,
) (*Client, error) {
	return &Client{
		connector:    connector,
		createClient: createClient,
		timeout:      defaultTimeout,
	}, nil
}

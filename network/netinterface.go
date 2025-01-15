// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2025 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI network functions
package network

import (
	"context"
	"log"

	grpcOpi "github.com/opiproject/godpu/grpc"
	pb "github.com/opiproject/opi-api/network/opinetcommon/v1alpha1/gen/go"
	"google.golang.org/grpc"
)

// CreateNetInterfaceClient defines the function type used to retrieve NetInterfaceServiceClient
type CreateNetInterfaceClient func(cc grpc.ClientConnInterface) pb.NetInterfaceServiceClient

// NetIntfClient is used for managing network interfaces on OPI server
type NetIntfClient struct {
	c                        grpcOpi.Connector
	createNetInterfaceClient CreateNetInterfaceClient
}

// NewNetInterface creates a network interface client for use with OPI server at the given address
func NewNetInterface(addr string, tls string) (*NetIntfClient, error) {
	c, err := grpcOpi.New(addr, tls)
	if err != nil {
		return nil, err
	}

	return NewNetInterfaceWithArgs(c, pb.NewNetInterfaceServiceClient)
}

// NewNetInterfaceWithArgs creates a new instance of the network interface client with non-default members
func NewNetInterfaceWithArgs(connector grpcOpi.Connector, createNetInterfaceClient CreateNetInterfaceClient) (*NetIntfClient, error) {
	return &NetIntfClient{
		c:                        connector,
		createNetInterfaceClient: createNetInterfaceClient,
	}, nil
}

// ListNetInterfaces retrieves a list of all network interface details from OPI server
func (c NetIntfClient) ListNetInterfaces(ctx context.Context, pageSize int32, pageToken string) (*pb.ListNetInterfacesResponse, error) {
	conn, closer, err := c.c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.createNetInterfaceClient(conn)

	data, err := client.ListNetInterfaces(ctx, &pb.ListNetInterfacesRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		log.Printf("error List Network Interfaces: %s\n", err)
		return nil, err
	}

	return data, nil
}

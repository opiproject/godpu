// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package inventory implements the go library for OPI to be used to query inventory
package inventory

import (
	"context"
	"errors"
	"log"

	grpcOpi "github.com/opiproject/godpu/grpc"
	pb "github.com/opiproject/opi-api/inventory/v1/gen/go"
	"google.golang.org/grpc"
)

// PbInvClientGetter defines the function type used to retrieve an inventory protobuf client
type PbInvClientGetter func(c grpc.ClientConnInterface) pb.InventorySvcClient

type invClientImpl struct {
	getInvClient PbInvClientGetter
	grpcOpi.Connector
}

// InvClient is an interface for querying inventory data from an OPI server
type InvClient interface {
	Get(ctx context.Context) (*pb.Inventory, error)
}

// New creates an inventory client for use with OPI server at the given address
func New(addr string) (InvClient, error) {
	c, err := grpcOpi.New(addr)
	if err != nil {
		return nil, err
	}

	// Default is to use the OPI grpc client and the pb client generated from the protobuf spec
	return NewWithArgs(c, pb.NewInventorySvcClient)
}

// NewWithArgs creates an inventory client for use with OPI server using the given gRPC client and the given function for
// retrieving an inventory protobuf client
func NewWithArgs(c grpcOpi.Connector, getter PbInvClientGetter) (InvClient, error) {
	if c == nil {
		return nil, errors.New("grpc connector is nil")
	}

	if getter == nil {
		return nil, errors.New("protobuf client getter is nil")
	}

	return invClientImpl{
		getInvClient: getter,
		Connector:    c,
	}, nil
}

// Get returns inventory information an OPI server
func (c invClientImpl) Get(ctx context.Context) (*pb.Inventory, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getInvClient(conn)

	data, err := client.GetInventory(ctx, &pb.GetInventoryRequest{})
	if err != nil {
		log.Printf("error getting inventory: %s\n", err)
		return nil, err
	}

	return data, nil
}

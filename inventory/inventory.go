// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package inventory implements the go library for OPI to be used to query inventory
package inventory

import (
	"context"
	"github.com/opiproject/godpu/common"
	pb "github.com/opiproject/opi-api/common/v1/gen/go"
	"google.golang.org/grpc"
	"log"
)

// PbInvClientGetter defines the function type used to retrieve an inventory protobuf client
type PbInvClientGetter func(c grpc.ClientConnInterface) pb.InventorySvcClient

type clientImpl struct {
	getInvClient PbInvClientGetter
	common.Client
}

// Client provides a function to retrieve inventory from the target OPI server
type Client interface {
	Get(ctx context.Context) (*pb.InventoryGetResponse, error)
}

// NewClient creates an inventory client for use with OPI server at the given address
func NewClient(addr string) (Client, error) {
	return NewClientWithPb(addr, pb.NewInventorySvcClient)
}

// NewClientWithPb creates an inventory client for use with OPI server at the given address using the given function for
// retrieving a inventory protobuf client
func NewClientWithPb(addr string, getter PbInvClientGetter) (Client, error) {
	c, err := common.NewClient(addr)
	if err != nil {
		return nil, err
	}
	return clientImpl{
		// Default is to use the client generated from protobuf spec
		getInvClient: pb.NewInventorySvcClient,
		Client:       c,
	}, nil
}

// Get returns inventory information an OPI server
func (c clientImpl) Get(ctx context.Context) (*pb.InventoryGetResponse, error) {
	conn, closer, err := c.NewGrpcConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getInvClient(conn)

	data, err := client.InventoryGet(ctx, &pb.InventoryGetRequest{})
	if err != nil {
		log.Printf("error getting inventory: %s\n", err)
		return nil, err
	}

	return data, nil
}

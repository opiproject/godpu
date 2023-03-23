// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package inventory implements the go library for OPI to be used to query inventory
package inventory

import (
	"context"
	pb "github.com/opiproject/opi-api/common/v1/gen/go"
	"google.golang.org/grpc"
	"log"
)

// SetAddress changes the gRPC connect address of the DPU
func SetAddress(addr string) {
	address = addr
}

// Get returns inventory information from DPUs
func Get(ctx context.Context, conn *grpc.ClientConn) error {
	client := pb.NewInventorySvcClient(conn)

	data, err := client.InventoryGet(ctx, &pb.InventoryGetRequest{})
	if err != nil {
		return err
	}

	log.Println(data)
	return nil
}

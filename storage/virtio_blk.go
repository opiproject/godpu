// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"

	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// CreateVirtioBlk creates a virtio-blk controller
func (c *Client) CreateVirtioBlk(
	ctx context.Context,
	id, volume string,
	port, pf, vf, maxIoQPS uint,
) (*pb.VirtioBlk, error) {
	conn, connClose, err := c.connector.NewConn()
	if err != nil {
		return nil, err
	}
	defer connClose()

	client := c.createVirtioBlkClient(conn)
	response, err := client.CreateVirtioBlk(
		ctx,
		&pb.CreateVirtioBlkRequest{
			VirtioBlkId: id,
			VirtioBlk: &pb.VirtioBlk{
				PcieId: &pb.PciEndpoint{
					PortId:           wrapperspb.Int32(int32(port)),
					PhysicalFunction: wrapperspb.Int32(int32(pf)),
					VirtualFunction:  wrapperspb.Int32(int32(vf)),
				},
				VolumeNameRef: volume,
				MaxIoQps:      int64(maxIoQPS),
			},
		})

	return response, err
}

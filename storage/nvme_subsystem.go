// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"

	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
)

// CreateNvmeSubsystem creates an nvme subsystem
func (c *Client) CreateNvmeSubsystem(
	ctx context.Context,
	id, nqn, hostnqn string,
) (*pb.NvmeSubsystem, error) {
	conn, connClose, err := c.connector.NewConn()
	if err != nil {
		return nil, err
	}
	defer connClose()

	client := c.createClient(conn)
	response, err := client.CreateNvmeSubsystem(
		ctx,
		&pb.CreateNvmeSubsystemRequest{
			NvmeSubsystemId: id,
			NvmeSubsystem: &pb.NvmeSubsystem{
				Spec: &pb.NvmeSubsystemSpec{
					Nqn:     nqn,
					Hostnqn: hostnqn,
				},
			},
		})

	return response, err
}

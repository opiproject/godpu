// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"

	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
)

// CreateNvmeNamespace creates an nvme namespace
func (c *Client) CreateNvmeNamespace(
	ctx context.Context,
	id, subsystem, volume string,
) (*pb.NvmeNamespace, error) {
	conn, connClose, err := c.connector.NewConn()
	if err != nil {
		return nil, err
	}
	defer connClose()

	client := c.createClient(conn)
	response, err := client.CreateNvmeNamespace(
		ctx,
		&pb.CreateNvmeNamespaceRequest{
			Parent:          subsystem,
			NvmeNamespaceId: id,
			NvmeNamespace: &pb.NvmeNamespace{
				Spec: &pb.NvmeNamespaceSpec{
					VolumeNameRef: volume,
				},
			},
		})

	return response, err
}

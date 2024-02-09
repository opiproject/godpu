// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package backend implements the go library for OPI backend storage
package backend

import (
	"context"

	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
)

// CreateNvmeController creates an nvme controller representing
// an external nvme device
func (c *Client) CreateNvmeController(
	ctx context.Context,
	id string,
	multipath pb.NvmeMultipath,
) (*pb.NvmeRemoteController, error) {
	conn, connClose, err := c.connector.NewConn()
	if err != nil {
		return nil, err
	}
	defer connClose()

	client := c.createNvmeClient(conn)
	response, err := client.CreateNvmeRemoteController(
		ctx,
		&pb.CreateNvmeRemoteControllerRequest{
			NvmeRemoteControllerId: id,
			NvmeRemoteController: &pb.NvmeRemoteController{
				Multipath: multipath,
			},
		})

	return response, err
}

// DeleteNvmeController deletes an nvme controller representing
// an external nvme device
func (c *Client) DeleteNvmeController(
	ctx context.Context,
	name string,
	allowMissing bool,
) error {
	conn, connClose, err := c.connector.NewConn()
	if err != nil {
		return err
	}
	defer connClose()

	client := c.createNvmeClient(conn)
	_, err = client.DeleteNvmeRemoteController(
		ctx,
		&pb.DeleteNvmeRemoteControllerRequest{
			Name:         name,
			AllowMissing: allowMissing,
		})

	return err
}

// GetNvmeController gets an nvme controller representing
// an external nvme device
func (c *Client) GetNvmeController(
	ctx context.Context,
	name string,
) (*pb.NvmeRemoteController, error) {
	conn, connClose, err := c.connector.NewConn()
	if err != nil {
		return nil, err
	}
	defer connClose()

	client := c.createNvmeClient(conn)
	return client.GetNvmeRemoteController(
		ctx,
		&pb.GetNvmeRemoteControllerRequest{
			Name: name,
		})
}

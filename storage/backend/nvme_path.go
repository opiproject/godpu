// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package backend implements the go library for OPI backend storage
package backend

import (
	"context"
	"fmt"
	"net"

	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
)

// CreateNvmeTCPPath creates a path to nvme TCP controller
func (c *Client) CreateNvmeTCPPath(
	ctx context.Context,
	id string,
	controller string,
	ip net.IP,
	port uint16,
	nqn, hostnqn string,
) (*pb.NvmePath, error) {
	var adrfam pb.NvmeAddressFamily
	switch {
	case ip.To4() != nil:
		adrfam = pb.NvmeAddressFamily_NVME_ADDRESS_FAMILY_IPV4
	case ip.To16() != nil:
		adrfam = pb.NvmeAddressFamily_NVME_ADDRESS_FAMILY_IPV6
	default:
		return nil, fmt.Errorf("invalid ip address format: %v", ip)
	}

	conn, connClose, err := c.connector.NewConn()
	if err != nil {
		return nil, err
	}
	defer connClose()

	client := c.createNvmeClient(conn)
	response, err := client.CreateNvmePath(
		ctx,
		&pb.CreateNvmePathRequest{
			NvmePathId: id,
			Parent:     controller,
			NvmePath: &pb.NvmePath{
				Trtype: pb.NvmeTransportType_NVME_TRANSPORT_TYPE_TCP,
				Traddr: ip.String(),
				Fabrics: &pb.FabricsPath{
					Trsvcid: int64(port),
					Subnqn:  nqn,
					Adrfam:  adrfam,
					Hostnqn: hostnqn,
				},
			},
		})

	return response, err
}

// CreateNvmePciePath creates a path to nvme PCIe controller
func (c *Client) CreateNvmePciePath(
	ctx context.Context,
	id string,
	controller string,
	bdf string,
) (*pb.NvmePath, error) {
	conn, connClose, err := c.connector.NewConn()
	if err != nil {
		return nil, err
	}
	defer connClose()

	client := c.createNvmeClient(conn)
	response, err := client.CreateNvmePath(
		ctx,
		&pb.CreateNvmePathRequest{
			NvmePathId: id,
			Parent:     controller,
			NvmePath: &pb.NvmePath{
				Trtype: pb.NvmeTransportType_NVME_TRANSPORT_TYPE_PCIE,
				Traddr: bdf,
			},
		})

	return response, err
}

// DeleteNvmePath deletes an nvme path to an external nvme controller
func (c *Client) DeleteNvmePath(
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
	_, err = client.DeleteNvmePath(
		ctx,
		&pb.DeleteNvmePathRequest{
			Name:         name,
			AllowMissing: allowMissing,
		})

	return err
}

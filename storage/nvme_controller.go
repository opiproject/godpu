// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"
	"fmt"
	"net"

	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
)

// CreateNvmeTCPController creates an nvme TCP controller
func (c *Client) CreateNvmeTCPController(
	ctx context.Context,
	id, subsystem string,
	ip net.IP,
	port uint16,
) (*pb.NvmeController, error) {
	conn, connClose, err := c.connector.NewConn()
	if err != nil {
		return nil, err
	}
	defer connClose()

	var adrfam pb.NvmeAddressFamily
	switch {
	case ip.To4() != nil:
		adrfam = pb.NvmeAddressFamily_NVME_ADRFAM_IPV4
	case ip.To16() != nil:
		adrfam = pb.NvmeAddressFamily_NVME_ADRFAM_IPV6
	default:
		return nil, fmt.Errorf("invalid ip address format: %v", ip)
	}

	client := c.createClient(conn)
	response, err := client.CreateNvmeController(
		ctx,
		&pb.CreateNvmeControllerRequest{
			Parent:           subsystem,
			NvmeControllerId: id,
			NvmeController: &pb.NvmeController{
				Spec: &pb.NvmeControllerSpec{
					Trtype: pb.NvmeTransportType_NVME_TRANSPORT_TCP,
					Endpoint: &pb.NvmeControllerSpec_FabricsId{
						FabricsId: &pb.FabricsEndpoint{
							Traddr:  ip.String(),
							Trsvcid: fmt.Sprint(port),
							Adrfam:  adrfam,
						},
					},
				},
			},
		})

	return response, err
}

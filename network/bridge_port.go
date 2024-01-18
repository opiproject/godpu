// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"context"
	"log"

	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// CreateBridgePort creates an Bridge Port an OPI server
func (c evpnClientImpl) CreateBridgePort(ctx context.Context, name string, mac string, bridgePortType string, logicalBridges []string) (*pb.BridgePort, error) {
	var typeOfPort pb.BridgePortType
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnBridgePortClient(conn)
	switch bridgePortType {
	case "access":
		typeOfPort = pb.BridgePortType_BRIDGE_PORT_TYPE_ACCESS
	case "trunk":
		typeOfPort = pb.BridgePortType_BRIDGE_PORT_TYPE_TRUNK
	default:
		typeOfPort = pb.BridgePortType_BRIDGE_PORT_TYPE_UNSPECIFIED
	}
	data, err := client.CreateBridgePort(ctx, &pb.CreateBridgePortRequest{
		BridgePortId: name,
		BridgePort: &pb.BridgePort{
			Spec: &pb.BridgePortSpec{
				MacAddress:     []byte(mac),
				Ptype:          typeOfPort,
				LogicalBridges: logicalBridges,
			},
		},
	})
	if err != nil {
		log.Printf("error creating Bridge Port: %s\n", err)
		return nil, err
	}

	return data, nil
}

// DeleteBridgePort delete an Bridge Port an OPI server
func (c evpnClientImpl) DeleteBridgePort(ctx context.Context, name string, allowMissing bool) (*emptypb.Empty, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnBridgePortClient(conn)
	data, err := client.DeleteBridgePort(ctx, &pb.DeleteBridgePortRequest{
		Name:         resourceIDToFullName("ports", name),
		AllowMissing: allowMissing,
	})
	if err != nil {
		log.Printf("error deleting  Bridge Port: %s\n", err)
		return nil, err
	}

	return data, nil
}

// GetBridgePort Get Bridge Port details
func (c evpnClientImpl) GetBridgePort(ctx context.Context, name string) (*pb.BridgePort, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnBridgePortClient(conn)
	data, err := client.GetBridgePort(ctx, &pb.GetBridgePortRequest{
		Name: resourceIDToFullName("ports", name),
	})
	if err != nil {
		log.Printf("error getting Bridge Port: %s\n", err)
		return nil, err
	}

	return data, nil
}

// ListBridgePorts list all the Bridge Port an OPI server
func (c evpnClientImpl) ListBridgePorts(ctx context.Context, pageSize int32, pageToken string) (*pb.ListBridgePortsResponse, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnBridgePortClient(conn)
	data, err := client.ListBridgePorts(ctx, &pb.ListBridgePortsRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		log.Printf("error list Bridge Port: %s\n", err)
		return nil, err
	}

	return data, nil
}

// UpdateBridgePort update the Bridge Port on OPI server
func (c evpnClientImpl) UpdateBridgePort(ctx context.Context, name string, updateMask []string, allowMissing bool) (*pb.BridgePort, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnBridgePortClient(conn)
	Port := &pb.BridgePort{
		Name: resourceIDToFullName("ports", name),
	}
	data, err := client.UpdateBridgePort(ctx, &pb.UpdateBridgePortRequest{
		BridgePort:   Port,
		UpdateMask:   &fieldmaskpb.FieldMask{Paths: updateMask},
		AllowMissing: allowMissing,
	})
	if err != nil {
		log.Printf("error updating Bridge Port: %s\n", err)
		return nil, err
	}

	return data, nil
}

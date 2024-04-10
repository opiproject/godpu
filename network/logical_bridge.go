// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"context"
	"errors"
	"log"

	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	pc "github.com/opiproject/opi-api/network/opinetcommon/v1alpha1/gen/go"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// CreateLogicalBridge creates an Logical Bridge an OPI server
func (c evpnClientImpl) CreateLogicalBridge(ctx context.Context, name string, vlanID uint32, vni *uint32, vtepIP string) (*pb.LogicalBridge, error) {
	var ipVtep *pc.IPPrefix

	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnLogicalBridgeClient(conn)

	if (vni == nil && vtepIP != "") || (vni != nil && vtepIP == "") {
		return nil, errors.New("one of the required together parameter [vni, vtep] wasn't passed ")
	}

	if vni != nil && vtepIP != "" {
		ipVtep, err = parseIPAndPrefix(vtepIP)
		if err != nil {
			log.Printf("parseIPAndPrefix: error creating Logical Bridge: %s\n", err)
			return nil, err
		}
	}
	data, err := client.CreateLogicalBridge(ctx, &pb.CreateLogicalBridgeRequest{
		LogicalBridgeId: name,
		LogicalBridge: &pb.LogicalBridge{
			Spec: &pb.LogicalBridgeSpec{
				VlanId:       vlanID,
				Vni:          vni,
				VtepIpPrefix: ipVtep,
			},
		},
	})
	if err != nil {
		log.Printf("error creating logical bridge: %s\n", err)
		return nil, err
	}

	return data, nil
}

// DeleteLogicalBridge deletes an Logical Bridge an OPI server
func (c evpnClientImpl) DeleteLogicalBridge(ctx context.Context, name string, allowMissing bool) (*emptypb.Empty, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	if name == "" {
		return nil, errors.New("required parameter [name] wasn't passed ")
	}
	client := c.getEvpnLogicalBridgeClient(conn)

	data, err := client.DeleteLogicalBridge(ctx, &pb.DeleteLogicalBridgeRequest{
		Name:         resourceIDToFullName("bridges", name),
		AllowMissing: allowMissing,
	})
	if err != nil {
		log.Printf("error deleting logical bridge: %s\n", err)
		return nil, err
	}

	return data, nil
}

// GetLogicalBridge get Logical Bridge details from OPI server
func (c evpnClientImpl) GetLogicalBridge(ctx context.Context, name string) (*pb.LogicalBridge, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	if name == "" {
		return nil, errors.New("required parameter [name] wasn't passed ")
	}
	client := c.getEvpnLogicalBridgeClient(conn)

	data, err := client.GetLogicalBridge(ctx, &pb.GetLogicalBridgeRequest{
		Name: resourceIDToFullName("bridges", name),
	})
	if err != nil {
		log.Printf("error getting logical bridge: %s\n", err)
		return nil, err
	}

	return data, nil
}

// ListLogicalBridges list all Logical Bridge details from OPI server
func (c evpnClientImpl) ListLogicalBridges(ctx context.Context, pageSize int32, pageToken string) (*pb.ListLogicalBridgesResponse, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnLogicalBridgeClient(conn)

	data, err := client.ListLogicalBridges(ctx, &pb.ListLogicalBridgesRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		log.Printf("error List logical bridge: %s\n", err)
		return nil, err
	}

	return data, nil
}

// UpdateLogicalBridge update Logical Bridge on OPI server
func (c evpnClientImpl) UpdateLogicalBridge(ctx context.Context, name string, updateMask []string, allowMissing bool) (*pb.LogicalBridge, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnLogicalBridgeClient(conn)
	Bridge := &pb.LogicalBridge{
		Name: resourceIDToFullName("bridges", name),
	}
	data, err := client.UpdateLogicalBridge(ctx, &pb.UpdateLogicalBridgeRequest{
		LogicalBridge: Bridge,
		UpdateMask:    &fieldmaskpb.FieldMask{Paths: updateMask},
		AllowMissing:  allowMissing,
	})
	if err != nil {
		log.Printf("error Update logical bridge: %s\n", err)
		return nil, err
	}

	return data, nil
}

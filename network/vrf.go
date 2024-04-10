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

// CreateVrf Create vrf on OPI Server
func (c evpnClientImpl) CreateVrf(ctx context.Context, name string, vni *uint32, loopbackIP string, vtepIP string) (*pb.Vrf, error) {
	conn, closer, err := c.NewConn()
	var ipVtep *pc.IPPrefix
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnVRFClient(conn)
	if loopbackIP == "" {
		return nil, errors.New("required together parameter [loopbackIP] wasn't passed ")
	}
	ipLoopback, err := parseIPAndPrefix(loopbackIP)
	if err != nil {
		log.Printf("parseIPAndPrefix: error creating vrf: %s\n", err)
		return nil, err
	}
	if vni != nil && vtepIP != "" {
		ipVtep, err = parseIPAndPrefix(vtepIP)
		if err != nil {
			log.Printf("parseIPAndPrefix: error creating vrf: %s\n", err)
			return nil, err
		}
	}
	data, err := client.CreateVrf(ctx, &pb.CreateVrfRequest{
		VrfId: name,
		Vrf: &pb.Vrf{
			Spec: &pb.VrfSpec{
				Vni:              vni,
				LoopbackIpPrefix: ipLoopback,
				VtepIpPrefix:     ipVtep,
			},
		},
	})
	if err != nil {
		log.Printf("error creating vrf: %s\n", err)
		return nil, err
	}

	return data, nil
}

// DeleteVrf update the vrf on OPI server
func (c evpnClientImpl) DeleteVrf(ctx context.Context, name string, allowMissing bool) (*emptypb.Empty, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	if name == "" {
		return nil, errors.New("required parameter [name] wasn't passed ")
	}

	client := c.getEvpnVRFClient(conn)
	data, err := client.DeleteVrf(ctx, &pb.DeleteVrfRequest{
		Name:         resourceIDToFullName("vrfs", name),
		AllowMissing: allowMissing,
	})
	if err != nil {
		log.Printf("error deleting vrf: %s\n", err)
		return nil, err
	}

	return data, nil
}

// GetVrf get vrf details from OPI server
func (c evpnClientImpl) GetVrf(ctx context.Context, name string) (*pb.Vrf, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	if name == "" {
		return nil, errors.New("required parameter [name] wasn't passed ")
	}

	client := c.getEvpnVRFClient(conn)
	data, err := client.GetVrf(ctx, &pb.GetVrfRequest{
		Name: resourceIDToFullName("vrfs", name),
	})
	if err != nil {
		log.Printf("error getting vrf: %s\n", err)
		return nil, err
	}

	return data, nil
}

// ListVrfs list all vrf's with details from OPI server
func (c evpnClientImpl) ListVrfs(ctx context.Context, pageSize int32, pageToken string) (*pb.ListVrfsResponse, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()
	client := c.getEvpnVRFClient(conn)
	data, err := client.ListVrfs(ctx, &pb.ListVrfsRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		log.Printf("error list Vrf: %s\n", err)
		return nil, err
	}

	return data, nil
}

// UpdateVrf update the vrf on OPI server
func (c evpnClientImpl) UpdateVrf(ctx context.Context, name string, updateMask []string, allowMissing bool) (*pb.Vrf, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()
	vrf := &pb.Vrf{
		Name: resourceIDToFullName("vrfs", name),
	}
	client := c.getEvpnVRFClient(conn)
	data, err := client.UpdateVrf(ctx, &pb.UpdateVrfRequest{
		Vrf:          vrf,
		UpdateMask:   &fieldmaskpb.FieldMask{Paths: updateMask},
		AllowMissing: allowMissing,
	})
	if err != nil {
		log.Printf("error updating vrf: %s\n", err)
		return nil, err
	}

	return data, nil
}

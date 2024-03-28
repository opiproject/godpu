// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// CreateSvi create svi on OPI server
func (c evpnClientImpl) CreateSvi(ctx context.Context, name string, vrf string, logicalBridge string, mac string, gwIPs []string, ebgp bool, remoteAS uint32) (*pb.Svi, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnSVIClient(conn)

	gwPrefixes, err := parseIPPrefixes(gwIPs)
	if err != nil {
		log.Printf("error parsing GwIPs: %s\n", err)
		return nil, err
	}
	macBytes, err := net.ParseMAC(mac)
	if err != nil {
		fmt.Println("Error parsing MAC address:", err)
		return nil, err
	}
	data, err := client.CreateSvi(ctx, &pb.CreateSviRequest{
		SviId: name,
		Svi: &pb.Svi{
			Spec: &pb.SviSpec{
				Vrf:           vrf,
				LogicalBridge: logicalBridge,
				MacAddress:    macBytes,
				GwIpPrefix:    gwPrefixes,
				EnableBgp:     ebgp,
				RemoteAs:      remoteAS,
			},
		},
	})
	if err != nil {
		log.Printf("error creating svi: %s\n", err)
		return nil, err
	}

	return data, nil
}

// DeleteSvi delete the svi on OPI server
func (c evpnClientImpl) DeleteSvi(ctx context.Context, name string, allowMissing bool) (*emptypb.Empty, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()
	client := c.getEvpnSVIClient(conn)
	data, err := client.DeleteSvi(ctx, &pb.DeleteSviRequest{
		Name:         resourceIDToFullName("svis", name),
		AllowMissing: allowMissing,
	})
	if err != nil {
		log.Printf("error deleting svi: %s\n", err)
		return nil, err
	}

	return data, nil
}

// GetSvi get svi details from OPI server
func (c evpnClientImpl) GetSvi(ctx context.Context, name string) (*pb.Svi, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()
	client := c.getEvpnSVIClient(conn)
	data, err := client.GetSvi(ctx, &pb.GetSviRequest{
		Name: resourceIDToFullName("svis", name),
	})
	if err != nil {
		log.Printf("error getting svi: %s\n", err)
		return nil, err
	}

	return data, nil
}

// ListSvis get all the svi's from OPI server
func (c evpnClientImpl) ListSvis(ctx context.Context, pageSize int32, pageToken string) (*pb.ListSvisResponse, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()
	client := c.getEvpnSVIClient(conn)
	data, err := client.ListSvis(ctx, &pb.ListSvisRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
	})
	if err != nil {
		log.Printf("error list svis: %s\n", err)
		return nil, err
	}

	return data, nil
}

// UpdateSvi update the svi on OPI server
func (c evpnClientImpl) UpdateSvi(ctx context.Context, name string, updateMask []string, allowMissing bool) (*pb.Svi, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()
	client := c.getEvpnSVIClient(conn)

	svi := &pb.Svi{
		Name: resourceIDToFullName("svis", name),
	}
	data, err := client.UpdateSvi(ctx, &pb.UpdateSviRequest{
		Svi:          svi,
		UpdateMask:   &fieldmaskpb.FieldMask{Paths: updateMask},
		AllowMissing: allowMissing,
	})
	if err != nil {
		log.Printf("error updating svi: %s\n", err)
		return nil, err
	}

	return data, nil
}

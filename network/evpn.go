// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"context"
	"errors"
	"log"

	grpcOpi "github.com/opiproject/godpu/grpc"
	pb "github.com/opiproject/opi-api/network/cloud/v1alpha1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// PbEvpnClientGetter defines the function type used to retrieve an evpn protobuf client
type PbEvpnClientGetter func(c grpc.ClientConnInterface) pb.CloudInfraServiceClient

type evpnClientImpl struct {
	getEvpnClient PbEvpnClientGetter
	grpcOpi.Connector
}

// EvpnClient is an interface for querying evpn data from an OPI server
type EvpnClient interface {
	// CreateSubnet(ctx context.Context) (*pb.Subnet, error)
	// DeleteSubnet(ctx context.Context) (*emptypb.Empty, error)
	CreateInterface(ctx context.Context) (*pb.Interface, error)
	GetInterface(ctx context.Context) (*pb.Interface, error)
	DeleteInterface(ctx context.Context) (*emptypb.Empty, error)
	CreateVpc(ctx context.Context) (*pb.Vpc, error)
	GetVpc(ctx context.Context) (*pb.Vpc, error)
	DeleteVpc(ctx context.Context) (*emptypb.Empty, error)
}

// New creates an evpn client for use with OPI server at the given address
func New(addr string) (EvpnClient, error) {
	c, err := grpcOpi.New(addr)
	if err != nil {
		return nil, err
	}

	// Default is to use the OPI grpc client and the pb client generated from the protobuf spec
	return NewWithArgs(c, pb.NewCloudInfraServiceClient)
}

// NewWithArgs creates an evpn client for use with OPI server using the given gRPC client and the given function for
// retrieving an evpn protobuf client
func NewWithArgs(c grpcOpi.Connector, getter PbEvpnClientGetter) (EvpnClient, error) {
	if c == nil {
		return nil, errors.New("grpc connector is nil")
	}

	if getter == nil {
		return nil, errors.New("protobuf client getter is nil")
	}

	return evpnClientImpl{
		getEvpnClient: getter,
		Connector:     c,
	}, nil
}

// CreateInterface creates an interface an OPI server
func (c evpnClientImpl) CreateInterface(ctx context.Context) (*pb.Interface, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnClient(conn)

	data, err := client.CreateInterface(ctx, &pb.CreateInterfaceRequest{
		Parent:      "todo",
		InterfaceId: "testinterface",
		Interface: &pb.Interface{
			Spec: &pb.InterfaceSpec{
				Ifid: 11,
			},
		},
	})
	if err != nil {
		log.Printf("error creating evpn: %s\n", err)
		return nil, err
	}

	return data, nil
}

// GetInterface creates an interface an OPI server
func (c evpnClientImpl) GetInterface(ctx context.Context) (*pb.Interface, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnClient(conn)

	data, err := client.GetInterface(ctx, &pb.GetInterfaceRequest{
		Name: "//network.opiproject.org/interfaces/testinterface",
	})
	if err != nil {
		log.Printf("error getting evpn: %s\n", err)
		return nil, err
	}

	return data, nil
}

// DeleteInterface creates an interface an OPI server
func (c evpnClientImpl) DeleteInterface(ctx context.Context) (*emptypb.Empty, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnClient(conn)

	data, err := client.DeleteInterface(ctx, &pb.DeleteInterfaceRequest{
		Name: "//network.opiproject.org/interfaces/testinterface",
	})
	if err != nil {
		log.Printf("error deleting evpn: %s\n", err)
		return nil, err
	}

	return data, nil
}

// CreateVpc creates an Vpc an OPI server
func (c evpnClientImpl) CreateVpc(ctx context.Context) (*pb.Vpc, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnClient(conn)

	data, err := client.CreateVpc(ctx, &pb.CreateVpcRequest{
		Parent: "todo",
		VpcId:  "testVpc",
		Vpc: &pb.Vpc{
			Spec: &pb.VpcSpec{
				V4RouteTableNameRef: "1000",
			},
		},
	})
	if err != nil {
		log.Printf("error creating evpn: %s\n", err)
		return nil, err
	}

	return data, nil
}

// GetVpc creates an Vpc an OPI server
func (c evpnClientImpl) GetVpc(ctx context.Context) (*pb.Vpc, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnClient(conn)

	data, err := client.GetVpc(ctx, &pb.GetVpcRequest{
		Name: "//network.opiproject.org/vpcs/testVpc",
	})
	if err != nil {
		log.Printf("error getting evpn: %s\n", err)
		return nil, err
	}

	return data, nil
}

// DeleteVpc creates an Vpc an OPI server
func (c evpnClientImpl) DeleteVpc(ctx context.Context) (*emptypb.Empty, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnClient(conn)

	data, err := client.DeleteVpc(ctx, &pb.DeleteVpcRequest{
		Name: "//network.opiproject.org/vpcs/testVpc",
	})
	if err != nil {
		log.Printf("error deleting evpn: %s\n", err)
		return nil, err
	}

	return data, nil
}

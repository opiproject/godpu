// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"context"
	"errors"
	"fmt"

	grpcOpi "github.com/opiproject/godpu/grpc"
	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// PbEvpnLogicalBridgeClientGetter defines the function type used to retrieve an evpn protobuf LogicalBridgeServiceClient
type PbEvpnLogicalBridgeClientGetter func(c grpc.ClientConnInterface) pb.LogicalBridgeServiceClient

// PbEvpnBridgePortClientGetter defines the function type used to retrieve an evpn protobuf BridgePortServiceClient
type PbEvpnBridgePortClientGetter func(c grpc.ClientConnInterface) pb.BridgePortServiceClient

// PbEvpnVRFClientGetter defines the function type used to retrieve an evpn protobuf VrfServiceClient
type PbEvpnVRFClientGetter func(c grpc.ClientConnInterface) pb.VrfServiceClient

// PbEvpnSVIClientGetter defines the function type used to retrieve an evpn protobuf SviServiceClient
type PbEvpnSVIClientGetter func(c grpc.ClientConnInterface) pb.SviServiceClient

type evpnClientImpl struct {
	// getEvpnClient PbEvpnClientGetter
	getEvpnLogicalBridgeClient PbEvpnLogicalBridgeClientGetter
	getEvpnBridgePortClient    PbEvpnBridgePortClientGetter
	getEvpnVRFClient           PbEvpnVRFClientGetter
	getEvpnSVIClient           PbEvpnSVIClientGetter
	grpcOpi.Connector
}

// EvpnClient is an interface for querying evpn data from an OPI server
type EvpnClient interface {

	// Logical Bridge interfaces
	CreateLogicalBridge(ctx context.Context, name string, vlanID uint32, vni uint32, vtepIP string) (*pb.LogicalBridge, error)
	DeleteLogicalBridge(ctx context.Context, name string, allowMissing bool) (*emptypb.Empty, error)
	GetLogicalBridge(ctx context.Context, name string) (*pb.LogicalBridge, error)
	ListLogicalBridges(ctx context.Context, pageSize int32, pageToken string) (*pb.ListLogicalBridgesResponse, error)
	UpdateLogicalBridge(ctx context.Context, name string, updateMask []string) (*pb.LogicalBridge, error)

	// Bridge Port Interfaces
	CreateBridgePort(ctx context.Context, name string, mac string, bridgePortType string, logicalBridges []string) (*pb.BridgePort, error)
	DeleteBridgePort(ctx context.Context, name string, allowMissing bool) (*emptypb.Empty, error)
	GetBridgePort(ctx context.Context, name string) (*pb.BridgePort, error)
	ListBridgePorts(ctx context.Context, pageSize int32, pageToken string) (*pb.ListBridgePortsResponse, error)
	UpdateBridgePort(ctx context.Context, name string, updateMask []string, allowMissing bool) (*pb.BridgePort, error)

	// VRF Interfaces
	CreateVrf(ctx context.Context, name string, vni uint32, loopback string, vtep string) (*pb.Vrf, error)
	DeleteVrf(ctx context.Context, name string, allowMissing bool) (*emptypb.Empty, error)
	GetVrf(ctx context.Context, name string) (*pb.Vrf, error)
	ListVrfs(ctx context.Context, pageSize int32, pageToken string) (*pb.ListVrfsResponse, error)
	UpdateVrf(ctx context.Context, name string, updateMask []string, allowMissing bool) (*pb.Vrf, error)

	// SVI Interfaces
	CreateSvi(ctx context.Context, name string, vrf string, logicalBridge string, mac string, gwIPs []string, ebgp bool, remoteAS uint32) (*pb.Svi, error)
	DeleteSvi(ctx context.Context, name string, allowMissing bool) (*emptypb.Empty, error)
	GetSvi(ctx context.Context, name string) (*pb.Svi, error)
	ListSvis(ctx context.Context, pageSize int32, pageToken string) (*pb.ListSvisResponse, error)
	UpdateSvi(ctx context.Context, name string, updateMask []string, allowMissing bool) (*pb.Svi, error)
}

func resourceIDToFullName(container string, resourceID string) string {
	return fmt.Sprintf("//network.opiproject.org/%s/%s", container, resourceID)
}

// NewLogicalBridge creates an evpn Logical Bridge client for use with OPI server at the given address
func NewLogicalBridge(addr string) (EvpnClient, error) {
	c, err := grpcOpi.New(addr)
	if err != nil {
		return nil, err
	}

	// Default is to use the OPI grpc client and the pb client generated from the protobuf spec

	return NewLogicalBridgeWithArgs(c, pb.NewLogicalBridgeServiceClient)
}

// NewLogicalBridgeWithArgs creates an evpn Logical Bridge client for use with OPI server using the given gRPC client and the given function for
// retrieving an evpn protobuf client
func NewLogicalBridgeWithArgs(c grpcOpi.Connector, getter PbEvpnLogicalBridgeClientGetter) (EvpnClient, error) {
	if c == nil {
		return nil, errors.New("grpc connector is nil")
	}

	if getter == nil {
		return nil, errors.New("protobuf client getter is nil")
	}

	return evpnClientImpl{
		getEvpnLogicalBridgeClient: getter,
		Connector:                  c,
	}, nil
}

// NewBridgePort creates an evpn Bridge Port client for use with OPI server at the given address
func NewBridgePort(addr string) (EvpnClient, error) {
	c, err := grpcOpi.New(addr)
	if err != nil {
		return nil, err
	}

	// Default is to use the OPI grpc client and the pb client generated from the protobuf spec

	return NewBridgePortWithArgs(c, pb.NewBridgePortServiceClient)
}

// NewBridgePortWithArgs creates an evpn Bridge Port client for use with OPI server using the given gRPC client and the given function for
// retrieving an evpn protobuf client
func NewBridgePortWithArgs(c grpcOpi.Connector, getter PbEvpnBridgePortClientGetter) (EvpnClient, error) {
	if c == nil {
		return nil, errors.New("grpc connector is nil")
	}

	if getter == nil {
		return nil, errors.New("protobuf client getter is nil")
	}

	return evpnClientImpl{
		getEvpnBridgePortClient: getter,
		Connector:               c,
	}, nil
}

// NewVRF creates an evpn VRF client for use with OPI server at the given address
func NewVRF(addr string) (EvpnClient, error) {
	c, err := grpcOpi.New(addr)
	if err != nil {
		return nil, err
	}

	// Default is to use the OPI grpc client and the pb client generated from the protobuf spec

	return NewVRFWithArgs(c, pb.NewVrfServiceClient)
}

// NewVRFWithArgs creates an evpn VRF client for use with OPI server using the given gRPC client and the given function for
// retrieving an evpn protobuf client
func NewVRFWithArgs(c grpcOpi.Connector, getter PbEvpnVRFClientGetter) (EvpnClient, error) {
	if c == nil {
		return nil, errors.New("grpc connector is nil")
	}

	if getter == nil {
		return nil, errors.New("protobuf client getter is nil")
	}

	return evpnClientImpl{
		getEvpnVRFClient: getter,
		Connector:        c,
	}, nil
}

// NewSVI creates an evpn SVI client for use with OPI server at the given address
func NewSVI(addr string) (EvpnClient, error) {
	c, err := grpcOpi.New(addr)
	if err != nil {
		return nil, err
	}

	// Default is to use the OPI grpc client and the pb client generated from the protobuf spec

	return NewSVIWithArgs(c, pb.NewSviServiceClient)
}

// NewSVIWithArgs creates an evpn SVI client for use with OPI server using the given gRPC client and the given function for
// retrieving an evpn protobuf client
func NewSVIWithArgs(c grpcOpi.Connector, getter PbEvpnSVIClientGetter) (EvpnClient, error) {
	if c == nil {
		return nil, errors.New("grpc connector is nil")
	}

	if getter == nil {
		return nil, errors.New("protobuf client getter is nil")
	}

	return evpnClientImpl{
		getEvpnSVIClient: getter,
		Connector:        c,
	}, nil
}

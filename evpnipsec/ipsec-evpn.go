// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package evpnipsec implements the go library for OPI to be used to establish networking
package evpnipsec

import (
	"context"
	"errors"

	grpcOpi "github.com/opiproject/godpu/grpc"
	pb "github.com/opiproject/opi-evpn-bridge/pkg/ipsec/gen/go"
	"google.golang.org/grpc"
)

// PbEvpnIPSecClientGetter defines the function type used to retrieve an evpn protobuf LogicalBridgeServiceClient
type PbEvpnIPSecClientGetter func(c grpc.ClientConnInterface) pb.IPUIPSecClient

// IPSecEvpnClientImpl defines the function type used to retrieve an evpn protobuf LogicalBridgeServiceClient
type IPSecEvpnClientImpl struct {
	// getIPSecClient PbEvpnIPSecClientGetter
	getIPSecClient PbEvpnIPSecClientGetter

	grpcOpi.Connector
}

// IPSecClient is an interface for querying evpn data from an OPI server
type IPSecClient interface {

	// AddSA interfaces
	AddSA(ctx context.Context, src string, dst string, spi uint32, proto int32, ifID uint32, reqid uint32, mode int32, intrface string, encAlg int32, encKey string,
		intAlg int32, intKey string, replayWindow uint32, tfc uint32, encap int32, esn int32, copyDf int32, copyEcn int32, copyDscp int32, initiator int32, inbound int32,
		update int32) (*pb.AddSAResp, error)
	DelSA(ctx context.Context, src string, dst string, spi uint32, proto int32, ifID uint32) (*pb.DeleteSAResp, error)
}

// NewIPSecClient creates an evpn Logical Bridge client for use with OPI server at the given address
func NewIPSecClient(addr string, tls string) (IPSecClient, error) {
	c, err := grpcOpi.New(addr, tls)
	if err != nil {
		return nil, err
	}

	// Default is to use the OPI grpc client and the pb client generated from the protobuf spec

	return NewIPSecWithArgs(c, pb.NewIPUIPSecClient)
}

// NewIPSecWithArgs creates an evpn Logical Bridge client for use with OPI server using the given gRPC client and the given function for
// retrieving an evpn protobuf client
func NewIPSecWithArgs(c grpcOpi.Connector, getter PbEvpnIPSecClientGetter) (IPSecClient, error) {
	if c == nil {
		return nil, errors.New("grpc connector is nil")
	}

	if getter == nil {
		return nil, errors.New("protobuf client getter is nil")
	}

	return IPSecEvpnClientImpl{
		getIPSecClient: getter,
		Connector:      c,
	}, nil
}

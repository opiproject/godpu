// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	grpcOpi "github.com/opiproject/godpu/grpc"
	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	pc "github.com/opiproject/opi-api/network/opinetcommon/v1alpha1/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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
	CreateLogicalBridge(ctx context.Context, name string, vlanID uint32, vni uint32) (*pb.LogicalBridge, error)
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

// utils

// Function to convert IPv4 address from net.IP to uint32
func ip4ToInt(ip net.IP) uint32 {
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

// Function to parse IP address and prefix from a string of the form "IP/PREFIX"
func parseIPAndPrefix(ipPrefixStr string) (*pc.IPPrefix, error) {
	parts := strings.Split(ipPrefixStr, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid IP address with prefix: %s", ipPrefixStr)
	}

	ipAddress := parts[0]
	prefixLength64, err := strconv.ParseInt(parts[1], 10, 32)

	if err != nil {
		return nil, fmt.Errorf("failed to parse prefix length: %s", err)
	}

	prefixLength := int32(prefixLength64)

	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address: %s", ipAddress)
	}

	addressFamily := int32(0)
	var ipv4Addr uint32
	var ipv6Addr []byte
	var addr *pc.IPAddress
	if ip.To4() != nil {
		addressFamily = 1 // IPv4
		ipv4Addr = ip4ToInt(ip.To4())
		addr = &pc.IPAddress{
			Af: pc.IpAf(addressFamily),
			V4OrV6: &pc.IPAddress_V4Addr{
				V4Addr: ipv4Addr,
			},
		}
	} else {
		addressFamily = 2 // IPv6
		ipv6Addr = ip.To16()
		addr = &pc.IPAddress{
			Af: pc.IpAf(addressFamily),
			V4OrV6: &pc.IPAddress_V6Addr{
				V6Addr: ipv6Addr,
			},
		}
	}
	return &pc.IPPrefix{
		Addr: addr,
		Len:  prefixLength,
	}, nil
}

// Function to parse an array of IP prefixes from strings to pb.IPPrefix messages
func parseIPPrefixes(ipPrefixesStr []string) ([]*pc.IPPrefix, error) {
	const maxPrefixes = 10 // Update this to your expected maximum number of prefixes
	var ipPrefixes = make([]*pc.IPPrefix, 0, maxPrefixes)

	for _, ipPrefixStr := range ipPrefixesStr {
		ipPrefix, err := parseIPAndPrefix(ipPrefixStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse IP prefix: %v", err)
		}
		ipPrefixes = append(ipPrefixes, ipPrefix)
	}
	return ipPrefixes, nil
}

// CreateLogicalBridge creates an Logical Bridge an OPI server
func (c evpnClientImpl) CreateLogicalBridge(ctx context.Context, name string, vlanID uint32, vni uint32) (*pb.LogicalBridge, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnLogicalBridgeClient(conn)

	data, err := client.CreateLogicalBridge(ctx, &pb.CreateLogicalBridgeRequest{
		LogicalBridgeId: name,
		LogicalBridge: &pb.LogicalBridge{
			Spec: &pb.LogicalBridgeSpec{
				VlanId: vlanID,
				Vni:    vni,
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
func (c evpnClientImpl) UpdateLogicalBridge(ctx context.Context, name string, updateMask []string) (*pb.LogicalBridge, error) {
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
	})
	if err != nil {
		log.Printf("error Update logical bridge: %s\n", err)
		return nil, err
	}

	return data, nil
}

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
		typeOfPort = pb.BridgePortType_ACCESS
	case "trunk":
		typeOfPort = pb.BridgePortType_TRUNK
	default:
		typeOfPort = pb.BridgePortType_UNKNOWN
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

// CreateVrf Create vrf on OPI Server
func (c evpnClientImpl) CreateVrf(ctx context.Context, name string, vni uint32, loopback string, vtep string) (*pb.Vrf, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnVRFClient(conn)
	ipLoopback, err := parseIPAndPrefix(loopback)
	if err != nil {
		log.Printf("parseIPAndPrefix: error creating vrf: %s\n", err)
		return nil, err
	}
	ipVtep, err := parseIPAndPrefix(vtep)
	if err != nil {
		log.Printf("parseIPAndPrefix: error creating vrf: %s\n", err)
		return nil, err
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
	client := c.getEvpnVRFClient(conn)
	vrf, err := client.GetVrf(ctx, &pb.GetVrfRequest{
		Name: resourceIDToFullName("vrfs", name),
	})
	if err != nil {
		log.Printf("error updating vrf: %s\n", err)
		return nil, err
	}
	data, err := client.UpdateVrf(ctx, &pb.UpdateVrfRequest{
		Vrf:          vrf,
		UpdateMask:   &fieldmaskpb.FieldMask{Paths: updateMask},
		AllowMissing: allowMissing,
	})
	if err != nil {
		log.Printf("error creating evpn: %s\n", err)
		return nil, err
	}

	return data, nil
}

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
	data, err := client.CreateSvi(ctx, &pb.CreateSviRequest{
		SviId: name,
		Svi: &pb.Svi{
			Spec: &pb.SviSpec{
				Vrf:           vrf,
				LogicalBridge: logicalBridge,
				MacAddress:    []byte(mac),
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
	svi, err := client.GetSvi(ctx, &pb.GetSviRequest{
		Name: resourceIDToFullName("svis", name),
	})
	if err != nil {
		log.Printf("error getting svi: %s\n", err)
		return nil, err
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

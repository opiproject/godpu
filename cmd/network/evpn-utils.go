// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (c) 2024 Ericsson AB.

// Package network implements the network related CLI commands
package network

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/PraserX/ipconv"
	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	pc "github.com/opiproject/opi-api/network/opinetcommon/v1alpha1/gen/go"
)

// ComposeComponentsInfo composes the components with their details
func ComposeComponentsInfo(comp []*pb.Component) string {
	var status string
	for i := 0; i < len(comp); i++ {
		str := fmt.Sprintf(" %+v\n", comp[i])
		status += str
	}
	return status
}

// ComposeGwIps compose the gw ips string
func ComposeGwIps(comp []*pc.IPPrefix) string {
	var status string
	for i := 0; i < len(comp); i++ {
		str := fmt.Sprintf("%+v/%v ", ipconv.IntToIPv4(comp[i].GetAddr().GetV4Addr()), comp[i].GetLen())
		status += str
	}
	return status
}

// ExtractShortName takes a full name and returns the short name.
func ExtractShortName(fullName string) string {
	parts := strings.Split(fullName, "/")
	return parts[len(parts)-1] // Return the last part of the split name.
}

// PrintLB prints the logical bridge fields in human readable format
func PrintLB(lb *pb.LogicalBridge) {
	shortName := ExtractShortName(lb.GetName())
	log.Println("name:", shortName)

	log.Println("status:", lb.GetStatus().GetOperStatus().String())
	log.Println("vlan:", lb.GetSpec().GetVlanId())
	if lb.GetSpec().GetVni() != 0 {
		log.Println("vni:", lb.GetSpec().GetVni())
	}
	if lb.GetSpec().GetVtepIpPrefix().GetAddr().GetV4Addr() != 0 {
		Vteip := fmt.Sprintf("%+v/%v", ipconv.IntToIPv4(lb.GetSpec().GetVtepIpPrefix().GetAddr().GetV4Addr()), lb.GetSpec().GetVtepIpPrefix().GetLen())
		log.Println("vtep ip:", Vteip)
	}

	log.Println("Component Status:")
	log.Println(ComposeComponentsInfo(lb.GetStatus().GetComponents()))
}

// PrintBP prints the bridge Port fields in human readable format
func PrintBP(bp *pb.BridgePort) {
	shortName := ExtractShortName(bp.GetName())
	log.Println("name:", shortName)

	log.Println("status:", bp.GetStatus().GetOperStatus().String())
	log.Println("ptype:", bp.GetSpec().GetPtype())
	log.Println("MacAddress:", net.HardwareAddr(bp.GetSpec().GetMacAddress()).String())

	// Extract short names for the logical bridges
	bridges := bp.GetSpec().GetLogicalBridges()
	shortBridgeNames := make([]string, len(bridges))
	for i, bridge := range bridges {
		shortBridgeNames[i] = ExtractShortName(bridge)
	}
	log.Println("bridges:", shortBridgeNames)

	log.Println("Component Status:")
	log.Println(ComposeComponentsInfo(bp.GetStatus().GetComponents()))
}

// PrintSvi prints the svi fields in human readable format
func PrintSvi(svi *pb.Svi) {
	shortName := ExtractShortName(svi.GetName())
	log.Println("name:", shortName)

	log.Println("status:", svi.GetStatus().GetOperStatus().String())

	shortName = ExtractShortName(svi.GetSpec().GetVrf())
	log.Println("Vrf:", shortName)

	shortName = ExtractShortName(svi.GetSpec().GetLogicalBridge())
	log.Println("LogicalBridge:", shortName)
	log.Println("MacAddress:", net.HardwareAddr(svi.GetSpec().GetMacAddress()).String())
	log.Println("GwIPs:", ComposeGwIps(svi.GetSpec().GetGwIpPrefix()))
	if svi.GetSpec().GetRemoteAs() != 0 {
		log.Println("remoteAS:", svi.GetSpec().GetRemoteAs())
	}
	if svi.GetSpec().GetEnableBgp() {
		log.Println("EnableBgp:", svi.GetSpec().GetEnableBgp())
	}
	log.Println("Component Status:")
	log.Println(ComposeComponentsInfo(svi.GetStatus().GetComponents()))
}

// PrintVrf prints the vrf fields in human readable format
func PrintVrf(vrf *pb.Vrf) {
	Loopback := fmt.Sprintf("%+v/%+v", ipconv.IntToIPv4(vrf.GetSpec().GetLoopbackIpPrefix().GetAddr().GetV4Addr()), vrf.GetSpec().GetLoopbackIpPrefix().GetLen())

	shortName := ExtractShortName(vrf.GetName())
	log.Println("name:", shortName)

	log.Println("operation status:", vrf.GetStatus().GetOperStatus().String())

	if vrf.GetSpec().GetVni() != 0 {
		log.Println("vni:", vrf.GetSpec().GetVni())
	}
	if vrf.GetSpec().GetVtepIpPrefix().GetAddr().GetV4Addr() != 0 {
		Vteip := fmt.Sprintf("%+v/%v", ipconv.IntToIPv4(vrf.GetSpec().GetVtepIpPrefix().GetAddr().GetV4Addr()), vrf.GetSpec().GetVtepIpPrefix().GetLen())
		log.Println("vtep ip:", Vteip)
	}
	log.Println("loopback ip:", Loopback)
	log.Println("Component Status:")
	log.Println(ComposeComponentsInfo(vrf.GetStatus().GetComponents()))
}

// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (c) 2024 Ericsson AB.

// Package network implements the network related CLI commands
package network

import (
	"fmt"

	"github.com/PraserX/ipconv"
	pb "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	pc "github.com/opiproject/opi-api/network/opinetcommon/v1alpha1/gen/go"
)

// PrintComponents prints the components with their details
func PrintComponents(comp []*pb.Component) string {
	var status string
	for i := 0; i < len(comp); i++ {
		str := fmt.Sprintf(" %+v\n", comp[i])
		status += str
	}
	return status
}

// PrintGWIPs prints the gw ips
func PrintGWIPs(comp []*pc.IPPrefix) string {
	var status string
	for i := 0; i < len(comp); i++ {
		str := fmt.Sprintf("%+v/%v ", ipconv.IntToIPv4(comp[i].GetAddr().GetV4Addr()), comp[i].GetLen())
		status += str
	}
	return status
}

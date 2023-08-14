// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"fmt"
	"net"

	pc "github.com/opiproject/opi-api/network/opinetcommon/v1alpha1/gen/go"
)

// Function to convert IPv4 address from net.IP to uint32
func ip4ToInt(ip net.IP) uint32 {
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}

// Function to parse IP address and prefix from a string of the form "IP/PREFIX"
func parseIPAndPrefix(ipPrefixStr string) (*pc.IPPrefix, error) {
	ip, ipnet, err := net.ParseCIDR(ipPrefixStr)
	if err != nil {
		return nil, err
	}

	prefixLength, _ := ipnet.Mask.Size()
	addr := &pc.IPAddress{}

	if ip.To4() != nil {
		addr.Af = pc.IpAf_IP_AF_INET
		addr.V4OrV6 = &pc.IPAddress_V4Addr{
			V4Addr: ip4ToInt(ip.To4()),
		}
	} else {
		addr.Af = pc.IpAf_IP_AF_INET6
		addr.V4OrV6 = &pc.IPAddress_V6Addr{
			V6Addr: ip.To16(),
		}
	}

	return &pc.IPPrefix{
		Addr: addr,
		Len:  int32(prefixLength),
	}, nil
}

// Function to parse an array of IP prefixes from strings to pb.IPPrefix messages
func parseIPPrefixes(ipPrefixesStr []string) ([]*pc.IPPrefix, error) {
	ipPrefixes := make([]*pc.IPPrefix, len(ipPrefixesStr))

	for i, ipPrefixStr := range ipPrefixesStr {
		ipPrefix, err := parseIPAndPrefix(ipPrefixStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse IP prefix: %v", err)
		}
		ipPrefixes[i] = ipPrefix
	}

	return ipPrefixes, nil
}

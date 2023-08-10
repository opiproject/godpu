// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	pc "github.com/opiproject/opi-api/network/opinetcommon/v1alpha1/gen/go"
)

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

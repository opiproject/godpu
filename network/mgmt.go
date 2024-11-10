// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"context"
	"log"

	pm "github.com/opiproject/opi-evpn-bridge/pkg/netlink/proto/gen/go"
)

// DumpNetlinkDb get netlink DB details from OPI server
func (c evpnClientImpl) DumpNetlinkDB(ctx context.Context, details bool) (*pm.DumpNetlinkDbResult, error) {
	conn, closer, err := c.NewConn()
	if err != nil {
		log.Printf("error creating connection: %s\n", err)
		return nil, err
	}
	defer closer()

	client := c.getEvpnMgmtClient(conn)
	data, err := client.DumpNetlinkDB(ctx, &pm.DumpNetlinkDbRequest{
		Details: details,
	})
	if err != nil {
		log.Printf("error getting vrf: %s\n", err)
		return nil, err
	}

	return data, nil
}

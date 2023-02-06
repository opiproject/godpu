// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

// Package main implements the CLI commands
package main

import (
	"github.com/opiproject/godpu/cmd/inventory"
	"github.com/opiproject/godpu/cmd/ipsec"
)

func main() {
	ipsec.Execute()
	inventory.Execute()
}

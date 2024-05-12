// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package main implements the CLI commands
package main

import (
	"log"

	"github.com/opiproject/godpu/cmd"
)

func main() {
	command := cmd.NewCommand()
	if err := command.Execute(); err != nil {
		log.Fatalf("[ERROR] %s", err.Error())
	}
}

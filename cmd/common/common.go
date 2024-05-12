// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023-2024 Intel Corporation

// Package common has common constants, functions for all storage commands
package common

import (
	"fmt"
	"os"
)

// AddrCmdLineArg cmdline arg name for address
const AddrCmdLineArg = "addr"

// TimeoutCmdLineArg cmdline arg name for timeout
const TimeoutCmdLineArg = "timeout"

// TLSFiles cmdline arg name for tls files
const TLSFiles = "tlsfiles"

// PrintResponse prints only response string into stdout without any
// additional information
func PrintResponse(response string) {
	fmt.Fprintln(os.Stdout, response)
}

// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package cmd implements the CLI commands
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const addrCmdLineArg = "addr"
const timeoutCmdLineArg = "timeout"

// NewStorageCommand tests the storage functionality
func NewStorageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "storage",
		Aliases: []string{"g"},
		Short:   "Tests storage functionality",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	flags := cmd.PersistentFlags()
	flags.String(addrCmdLineArg, "localhost:50151", "address of OPI gRPC server")
	flags.Duration(timeoutCmdLineArg, 10*time.Second, "timeout for a cmd")

	cmd.AddCommand(newStorageCreateCommand())
	cmd.AddCommand(newStorageDeleteCommand())
	cmd.AddCommand(newStorageTestCommand())

	return cmd
}

func newStorageCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Creates resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newCreateNvmeCommand())

	return cmd
}

func newCreateNvmeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nvme",
		Aliases: []string{"n"},
		Short:   "Creates nvme resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newCreateNvmeSubsystemCommand())
	cmd.AddCommand(newCreateNvmeNamespaceCommand())

	return cmd
}

func newStorageDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"d"},
		Short:   "Deletes resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newDeleteNvmeCommand())

	return cmd
}

func newDeleteNvmeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nvme",
		Aliases: []string{"n"},
		Short:   "Deletes nvme resource",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			err := c.Help()
			cobra.CheckErr(err)
		},
	}

	cmd.AddCommand(newDeleteNvmeSubsystemCommand())

	return cmd
}

func printResponse(response string) {
	fmt.Fprintln(os.Stdout, response)
}

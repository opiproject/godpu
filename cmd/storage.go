// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package cmd implements the CLI commands
package cmd

import (
	"github.com/spf13/cobra"
)

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

	cmd.AddCommand(newStorageTestCommand())

	return cmd
}

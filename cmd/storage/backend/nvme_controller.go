// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2024 Intel Corporation

// Package backend implements the CLI commands for storage backend
package backend

import (
	"context"
	"fmt"
	"strings"

	"github.com/opiproject/godpu/cmd/storage/common"
	backendclient "github.com/opiproject/godpu/storage/backend"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"github.com/spf13/cobra"
)

func newCreateNvmeControllerCommand() *cobra.Command {
	id := ""
	multipath := ""
	cmd := &cobra.Command{
		Use:     "controller",
		Aliases: []string{"c"},
		Short:   "Creates nvme controller representing an external nvme device",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(common.TimeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := backendclient.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			allowedModes := map[string]pb.NvmeMultipath{
				"disable":   pb.NvmeMultipath_NVME_MULTIPATH_DISABLE,
				"failover":  pb.NvmeMultipath_NVME_MULTIPATH_FAILOVER,
				"multipath": pb.NvmeMultipath_NVME_MULTIPATH_MULTIPATH,
			}

			mode, ok := allowedModes[strings.ToLower(multipath)]
			if !ok {
				cobra.CheckErr(fmt.Errorf("not allowed multipath mode: '%s'", multipath))
			}

			response, err := client.CreateNvmeController(ctx, id, mode)
			cobra.CheckErr(err)

			common.PrintResponse(response.Name)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "id for created resource. Assigned by server if omitted.")
	cmd.Flags().StringVar(&multipath, "multipath", "disable", "multipath mode (disable, failover, enable)")

	return cmd
}

func newDeleteNvmeControllerCommand() *cobra.Command {
	name := ""
	allowMissing := false
	cmd := &cobra.Command{
		Use:     "controller",
		Aliases: []string{"c"},
		Short:   "Deletes nvme controller representing an external nvme device",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			addr, err := c.Flags().GetString(common.AddrCmdLineArg)
			cobra.CheckErr(err)

			timeout, err := c.Flags().GetDuration(common.TimeoutCmdLineArg)
			cobra.CheckErr(err)

			client, err := backendclient.New(addr)
			cobra.CheckErr(err)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			err = client.DeleteNvmeController(ctx, name, allowMissing)
			cobra.CheckErr(err)
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "name of deleted remote controller")
	cmd.Flags().BoolVar(&allowMissing, "allowMissing", false, "cmd succeeds if attempts to delete a resource that is not present")

	cobra.CheckErr(cmd.MarkFlagRequired("name"))

	return cmd
}

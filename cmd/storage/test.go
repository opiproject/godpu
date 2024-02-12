// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (C) 2023-2024 Intel Corporation

// Package storage implements the storage related CLI commands
package storage

import (
	"context"
	"log"

	"github.com/opiproject/godpu/cmd/storage/common"
	"github.com/opiproject/godpu/grpc"
	"github.com/opiproject/godpu/storage/test"
	"github.com/spf13/cobra"
)

type storagePartition string

const (
	storagePartitionFrontend  storagePartition = "frontend"
	storagePartitionBackend   storagePartition = "backend"
	storagePartitionMiddleend storagePartition = "middleend"
)

var allStoragePartitions = []storagePartition{
	storagePartitionFrontend,
	storagePartitionBackend,
	storagePartitionMiddleend,
}

func newStorageTestCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "test",
		Aliases: []string{"s"},
		Short:   "Test storage functionality",
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			runTests(
				c,
				allStoragePartitions,
				test.AllFrontendPartitions,
			)
		},
	}

	cmd.AddCommand(newStorageTestFrontendCommand())
	cmd.AddCommand(newStorageTestBackendCommand())
	cmd.AddCommand(newStorageTestMiddleendCommand())

	return cmd
}

func newStorageTestFrontendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   string(storagePartitionFrontend),
		Short: "Tests storage frontend API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionFrontend},
				test.AllFrontendPartitions,
			)
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "nvme",
		Short: "Tests storage frontend nvme API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionFrontend},
				[]test.FrontendPartition{test.FrontendPartitionNvme},
			)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "virtio-blk",
		Short: "Tests storage frontend virtio-blk API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionFrontend},
				[]test.FrontendPartition{test.FrontendPartitionVirtioBlk},
			)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "scsi",
		Short: "Tests storage frontend scsi API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionFrontend},
				[]test.FrontendPartition{test.FrontendPartitionScsi},
			)
		},
	})

	return cmd
}

func newStorageTestBackendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   string(storagePartitionBackend),
		Short: "Tests storage backend API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionBackend},
				nil,
			)
		},
	}

	return cmd
}

func newStorageTestMiddleendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   string(storagePartitionMiddleend),
		Short: "Tests storage middleend API",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, _ []string) {
			runTests(
				c,
				[]storagePartition{storagePartitionMiddleend},
				nil,
			)
		},
	}

	return cmd
}

func runTests(
	cmd *cobra.Command,
	partitions []storagePartition,
	frontendPartitions []test.FrontendPartition,
) {
	addr, err := cmd.Flags().GetString(common.AddrCmdLineArg)
	if err != nil {
		log.Fatalf("error getting %v argument: %v", common.AddrCmdLineArg, err)
	}

	timeout, err := cmd.Flags().GetDuration(common.TimeoutCmdLineArg)
	if err != nil {
		log.Fatalf("error getting %v argument: %v", common.AddrCmdLineArg, err)
	}

	// Set up a connection to the server.
	client, err := grpc.New(addr)
	if err != nil {
		log.Fatalf("error creating new client: %v", err)
	}

	// Contact the server and print out its response.
	conn, closer, err := client.NewConn()
	if err != nil {
		log.Fatalf("error creating gRPC connection: %v", err)
	}
	defer closer()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for _, partition := range partitions {
		log.Printf("==============================================================================")
		log.Printf("Test %v", partition)
		log.Printf("==============================================================================")

		var err error
		switch partition {
		case storagePartitionFrontend:
			err = test.DoFrontend(ctx, conn, frontendPartitions)
		case storagePartitionBackend:
			err = test.DoBackend(ctx, conn)
		case storagePartitionMiddleend:
			err = test.DoMiddleend(ctx, conn)
		default:
			log.Panicf("Unknown storage partition: %v", partition)
		}

		if err != nil {
			log.Panicf("%v tests failed with error: %v", partition, err)
		}
	}
}

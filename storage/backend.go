// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"
	"fmt"
	"log"
	"net"
	"path"

	"github.com/google/uuid"
	pc "github.com/opiproject/opi-api/common/v1/gen/go"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/grpc"
)

// DoBackend executes the back end code
func DoBackend(ctx context.Context, conn grpc.ClientConnInterface) error {
	nvme := pb.NewNVMfRemoteControllerServiceClient(conn)
	null := pb.NewNullDebugServiceClient(conn)
	aio := pb.NewAioControllerServiceClient(conn)

	err := executeNVMfRemoteController(ctx, nvme)
	if err != nil {
		return err
	}
	err = executeNullDebug(ctx, null)
	if err != nil {
		return err
	}
	err = executeAioController(ctx, aio)
	if err != nil {
		return err
	}
	return nil
}

func executeNVMfRemoteController(ctx context.Context, c4 pb.NVMfRemoteControllerServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewNVMfRemoteControllerServiceClient")
	log.Printf("=======================================")

	addr, err := net.LookupIP("spdk")
	if err != nil {
		return err
	}

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"opi-nvme8"} {
		rr0, err := c4.CreateNVMfRemoteController(ctx, &pb.CreateNVMfRemoteControllerRequest{
			NvMfRemoteControllerId: resourceID,
			NvMfRemoteController: &pb.NVMfRemoteController{
				Trtype:  pb.NvmeTransportType_NVME_TRANSPORT_TCP,
				Adrfam:  pb.NvmeAddressFamily_NVMF_ADRFAM_IPV4,
				Traddr:  addr[0].String(),
				Trsvcid: 4444,
				Subnqn:  "nqn.2016-06.io.spdk:cnode1",
				Hostnqn: "nqn.2014-08.org.nvmexpress:uuid:feb98abe-d51f-40c8-b348-2753f3571d3c"}})
		if err != nil {
			return err
		}
		// verify
		newResourceID := resourceID
		if resourceID == "" {
			parsed, err := uuid.Parse(path.Base(rr0.Name))
			if err != nil {
				return err
			}
			newResourceID = parsed.String()
		}
		fullname := fmt.Sprintf("//storage.opiproject.org/volumes/%s", newResourceID)
		if rr0.Name != fullname {
			return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rr0.Name, fullname)
		}
		log.Printf("Connected NVMf: %v", rr0)
		// continue
		rr2, err := c4.NVMfRemoteControllerReset(ctx, &pb.NVMfRemoteControllerResetRequest{Id: &pc.ObjectKey{Value: rr0.Name}})
		if err != nil {
			return err
		}
		log.Printf("Reset NVMf: %v", rr2)
		rr3, err := c4.ListNVMfRemoteControllers(ctx, &pb.ListNVMfRemoteControllersRequest{})
		if err != nil {
			return err
		}
		log.Printf("List NVMf: %v", rr3)
		rr4, err := c4.GetNVMfRemoteController(ctx, &pb.GetNVMfRemoteControllerRequest{Name: rr0.Name})
		if err != nil {
			return err
		}
		log.Printf("Got NVMf: %v", rr4)
		rr5, err := c4.NVMfRemoteControllerStats(ctx, &pb.NVMfRemoteControllerStatsRequest{Id: &pc.ObjectKey{Value: rr0.Name}})
		if err != nil {
			return err
		}
		log.Printf("Stats NVMf: %v", rr5)
		rr1, err := c4.DeleteNVMfRemoteController(ctx, &pb.DeleteNVMfRemoteControllerRequest{Name: rr0.Name})
		if err != nil {
			return err
		}
		log.Printf("Disconnected NVMf: %v -> %v", rr0, rr1)
	}
	return nil
}

func executeNullDebug(ctx context.Context, c1 pb.NullDebugServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewNullDebugServiceClient")
	log.Printf("=======================================")

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"opi-null9", ""} {
		rs1, err := c1.CreateNullDebug(ctx, &pb.CreateNullDebugRequest{
			NullDebugId: resourceID,
			NullDebug:   &pb.NullDebug{BlockSize: 512, BlocksCount: 64}})
		if err != nil {
			return err
		}
		// verify
		newResourceID := resourceID
		if resourceID == "" {
			parsed, err := uuid.Parse(path.Base(rs1.Name))
			if err != nil {
				return err
			}
			newResourceID = parsed.String()
		}
		fullname := fmt.Sprintf("//storage.opiproject.org/volumes/%s", newResourceID)
		if rs1.Name != fullname {
			return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rs1.Name, fullname)
		}
		log.Printf("Added Null: %v", rs1)
		// continue
		rs3, err := c1.UpdateNullDebug(ctx, &pb.UpdateNullDebugRequest{NullDebug: &pb.NullDebug{Name: rs1.Name}})
		if err != nil {
			return err
		}
		log.Printf("Updated Null: %v", rs3)
		rs4, err := c1.ListNullDebugs(ctx, &pb.ListNullDebugsRequest{})
		if err != nil {
			return err
		}
		log.Printf("Listed Null: %v", rs4)
		rs5, err := c1.GetNullDebug(ctx, &pb.GetNullDebugRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got Null: %s", rs5.Name)
		rs6, err := c1.NullDebugStats(ctx, &pb.NullDebugStatsRequest{Handle: &pc.ObjectKey{Value: rs1.Name}})
		if err != nil {
			return err
		}
		log.Printf("Stats Null: %s", rs6.Stats)
		rs2, err := c1.DeleteNullDebug(ctx, &pb.DeleteNullDebugRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted Null: %v -> %v", rs1, rs2)
	}
	return nil
}

func executeAioController(ctx context.Context, c2 pb.AioControllerServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewAioControllerServiceClient")
	log.Printf("=======================================")

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"opi-aio4", ""} {
		ra1, err := c2.CreateAioController(ctx, &pb.CreateAioControllerRequest{
			AioControllerId: resourceID,
			AioController:   &pb.AioController{BlockSize: 512, BlocksCount: 12, Filename: "/tmp/aio_bdev_file"}})
		if err != nil {
			return err
		}
		// verify
		newResourceID := resourceID
		if resourceID == "" {
			parsed, err := uuid.Parse(path.Base(ra1.Name))
			if err != nil {
				return err
			}
			newResourceID = parsed.String()
		}
		fullname := fmt.Sprintf("//storage.opiproject.org/volumes/%s", newResourceID)
		if ra1.Name != fullname {
			return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", ra1.Name, fullname)
		}
		log.Printf("Added Aio: %v", ra1)
		// continue
		ra3, err := c2.UpdateAioController(ctx, &pb.UpdateAioControllerRequest{AioController: &pb.AioController{Name: ra1.Name, Filename: "/tmp/aio_bdev_file"}})
		if err != nil {
			return err
		}
		log.Printf("Updated Aio: %v", ra3)
		ra4, err := c2.ListAioControllers(ctx, &pb.ListAioControllersRequest{})
		if err != nil {
			return err
		}
		log.Printf("Listed Aio: %v", ra4)
		ra5, err := c2.GetAioController(ctx, &pb.GetAioControllerRequest{Name: ra1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got Aio: %s", ra5.Name)
		ra6, err := c2.AioControllerStats(ctx, &pb.AioControllerStatsRequest{Handle: &pc.ObjectKey{Value: ra1.Name}})
		if err != nil {
			return err
		}
		log.Printf("Stats Aio: %s", ra6.Stats)
		ra2, err := c2.DeleteAioController(ctx, &pb.DeleteAioControllerRequest{Name: ra1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted Aio: %v -> %v", ra1, ra2)
	}
	return nil
}

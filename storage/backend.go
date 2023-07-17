// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
// Copyright (C) 2023 Intel Corporation

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"
	"fmt"
	"log"
	"net"
	"path"
	"time"

	"github.com/google/uuid"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// DoBackend executes the back end code
func DoBackend(ctx context.Context, conn grpc.ClientConnInterface) error {
	nvme := pb.NewNvmeRemoteControllerServiceClient(conn)
	null := pb.NewNullVolumeServiceClient(conn)
	aio := pb.NewAioVolumeServiceClient(conn)

	err := executeNvmeRemoteController(ctx, nvme)
	if err != nil {
		return err
	}
	err = executeNvmePath(ctx, nvme, false)
	if err != nil {
		return err
	}
	err = executeNvmePath(ctx, nvme, true)
	if err != nil {
		return err
	}
	err = executeNullVolume(ctx, null)
	if err != nil {
		return err
	}
	err = executeAioVolume(ctx, aio)
	if err != nil {
		return err
	}
	return nil
}

func executeNvmeRemoteController(ctx context.Context, c4 pb.NvmeRemoteControllerServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewNvmeRemoteControllerServiceClient")
	log.Printf("=======================================")

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"opi-nvme8", ""} {
		rr0, err := c4.CreateNvmeRemoteController(ctx, &pb.CreateNvmeRemoteControllerRequest{
			NvmeRemoteControllerId: resourceID,
			NvmeRemoteController: &pb.NvmeRemoteController{
				Multipath: pb.NvmeMultipath_NVME_MULTIPATH_MULTIPATH,
				Tcp: &pb.TcpController{
					Hdgst: false,
					Ddgst: false,
				},
			}})
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
		log.Printf("Created Nvme controller: %v", rr0)
		// continue
		rr2, err := c4.ResetNvmeRemoteController(ctx, &pb.ResetNvmeRemoteControllerRequest{Name: rr0.Name})
		if err != nil {
			return err
		}
		log.Printf("Reset Nvme: %v", rr2)
		rr3, err := c4.ListNvmeRemoteControllers(ctx, &pb.ListNvmeRemoteControllersRequest{Parent: "todo"})
		if err != nil {
			return err
		}
		log.Printf("List Nvme: %v", rr3)
		rr4, err := c4.GetNvmeRemoteController(ctx, &pb.GetNvmeRemoteControllerRequest{Name: rr0.Name})
		if err != nil {
			return err
		}
		log.Printf("Got Nvme: %v", rr4)
		rr5, err := c4.StatsNvmeRemoteController(ctx, &pb.StatsNvmeRemoteControllerRequest{Name: rr0.Name})
		if err != nil {
			return err
		}
		log.Printf("Stats Nvme: %v", rr5)
		rr1, err := c4.DeleteNvmeRemoteController(ctx, &pb.DeleteNvmeRemoteControllerRequest{Name: rr0.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted Nvme controller: %v -> %v", rr0, rr1)

		// wait for some time for the backend to delete above objects
		time.Sleep(time.Second)
	}
	return nil
}

func executeNvmePath(ctx context.Context, c5 pb.NvmeRemoteControllerServiceClient, tlsEnabled bool) error {
	log.Printf("=======================================")
	log.Printf("Testing NewNvmePathClient TLS=%v", tlsEnabled)
	log.Printf("=======================================")

	addr, err := net.LookupIP("spdk")
	if err != nil {
		return err
	}

	port := 4444
	psk := []byte{}
	if tlsEnabled {
		port = 5555
		psk = []byte("NVMeTLSkey-1:01:MDAxMTIyMzM0NDU1NjY3Nzg4OTlhYWJiY2NkZGVlZmZwJEiQ:")
	}

	ctrlrResourceID := "opi-nvme8"
	rr0, err := c5.CreateNvmeRemoteController(ctx, &pb.CreateNvmeRemoteControllerRequest{
		NvmeRemoteControllerId: ctrlrResourceID,
		NvmeRemoteController: &pb.NvmeRemoteController{
			Multipath: pb.NvmeMultipath_NVME_MULTIPATH_MULTIPATH,
			Tcp: &pb.TcpController{
				Hdgst: false,
				Ddgst: false,
				Psk:   psk,
			},
		}})
	if err != nil {
		return err
	}
	log.Printf("Created Nvme controller: %v", rr0)

	for _, resourceID := range []string{"opi-nvme8-path", ""} {
		np0, err := c5.CreateNvmePath(ctx, &pb.CreateNvmePathRequest{
			NvmePathId: resourceID,
			NvmePath: &pb.NvmePath{
				Trtype:            pb.NvmeTransportType_NVME_TRANSPORT_TCP,
				Traddr:            addr[0].String(),
				ControllerNameRef: rr0.Name,
				Fabrics: &pb.FabricsPath{
					Adrfam:  pb.NvmeAddressFamily_NVME_ADRFAM_IPV4,
					Trsvcid: int64(port),
					Subnqn:  "nqn.2016-06.io.spdk:cnode1",
					Hostnqn: "nqn.2014-08.org.nvmexpress:uuid:feb98abe-d51f-40c8-b348-2753f3571d3c",
				},
			}})
		if err != nil {
			return err
		}
		log.Printf("Created Nvme path: %v", np0)

		newResourceID := resourceID
		if resourceID == "" {
			parsed, err := uuid.Parse(path.Base(np0.Name))
			if err != nil {
				return err
			}
			newResourceID = parsed.String()
		}
		fullname := fmt.Sprintf("//storage.opiproject.org/volumes/%s", newResourceID)
		if np0.Name != fullname {
			return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", np0.Name, fullname)
		}
		log.Printf("Created Nvme path: %v", np0)
		np3, err := c5.UpdateNvmePath(ctx, &pb.UpdateNvmePathRequest{
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			NvmePath: &pb.NvmePath{
				Name:              np0.Name,
				Trtype:            pb.NvmeTransportType_NVME_TRANSPORT_TCP,
				Traddr:            addr[0].String(),
				ControllerNameRef: rr0.Name,
				Fabrics: &pb.FabricsPath{
					Adrfam:  pb.NvmeAddressFamily_NVME_ADRFAM_IPV4,
					Trsvcid: int64(port),
					Subnqn:  "nqn.2016-06.io.spdk:cnode1",
					Hostnqn: "nqn.2014-08.org.nvmexpress:uuid:feb98abe-d51f-40c8-b348-2753f3571d3c",
				},
			}})
		if err != nil {
			return err
		}
		log.Printf("Updated Nvme path: %v", np3)
		np4, err := c5.ListNvmePaths(ctx, &pb.ListNvmePathsRequest{Parent: "todo"})
		if err != nil {
			return err
		}
		log.Printf("Listed Nvme path: %v", np4)
		np5, err := c5.GetNvmePath(ctx, &pb.GetNvmePathRequest{Name: np0.Name})
		if err != nil {
			return err
		}
		log.Printf("Got Nvme path: %s", np5.Name)
		np6, err := c5.StatsNvmePath(ctx, &pb.StatsNvmePathRequest{Name: np0.Name})
		if err != nil {
			return err
		}
		log.Printf("Stats Nvme path: %s", np6.Stats)
		np1, err := c5.DeleteNvmePath(ctx, &pb.DeleteNvmePathRequest{
			Name: np0.Name,
		})
		if err != nil {
			return err
		}
		log.Printf("Deleted Nvme path: %v -> %v", np0, np1)

		// wait for some time for the backend to delete above objects
		time.Sleep(time.Second)
	}

	rr1, err := c5.DeleteNvmeRemoteController(ctx, &pb.DeleteNvmeRemoteControllerRequest{Name: rr0.Name})
	if err != nil {
		return err
	}
	log.Printf("Deleted Nvme controller: %v -> %v", rr0, rr1)

	return nil
}

func executeNullVolume(ctx context.Context, c1 pb.NullVolumeServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewNullVolumeServiceClient")
	log.Printf("=======================================")

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"opi-null9", ""} {
		rs1, err := c1.CreateNullVolume(ctx, &pb.CreateNullVolumeRequest{
			NullVolumeId: resourceID,
			NullVolume:   &pb.NullVolume{BlockSize: 512, BlocksCount: 64}})
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
		rs3, err := c1.UpdateNullVolume(ctx, &pb.UpdateNullVolumeRequest{
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			NullVolume: &pb.NullVolume{
				Name:        rs1.Name,
				BlockSize:   512,
				BlocksCount: 128,
			}})
		if err != nil {
			return err
		}
		log.Printf("Updated Null: %v", rs3)
		rs4, err := c1.ListNullVolumes(ctx, &pb.ListNullVolumesRequest{Parent: "todo"})
		if err != nil {
			return err
		}
		log.Printf("Listed Null: %v", rs4)
		rs5, err := c1.GetNullVolume(ctx, &pb.GetNullVolumeRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got Null: %s", rs5.Name)
		rs6, err := c1.StatsNullVolume(ctx, &pb.StatsNullVolumeRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Stats Null: %s", rs6.Stats)
		rs2, err := c1.DeleteNullVolume(ctx, &pb.DeleteNullVolumeRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted Null: %v -> %v", rs1, rs2)
	}
	return nil
}

func executeAioVolume(ctx context.Context, c2 pb.AioVolumeServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewAioVolumeServiceClient")
	log.Printf("=======================================")

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"opi-aio4", ""} {
		ra1, err := c2.CreateAioVolume(ctx, &pb.CreateAioVolumeRequest{
			AioVolumeId: resourceID,
			AioVolume:   &pb.AioVolume{BlockSize: 512, BlocksCount: 12, Filename: "/tmp/aio_bdev_file"}})
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
		ra3, err := c2.UpdateAioVolume(ctx, &pb.UpdateAioVolumeRequest{
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			AioVolume:  &pb.AioVolume{Name: ra1.Name, Filename: "/tmp/aio_bdev_file"}})
		if err != nil {
			return err
		}
		log.Printf("Updated Aio: %v", ra3)
		ra4, err := c2.ListAioVolumes(ctx, &pb.ListAioVolumesRequest{Parent: "todo"})
		if err != nil {
			return err
		}
		log.Printf("Listed Aio: %v", ra4)
		ra5, err := c2.GetAioVolume(ctx, &pb.GetAioVolumeRequest{Name: ra1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got Aio: %s", ra5.Name)
		ra6, err := c2.StatsAioVolume(ctx, &pb.StatsAioVolumeRequest{Name: ra1.Name})
		if err != nil {
			return err
		}
		log.Printf("Stats Aio: %s", ra6.Stats)
		ra2, err := c2.DeleteAioVolume(ctx, &pb.DeleteAioVolumeRequest{Name: ra1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted Aio: %v -> %v", ra1, ra2)
	}
	return nil
}

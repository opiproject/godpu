// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"
	"log"
	"time"

	pbc "github.com/opiproject/opi-api/common/v1/gen/go"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/grpc"
)

// DoFrontend executes the front end code
func DoFrontend(ctx context.Context, conn grpc.ClientConnInterface) error {
	nvme := pb.NewFrontendNvmeServiceClient(conn)
	blk := pb.NewFrontendVirtioBlkServiceClient(conn)
	scsi := pb.NewFrontendVirtioScsiServiceClient(conn)

	err := executeNVMeSubsystem(ctx, nvme)
	if err != nil {
		return err
	}
	err = executeNVMeController(ctx, nvme)
	if err != nil {
		return err
	}
	err = executeNVMeNamespace(ctx, nvme)
	if err != nil {
		return err
	}
	err = executeVirtioBlk(ctx, blk)
	if err != nil {
		return err
	}
	err = executeVirtioScsiController(ctx, scsi)
	if err != nil {
		return err
	}
	err = executeVirtioScsiLun(ctx, scsi)
	if err != nil {
		return err
	}
	return nil
}

func executeVirtioScsiLun(ctx context.Context, c6 pb.FrontendVirtioScsiServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing VirtioScsiLun")
	log.Printf("=======================================")
	// pre create: controller
	rss1, err := c6.CreateVirtioScsiController(ctx, &pb.CreateVirtioScsiControllerRequest{VirtioScsiController: &pb.VirtioScsiController{Id: &pbc.ObjectKey{Value: "OPI-VirtioScsi8"}}})
	if err != nil {
		return err
	}
	rl1, err := c6.CreateVirtioScsiLun(ctx, &pb.CreateVirtioScsiLunRequest{VirtioScsiLun: &pb.VirtioScsiLun{Id: &pbc.ObjectKey{Value: "OPI-VirtioScsi8"}, TargetId: &pbc.ObjectKey{Value: "OPI-VirtioScsi8"}, VolumeId: &pbc.ObjectKey{Value: "Malloc1"}}})
	if err != nil {
		return err
	}
	log.Printf("Added VirtioScsiLun: %v", rl1)
	rl3, err := c6.UpdateVirtioScsiLun(ctx, &pb.UpdateVirtioScsiLunRequest{VirtioScsiLun: &pb.VirtioScsiLun{Id: &pbc.ObjectKey{Value: "OPI-VirtioScsi8"}, TargetId: &pbc.ObjectKey{Value: "OPI-VirtioScsi8"}, VolumeId: &pbc.ObjectKey{Value: "Malloc1"}}})
	if err != nil {
		return err
	}
	log.Printf("Updated VirtioScsiLun: %v", rl3)
	rl4, err := c6.ListVirtioScsiLuns(ctx, &pb.ListVirtioScsiLunsRequest{Parent: "OPI-VirtioScsi8"})
	if err != nil {
		return err
	}
	log.Printf("Listed VirtioScsiLun: %v", rl4)
	rl5, err := c6.GetVirtioScsiLun(ctx, &pb.GetVirtioScsiLunRequest{Name: "OPI-VirtioScsi8"})
	if err != nil {
		return err
	}
	log.Printf("Got VirtioScsiLun: %v", rl5.VolumeId.Value)
	rl6, err := c6.VirtioScsiLunStats(ctx, &pb.VirtioScsiLunStatsRequest{ControllerId: &pbc.ObjectKey{Value: "OPI-VirtioScsi8"}})
	if err != nil {
		return err
	}
	log.Printf("Stats VirtioScsiLun: %v", rl6.Stats)
	rl2, err := c6.DeleteVirtioScsiLun(ctx, &pb.DeleteVirtioScsiLunRequest{Name: "OPI-VirtioScsi8"})
	if err != nil {
		return err
	}
	log.Printf("Deleted VirtioScsiLun: %v -> %v", rl1, rl2)
	rss2, err := c6.DeleteVirtioScsiController(ctx, &pb.DeleteVirtioScsiControllerRequest{Name: "OPI-VirtioScsi8"})
	if err != nil {
		return err
	}
	log.Printf("Deleted VirtioScsiController: %v -> %v", rss1, rss2)
	return err
}

func executeVirtioScsiController(ctx context.Context, c5 pb.FrontendVirtioScsiServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing VirtioScsiController")
	log.Printf("=======================================")
	rss1, err := c5.CreateVirtioScsiController(ctx, &pb.CreateVirtioScsiControllerRequest{VirtioScsiController: &pb.VirtioScsiController{Id: &pbc.ObjectKey{Value: "OPI-VirtioScsi8"}}})
	if err != nil {
		return err
	}
	log.Printf("Added VirtioScsiController: %v", rss1)
	rss3, err := c5.UpdateVirtioScsiController(ctx, &pb.UpdateVirtioScsiControllerRequest{VirtioScsiController: &pb.VirtioScsiController{Id: &pbc.ObjectKey{Value: "OPI-VirtioScsi8"}}})
	if err != nil {
		return err
	}
	log.Printf("Updated VirtioScsiController: %v", rss3)
	rss4, err := c5.ListVirtioScsiControllers(ctx, &pb.ListVirtioScsiControllersRequest{})
	if err != nil {
		return err
	}
	log.Printf("Listed VirtioScsiControllers: %s", rss4)
	rss5, err := c5.GetVirtioScsiController(ctx, &pb.GetVirtioScsiControllerRequest{Name: "OPI-VirtioScsi8"})
	if err != nil {
		return err
	}
	log.Printf("Got VirtioScsiController: %s", rss5.Id.Value)
	rss6, err := c5.VirtioScsiControllerStats(ctx, &pb.VirtioScsiControllerStatsRequest{ControllerId: &pbc.ObjectKey{Value: "OPI-VirtioScsi8"}})
	if err != nil {
		return err
	}
	log.Printf("Stats VirtioScsiController: %s", rss6.Stats)
	rss2, err := c5.DeleteVirtioScsiController(ctx, &pb.DeleteVirtioScsiControllerRequest{Name: "OPI-VirtioScsi8"})
	if err != nil {
		return err
	}
	log.Printf("Deleted VirtioScsiController: %v -> %v", rss1, rss2)
	return err
}

func executeVirtioBlk(ctx context.Context, c4 pb.FrontendVirtioBlkServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing VirtioBlk")
	log.Printf("=======================================")
	rv1, err := c4.CreateVirtioBlk(ctx, &pb.CreateVirtioBlkRequest{VirtioBlk: &pb.VirtioBlk{Id: &pbc.ObjectKey{Value: "VirtioBlk8"}, VolumeId: &pbc.ObjectKey{Value: "Malloc1"}}})
	if err != nil {
		return err
	}
	log.Printf("Added VirtioBlk: %v", rv1)
	rv3, err := c4.UpdateVirtioBlk(ctx, &pb.UpdateVirtioBlkRequest{VirtioBlk: &pb.VirtioBlk{Id: &pbc.ObjectKey{Value: "VirtioBlk8"}}})
	if err != nil {
		// UpdateVirtioBlk is not implemented, so no error here
		log.Printf("could not update VirtioBlk: %v", err)
	}
	log.Printf("Updated VirtioBlk: %v", rv3)
	rv4, err := c4.ListVirtioBlks(ctx, &pb.ListVirtioBlksRequest{})
	if err != nil {
		return err
	}
	log.Printf("Listed VirtioBlks: %v", rv4)
	rv5, err := c4.GetVirtioBlk(ctx, &pb.GetVirtioBlkRequest{Name: "VirtioBlk8"})
	if err != nil {
		return err
	}
	log.Printf("Got VirtioBlk: %v", rv5.Id.Value)
	rv6, err := c4.VirtioBlkStats(ctx, &pb.VirtioBlkStatsRequest{ControllerId: &pbc.ObjectKey{Value: "VirtioBlk8"}})
	if err != nil {
		// VirtioBlkStats is not implemented, so no error here
		log.Printf("could not stats VirtioBlk: %v", err)
	}
	log.Printf("Stats VirtioBlk: %v", rv6)
	rv2, err := c4.DeleteVirtioBlk(ctx, &pb.DeleteVirtioBlkRequest{Name: "VirtioBlk8"})
	if err != nil {
		return err
	}
	log.Printf("Deleted VirtioBlk: %v -> %v", rv1, rv2)

	return err
}

func executeNVMeNamespace(ctx context.Context, c2 pb.FrontendNvmeServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NVMeNamespace")
	log.Printf("=======================================")
	// pre create: subsystem and controller
	rs1, err := c2.CreateNVMeSubsystem(ctx, &pb.CreateNVMeSubsystemRequest{
		NvMeSubsystem: &pb.NVMeSubsystem{
			Spec: &pb.NVMeSubsystemSpec{
				Id:            &pbc.ObjectKey{Value: "namespace-test-ss"},
				ModelNumber:   "OPI Model",
				SerialNumber:  "OPI SN",
				MaxNamespaces: 10,
				Nqn:           "nqn.2022-09.io.spdk:opi1"}}})
	if err != nil {
		return err
	}
	log.Printf("Added NVMeSubsystem: %v", rs1)
	rc1, err := c2.CreateNVMeController(ctx, &pb.CreateNVMeControllerRequest{
		NvMeController: &pb.NVMeController{
			Spec: &pb.NVMeControllerSpec{
				Id:               &pbc.ObjectKey{Value: "namespace-test-ctrler"},
				SubsystemId:      &pbc.ObjectKey{Value: "namespace-test-ss"},
				PcieId:           &pb.PciEndpoint{PhysicalFunction: 1, VirtualFunction: 2, PortId: 3},
				MaxNsq:           5,
				MaxNcq:           6,
				Sqes:             7,
				Cqes:             8,
				NvmeControllerId: 1}}})
	if err != nil {
		return err
	}
	log.Printf("Added NVMeController: %v", rc1)

	// wait for some time for the backend to created above objects
	time.Sleep(time.Second)

	// NVMeNamespace
	rn1, err := c2.CreateNVMeNamespace(ctx, &pb.CreateNVMeNamespaceRequest{
		NvMeNamespace: &pb.NVMeNamespace{
			Spec: &pb.NVMeNamespaceSpec{
				Id:          &pbc.ObjectKey{Value: "namespace-test"},
				SubsystemId: &pbc.ObjectKey{Value: "namespace-test-ss"},
				VolumeId:    &pbc.ObjectKey{Value: "Malloc1"},
				Uuid:        &pbc.Uuid{Value: "1b4e28ba-2fa1-11d2-883f-b9a761bde3fb"},
				Nguid:       "1b4e28ba-2fa1-11d2-883f-b9a761bde3fb",
				Eui64:       1967554867335598546,
				HostNsid:    1}}})
	if err != nil {
		return err
	}
	log.Printf("Added NVMeNamespace: %v", rn1)
	rn3, err := c2.UpdateNVMeNamespace(ctx, &pb.UpdateNVMeNamespaceRequest{
		NvMeNamespace: &pb.NVMeNamespace{
			Spec: &pb.NVMeNamespaceSpec{
				Id:          &pbc.ObjectKey{Value: "namespace-test"},
				SubsystemId: &pbc.ObjectKey{Value: "namespace-test-ss"},
				VolumeId:    &pbc.ObjectKey{Value: "Malloc1"},
				Uuid:        &pbc.Uuid{Value: "1b4e28ba-2fa1-11d2-883f-b9a761bde3fb"},
				Nguid:       "1b4e28ba-2fa1-11d2-883f-b9a761bde3fb",
				Eui64:       1967554867335598546,
				HostNsid:    1}}})
	if err != nil {
		return err
	}
	log.Printf("Updated NVMeNamespace: %v", rn3)
	rn4, err := c2.ListNVMeNamespaces(ctx, &pb.ListNVMeNamespacesRequest{Parent: "namespace-test-ss"})
	if err != nil {
		return err
	}
	log.Printf("Listed NVMeNamespaces: %v", rn4)
	rn5, err := c2.GetNVMeNamespace(ctx, &pb.GetNVMeNamespaceRequest{Name: "namespace-test"})
	if err != nil {
		return err
	}
	log.Printf("Got NVMeNamespace: %v", rn5.Spec.Id.Value)
	rn6, err := c2.NVMeNamespaceStats(ctx, &pb.NVMeNamespaceStatsRequest{NamespaceId: &pbc.ObjectKey{Value: "namespace-test"}})
	if err != nil {
		return err
	}
	log.Printf("Stats NVMeNamespace: %v", rn6.Stats)
	rn2, err := c2.DeleteNVMeNamespace(ctx, &pb.DeleteNVMeNamespaceRequest{Name: "namespace-test"})
	if err != nil {
		return err
	}
	log.Printf("Deleted NVMeNamespace:  %v -> %v", rn1, rn2)

	// post cleanup: controller and subsystem
	rc2, err := c2.DeleteNVMeController(ctx, &pb.DeleteNVMeControllerRequest{Name: "namespace-test-ctrler"})
	if err != nil {
		return err
	}
	log.Printf("Deleted NVMeController: %v", rc2)

	// post cleanup: subsystem
	rs2, err := c2.DeleteNVMeSubsystem(ctx, &pb.DeleteNVMeSubsystemRequest{Name: "namespace-test-ss"})
	if err != nil {
		return err
	}
	log.Printf("Deleted NVMeSubsystem: %v -> %v", rs1, rs2)
	return nil
}

func executeNVMeController(ctx context.Context, c2 pb.FrontendNvmeServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NVMeController")
	log.Printf("=======================================")
	// pre create: subsystem
	rs1, err := c2.CreateNVMeSubsystem(ctx, &pb.CreateNVMeSubsystemRequest{
		NvMeSubsystem: &pb.NVMeSubsystem{
			Spec: &pb.NVMeSubsystemSpec{
				Id:            &pbc.ObjectKey{Value: "controller-test-ss"},
				ModelNumber:   "OPI Model",
				SerialNumber:  "OPI SN",
				MaxNamespaces: 10,
				Nqn:           "nqn.2022-09.io.spdk:opi2"}}})
	if err != nil {
		return err
	}
	log.Printf("Added NVMeSubsystem: %v", rs1)

	// NVMeController
	rc1, err := c2.CreateNVMeController(ctx, &pb.CreateNVMeControllerRequest{
		NvMeController: &pb.NVMeController{
			Spec: &pb.NVMeControllerSpec{
				Id:               &pbc.ObjectKey{Value: "controller-test"},
				SubsystemId:      &pbc.ObjectKey{Value: "controller-test-ss"},
				PcieId:           &pb.PciEndpoint{PhysicalFunction: 1, VirtualFunction: 2, PortId: 3},
				MaxNsq:           5,
				MaxNcq:           6,
				Sqes:             7,
				Cqes:             8,
				NvmeControllerId: 1}}})
	if err != nil {
		return err
	}
	log.Printf("Added NVMeController: %v", rc1)

	rc3, err := c2.UpdateNVMeController(ctx, &pb.UpdateNVMeControllerRequest{
		NvMeController: &pb.NVMeController{
			Spec: &pb.NVMeControllerSpec{
				Id:               &pbc.ObjectKey{Value: "controller-test"},
				SubsystemId:      &pbc.ObjectKey{Value: "controller-test-ss"},
				PcieId:           &pb.PciEndpoint{PhysicalFunction: 1, VirtualFunction: 2, PortId: 3},
				MaxNsq:           5,
				MaxNcq:           6,
				Sqes:             7,
				Cqes:             8,
				NvmeControllerId: 2}}})
	if err != nil {
		return err
	}
	log.Printf("Updated NVMeController: %v", rc3)

	rc4, err := c2.ListNVMeControllers(ctx, &pb.ListNVMeControllersRequest{Parent: "controller-test-ss"})
	if err != nil {
		return err
	}
	log.Printf("Listed NVMeControllers: %s", rc4)

	rc5, err := c2.GetNVMeController(ctx, &pb.GetNVMeControllerRequest{Name: "controller-test"})
	if err != nil {
		return err
	}
	log.Printf("Got NVMeController: %s", rc5.Spec.Id.Value)

	rc6, err := c2.NVMeControllerStats(ctx, &pb.NVMeControllerStatsRequest{Id: &pbc.ObjectKey{Value: "controller-test"}})
	if err != nil {
		return err
	}
	log.Printf("Stats NVMeController: %s", rc6.Stats)

	rc2, err := c2.DeleteNVMeController(ctx, &pb.DeleteNVMeControllerRequest{Name: "controller-test"})
	if err != nil {
		return err
	}
	log.Printf("Deleted NVMeController: %v -> %v", rc1, rc2)

	// post cleanup: subsystem
	rs2, err := c2.DeleteNVMeSubsystem(ctx, &pb.DeleteNVMeSubsystemRequest{Name: "controller-test-ss"})
	if err != nil {
		return err
	}
	log.Printf("Deleted NVMeSubsystem: %v -> %v", rs1, rs2)
	return nil
}

func executeNVMeSubsystem(ctx context.Context, c1 pb.FrontendNvmeServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NVMeSubsystem")
	log.Printf("=======================================")
	rs1, err := c1.CreateNVMeSubsystem(ctx, &pb.CreateNVMeSubsystemRequest{
		NvMeSubsystem: &pb.NVMeSubsystem{
			Spec: &pb.NVMeSubsystemSpec{
				Id:            &pbc.ObjectKey{Value: "subsystem-test"},
				ModelNumber:   "OPI Model",
				SerialNumber:  "OPI SN",
				MaxNamespaces: 10,
				Nqn:           "nqn.2022-09.io.spdk:opi3"}}})
	if err != nil {
		return err
	}
	log.Printf("Added NVMeSubsystem: %v", rs1)
	rs3, err := c1.UpdateNVMeSubsystem(ctx, &pb.UpdateNVMeSubsystemRequest{
		NvMeSubsystem: &pb.NVMeSubsystem{
			Spec: &pb.NVMeSubsystemSpec{
				Id:  &pbc.ObjectKey{Value: "subsystem-test"},
				Nqn: "nqn.2022-09.io.spdk:opi3"}}})
	if err != nil {
		// UpdateNVMeSubsystem is not implemented, so no error here
		log.Printf("could not update NVMe subsystem: %v", err)
	}
	log.Printf("Updated UpdateNVMeSubsystem: %v", rs3)
	rs4, err := c1.ListNVMeSubsystems(ctx, &pb.ListNVMeSubsystemsRequest{})
	if err != nil {
		return err
	}
	log.Printf("Listed UpdateNVMeSubsystems: %v", rs4)
	rs5, err := c1.GetNVMeSubsystem(ctx, &pb.GetNVMeSubsystemRequest{Name: "subsystem-test"})
	if err != nil {
		return err
	}
	log.Printf("Got UpdateNVMeSubsystem: %s", rs5.Spec.Nqn)
	rs6, err := c1.NVMeSubsystemStats(ctx, &pb.NVMeSubsystemStatsRequest{
		SubsystemId: &pbc.ObjectKey{Value: "subsystem-test"}})
	if err != nil {
		return err
	}
	log.Printf("Stats UpdateNVMeSubsystem: %s", rs6.Stats)

	// post cleanup: subsystem
	rs2, err := c1.DeleteNVMeSubsystem(ctx, &pb.DeleteNVMeSubsystemRequest{Name: "subsystem-test"})
	if err != nil {
		return err
	}
	log.Printf("Deleted NVMeSubsystem: %v -> %v", rs1, rs2)
	return nil
}

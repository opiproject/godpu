// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"
	"fmt"
	"log"
	"path"
	"time"

	"github.com/google/uuid"
	pbc "github.com/opiproject/opi-api/common/v1/gen/go"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/protobuf/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// DoFrontend executes the front end code
func DoFrontend(ctx context.Context, conn grpc.ClientConnInterface) error {
	nvme := pb.NewFrontendNvmeServiceClient(conn)
	blk := pb.NewFrontendVirtioBlkServiceClient(conn)
	scsi := pb.NewFrontendVirtioScsiServiceClient(conn)

	err := executeNvmeSubsystem(ctx, nvme)
	if err != nil {
		return err
	}
	err = executeNvmeController(ctx, nvme)
	if err != nil {
		return err
	}
	err = executeNvmeNamespace(ctx, nvme)
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
	const resourceID = "opi-virtio-scsi8"
	// pre create: controller
	rss1, err := c6.CreateVirtioScsiController(ctx, &pb.CreateVirtioScsiControllerRequest{
		VirtioScsiControllerId: resourceID,
		VirtioScsiController: &pb.VirtioScsiController{
			Name: "",
			PcieId: &pb.PciEndpoint{
				PhysicalFunction: wrapperspb.Int32(1),
				VirtualFunction:  wrapperspb.Int32(2),
				PortId:           wrapperspb.Int32(3)},
		}})
	if err != nil {
		return err
	}
	ctrlrName := fmt.Sprintf("//storage.opiproject.org/volumes/%s", resourceID)
	if rss1.Name != ctrlrName {
		return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rss1.Name, ctrlrName)
	}
	rl1, err := c6.CreateVirtioScsiLun(ctx, &pb.CreateVirtioScsiLunRequest{VirtioScsiLunId: resourceID, VirtioScsiLun: &pb.VirtioScsiLun{Name: "", TargetNameRef: resourceID, VolumeNameRef: "Malloc1"}})
	if err != nil {
		return err
	}
	fullname := fmt.Sprintf("//storage.opiproject.org/volumes/%s", resourceID)
	if rl1.Name != fullname {
		return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rl1.Name, fullname)
	}
	log.Printf("Added VirtioScsiLun: %v", rl1)
	rl3, err := c6.UpdateVirtioScsiLun(ctx, &pb.UpdateVirtioScsiLunRequest{
		UpdateMask:    &fieldmaskpb.FieldMask{Paths: []string{"*"}},
		VirtioScsiLun: &pb.VirtioScsiLun{Name: rl1.Name, TargetNameRef: resourceID, VolumeNameRef: "Malloc1"}})
	if err != nil {
		return err
	}
	log.Printf("Updated VirtioScsiLun: %v", rl3)
	rl4, err := c6.ListVirtioScsiLuns(ctx, &pb.ListVirtioScsiLunsRequest{Parent: rl1.Name})
	if err != nil {
		return err
	}
	log.Printf("Listed VirtioScsiLun: %v", rl4)
	rl5, err := c6.GetVirtioScsiLun(ctx, &pb.GetVirtioScsiLunRequest{Name: rl1.Name})
	if err != nil {
		return err
	}
	log.Printf("Got VirtioScsiLun: %v", rl5.VolumeNameRef)
	rl6, err := c6.StatsVirtioScsiLun(ctx, &pb.StatsVirtioScsiLunRequest{Name: rl1.Name})
	if err != nil {
		return err
	}
	log.Printf("Stats VirtioScsiLun: %v", rl6.Stats)
	rl2, err := c6.DeleteVirtioScsiLun(ctx, &pb.DeleteVirtioScsiLunRequest{Name: rl1.Name})
	if err != nil {
		return err
	}
	log.Printf("Deleted VirtioScsiLun: %v -> %v", rl1, rl2)
	rss2, err := c6.DeleteVirtioScsiController(ctx, &pb.DeleteVirtioScsiControllerRequest{Name: rl1.Name})
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

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"opi-virtio-scsi8", ""} {
		rss1, err := c5.CreateVirtioScsiController(ctx, &pb.CreateVirtioScsiControllerRequest{
			VirtioScsiControllerId: resourceID,
			VirtioScsiController: &pb.VirtioScsiController{
				Name: "",
				PcieId: &pb.PciEndpoint{
					PhysicalFunction: wrapperspb.Int32(1),
					VirtualFunction:  wrapperspb.Int32(2),
					PortId:           wrapperspb.Int32(3)},
			}})
		if err != nil {
			return err
		}
		// verify
		newResourceID := resourceID
		if resourceID == "" {
			parsed, err := uuid.Parse(path.Base(rss1.Name))
			if err != nil {
				return err
			}
			newResourceID = parsed.String()
		}
		fullname := fmt.Sprintf("//storage.opiproject.org/volumes/%s", newResourceID)
		if rss1.Name != fullname {
			return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rss1.Name, fullname)
		}
		log.Printf("Added VirtioScsiController: %v", rss1)
		rss3, err := c5.UpdateVirtioScsiController(ctx, &pb.UpdateVirtioScsiControllerRequest{
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			VirtioScsiController: &pb.VirtioScsiController{
				Name: rss1.Name,
				PcieId: &pb.PciEndpoint{
					PhysicalFunction: wrapperspb.Int32(1),
					VirtualFunction:  wrapperspb.Int32(2),
					PortId:           wrapperspb.Int32(3)},
			}})
		if err != nil {
			return err
		}
		log.Printf("Updated VirtioScsiController: %v", rss3)
		rss4, err := c5.ListVirtioScsiControllers(ctx, &pb.ListVirtioScsiControllersRequest{Parent: "todo"})
		if err != nil {
			return err
		}
		log.Printf("Listed VirtioScsiControllers: %s", rss4)
		rss5, err := c5.GetVirtioScsiController(ctx, &pb.GetVirtioScsiControllerRequest{Name: rss1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got VirtioScsiController: %s", rss5.Name)
		rss6, err := c5.StatsVirtioScsiController(ctx, &pb.StatsVirtioScsiControllerRequest{Name: rss1.Name})
		if err != nil {
			return err
		}
		log.Printf("Stats VirtioScsiController: %s", rss6.Stats)
		rss2, err := c5.DeleteVirtioScsiController(ctx, &pb.DeleteVirtioScsiControllerRequest{Name: rss1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted VirtioScsiController: %v -> %v", rss1, rss2)
	}
	return nil
}

func executeVirtioBlk(ctx context.Context, c4 pb.FrontendVirtioBlkServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing VirtioBlk")
	log.Printf("=======================================")

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"opi-virtio-blk8", ""} {
		rv1, err := c4.CreateVirtioBlk(ctx, &pb.CreateVirtioBlkRequest{
			VirtioBlkId: resourceID,
			VirtioBlk: &pb.VirtioBlk{
				Name:          "",
				VolumeNameRef: "Malloc1",
				PcieId: &pb.PciEndpoint{
					PhysicalFunction: wrapperspb.Int32(1),
					VirtualFunction:  wrapperspb.Int32(2),
					PortId:           wrapperspb.Int32(3)},
			}})
		if err != nil {
			return err
		}
		// verify
		newResourceID := resourceID
		if resourceID == "" {
			parsed, err := uuid.Parse(path.Base(rv1.Name))
			if err != nil {
				return err
			}
			newResourceID = parsed.String()
		}
		fullname := fmt.Sprintf("//storage.opiproject.org/volumes/%s", newResourceID)
		if rv1.Name != fullname {
			return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rv1.Name, fullname)
		}
		log.Printf("Added VirtioBlk: %v", rv1)
		rv3, err := c4.UpdateVirtioBlk(ctx, &pb.UpdateVirtioBlkRequest{
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			VirtioBlk:  &pb.VirtioBlk{Name: rv1.Name}})
		if err != nil {
			// UpdateVirtioBlk is not implemented, so no error here
			log.Printf("could not update VirtioBlk: %v", err)
		}
		log.Printf("Updated VirtioBlk: %v", rv3)
		rv4, err := c4.ListVirtioBlks(ctx, &pb.ListVirtioBlksRequest{Parent: "todo"})
		if err != nil {
			return err
		}
		log.Printf("Listed VirtioBlks: %v", rv4)
		rv5, err := c4.GetVirtioBlk(ctx, &pb.GetVirtioBlkRequest{Name: rv1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got VirtioBlk: %v", rv5.Name)
		rv6, err := c4.StatsVirtioBlk(ctx, &pb.StatsVirtioBlkRequest{Name: rv1.Name})
		if err != nil {
			// VirtioBlkStats is not implemented, so no error here
			log.Printf("could not stats VirtioBlk: %v", err)
		}
		log.Printf("Stats VirtioBlk: %v", rv6)
		rv2, err := c4.DeleteVirtioBlk(ctx, &pb.DeleteVirtioBlkRequest{Name: rv1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted VirtioBlk: %v -> %v", rv1, rv2)
	}
	return nil
}

func executeNvmeNamespace(ctx context.Context, c2 pb.FrontendNvmeServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NvmeNamespace")
	log.Printf("=======================================")

	// pre create: subsystem
	rs1, err := c2.CreateNvmeSubsystem(ctx, &pb.CreateNvmeSubsystemRequest{
		NvmeSubsystemId: "namespace-test-ss",
		NvmeSubsystem: &pb.NvmeSubsystem{
			Spec: &pb.NvmeSubsystemSpec{
				ModelNumber:   "OPI Model",
				SerialNumber:  "OPI SN",
				MaxNamespaces: 10,
				Nqn:           "nqn.2022-09.io.spdk:opi1"}}})
	if err != nil {
		return err
	}
	ssName := fmt.Sprintf("//storage.opiproject.org/volumes/%s", "namespace-test-ss")
	if rs1.Name != ssName {
		return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rs1.Name, ssName)
	}
	log.Printf("Added NvmeSubsystem: %v", rs1)

	// pre create: controller
	rc1, err := c2.CreateNvmeController(ctx, &pb.CreateNvmeControllerRequest{
		NvmeControllerId: "namespace-test-ctrler",
		NvmeController: &pb.NvmeController{
			Spec: &pb.NvmeControllerSpec{
				SubsystemNameRef: rs1.Name,
				PcieId: &pb.PciEndpoint{
					PhysicalFunction: wrapperspb.Int32(1),
					VirtualFunction:  wrapperspb.Int32(2),
					PortId:           wrapperspb.Int32(3)},
				MaxNsq:           5,
				MaxNcq:           6,
				Sqes:             7,
				Cqes:             8,
				NvmeControllerId: proto.Int32(1)}}})
	if err != nil {
		return err
	}
	ctrlrName := fmt.Sprintf("//storage.opiproject.org/volumes/%s", "namespace-test-ctrler")
	if rc1.Name != ctrlrName {
		return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rc1.Name, ctrlrName)
	}
	log.Printf("Added NvmeController: %v", rc1)

	// wait for some time for the backend to created above objects
	time.Sleep(time.Second)

	// NvmeNamespace

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"namespace-test", ""} {
		rn1, err := c2.CreateNvmeNamespace(ctx, &pb.CreateNvmeNamespaceRequest{
			NvmeNamespaceId: resourceID,
			NvmeNamespace: &pb.NvmeNamespace{
				Spec: &pb.NvmeNamespaceSpec{
					SubsystemNameRef: rs1.Name,
					VolumeNameRef:    "Malloc1",
					Uuid:             &pbc.Uuid{Value: "1b4e28ba-2fa1-11d2-883f-b9a761bde3fb"},
					Nguid:            "1b4e28ba-2fa1-11d2-883f-b9a761bde3fb",
					Eui64:            1967554867335598546,
					HostNsid:         1}}})
		if err != nil {
			return err
		}
		// verify
		newResourceID := resourceID
		if resourceID == "" {
			parsed, err := uuid.Parse(path.Base(rn1.Name))
			if err != nil {
				return err
			}
			newResourceID = parsed.String()
		}
		fullname := fmt.Sprintf("//storage.opiproject.org/volumes/%s", newResourceID)
		if rn1.Name != fullname {
			return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rn1.Name, fullname)
		}
		log.Printf("Added NvmeNamespace: %v", rn1)
		rn3, err := c2.UpdateNvmeNamespace(ctx, &pb.UpdateNvmeNamespaceRequest{
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			NvmeNamespace: &pb.NvmeNamespace{
				Name: rn1.Name,
				Spec: &pb.NvmeNamespaceSpec{
					SubsystemNameRef: rs1.Name,
					VolumeNameRef:    "Malloc1",
					Uuid:             &pbc.Uuid{Value: "1b4e28ba-2fa1-11d2-883f-b9a761bde3fb"},
					Nguid:            "1b4e28ba-2fa1-11d2-883f-b9a761bde3fb",
					Eui64:            1967554867335598546,
					HostNsid:         1}}})
		if err != nil {
			return err
		}
		log.Printf("Updated NvmeNamespace: %v", rn3)
		rn4, err := c2.ListNvmeNamespaces(ctx, &pb.ListNvmeNamespacesRequest{Parent: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Listed NvmeNamespaces: %v", rn4)
		rn5, err := c2.GetNvmeNamespace(ctx, &pb.GetNvmeNamespaceRequest{Name: rn1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got NvmeNamespace: %v", rn5.Name)
		rn6, err := c2.StatsNvmeNamespace(ctx, &pb.StatsNvmeNamespaceRequest{Name: rn1.Name})
		if err != nil {
			return err
		}
		log.Printf("Stats NvmeNamespace: %v", rn6.Stats)
		rn2, err := c2.DeleteNvmeNamespace(ctx, &pb.DeleteNvmeNamespaceRequest{Name: rn1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted NvmeNamespace:  %v -> %v", rn1, rn2)
	}

	// post cleanup: controller
	rc2, err := c2.DeleteNvmeController(ctx, &pb.DeleteNvmeControllerRequest{Name: rc1.Name})
	if err != nil {
		return err
	}
	log.Printf("Deleted NvmeController: %v", rc2)

	// post cleanup: subsystem
	rs2, err := c2.DeleteNvmeSubsystem(ctx, &pb.DeleteNvmeSubsystemRequest{Name: rs1.Name})
	if err != nil {
		return err
	}
	log.Printf("Deleted NvmeSubsystem: %v -> %v", rs1, rs2)
	return nil
}

func executeNvmeController(ctx context.Context, c2 pb.FrontendNvmeServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NvmeController")
	log.Printf("=======================================")
	// pre create: subsystem
	rs1, err := c2.CreateNvmeSubsystem(ctx, &pb.CreateNvmeSubsystemRequest{
		NvmeSubsystemId: "controller-test-ss",
		NvmeSubsystem: &pb.NvmeSubsystem{
			Spec: &pb.NvmeSubsystemSpec{
				ModelNumber:   "OPI Model",
				SerialNumber:  "OPI SN",
				MaxNamespaces: 10,
				Nqn:           "nqn.2022-09.io.spdk:opi2"}}})
	if err != nil {
		return err
	}
	ssName := fmt.Sprintf("//storage.opiproject.org/volumes/%s", "controller-test-ss")
	if rs1.Name != ssName {
		return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rs1.Name, ssName)
	}
	log.Printf("Added NvmeSubsystem: %v", rs1)

	// NvmeController

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"controller-test", ""} {
		rc1, err := c2.CreateNvmeController(ctx, &pb.CreateNvmeControllerRequest{
			NvmeControllerId: resourceID,
			NvmeController: &pb.NvmeController{
				Spec: &pb.NvmeControllerSpec{
					SubsystemNameRef: rs1.Name,
					PcieId: &pb.PciEndpoint{
						PhysicalFunction: wrapperspb.Int32(1),
						VirtualFunction:  wrapperspb.Int32(2),
						PortId:           wrapperspb.Int32(3)},
					MaxNsq:           5,
					MaxNcq:           6,
					Sqes:             7,
					Cqes:             8,
					NvmeControllerId: proto.Int32(1)}}})
		if err != nil {
			return err
		}
		// verify
		newResourceID := resourceID
		if resourceID == "" {
			parsed, err := uuid.Parse(path.Base(rc1.Name))
			if err != nil {
				return err
			}
			newResourceID = parsed.String()
		}
		fullname := fmt.Sprintf("//storage.opiproject.org/volumes/%s", newResourceID)
		if rc1.Name != fullname {
			return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rc1.Name, fullname)
		}
		log.Printf("Added NvmeController: %v", rc1)
		rc3, err := c2.UpdateNvmeController(ctx, &pb.UpdateNvmeControllerRequest{
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			NvmeController: &pb.NvmeController{
				Name: rc1.Name,
				Spec: &pb.NvmeControllerSpec{
					SubsystemNameRef: rs1.Name,
					PcieId: &pb.PciEndpoint{
						PhysicalFunction: wrapperspb.Int32(3),
						VirtualFunction:  wrapperspb.Int32(2),
						PortId:           wrapperspb.Int32(1)},
					MaxNsq:           8,
					MaxNcq:           7,
					Sqes:             6,
					Cqes:             5,
					NvmeControllerId: proto.Int32(1)}}})
		if err != nil {
			return err
		}
		log.Printf("Updated NvmeController: %v", rc3)

		rc4, err := c2.ListNvmeControllers(ctx, &pb.ListNvmeControllersRequest{Parent: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Listed NvmeControllers: %s", rc4)

		rc5, err := c2.GetNvmeController(ctx, &pb.GetNvmeControllerRequest{Name: rc1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got NvmeController: %s", rc5.Name)

		rc6, err := c2.StatsNvmeController(ctx, &pb.StatsNvmeControllerRequest{Name: rc1.Name})
		if err != nil {
			return err
		}
		log.Printf("Stats NvmeController: %s", rc6.Stats)

		rc2, err := c2.DeleteNvmeController(ctx, &pb.DeleteNvmeControllerRequest{Name: rc1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted NvmeController: %v -> %v", rc1, rc2)
	}

	// post cleanup: subsystem
	rs2, err := c2.DeleteNvmeSubsystem(ctx, &pb.DeleteNvmeSubsystemRequest{Name: rs1.Name})
	if err != nil {
		return err
	}
	log.Printf("Deleted NvmeSubsystem: %v -> %v", rs1, rs2)
	return nil
}

func executeNvmeSubsystem(ctx context.Context, c1 pb.FrontendNvmeServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NvmeSubsystem")
	log.Printf("=======================================")

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"subsystem-test", ""} {
		rs1, err := c1.CreateNvmeSubsystem(ctx, &pb.CreateNvmeSubsystemRequest{
			NvmeSubsystemId: resourceID,
			NvmeSubsystem: &pb.NvmeSubsystem{
				Spec: &pb.NvmeSubsystemSpec{
					ModelNumber:   "OPI Model",
					SerialNumber:  "OPI SN",
					MaxNamespaces: 10,
					Nqn:           "nqn.2022-09.io.spdk:opi3"}}})
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
		log.Printf("Added NvmeSubsystem: %v", rs1)
		rs3, err := c1.UpdateNvmeSubsystem(ctx, &pb.UpdateNvmeSubsystemRequest{
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			NvmeSubsystem: &pb.NvmeSubsystem{
				Name: rs1.Name,
				Spec: &pb.NvmeSubsystemSpec{
					Nqn: "nqn.2022-09.io.spdk:opi3"}}})
		if err != nil {
			// UpdateNvmeSubsystem is not implemented, so no error here
			log.Printf("could not update Nvme subsystem: %v", err)
		}
		log.Printf("Updated UpdateNvmeSubsystem: %v", rs3)
		rs4, err := c1.ListNvmeSubsystems(ctx, &pb.ListNvmeSubsystemsRequest{Parent: "todo"})
		if err != nil {
			return err
		}
		log.Printf("Listed UpdateNvmeSubsystems: %v", rs4)
		rs5, err := c1.GetNvmeSubsystem(ctx, &pb.GetNvmeSubsystemRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got UpdateNvmeSubsystem: %s", rs5.Spec.Nqn)
		rs6, err := c1.StatsNvmeSubsystem(ctx, &pb.StatsNvmeSubsystemRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Stats UpdateNvmeSubsystem: %s", rs6.Stats)

		// post cleanup: subsystem
		rs2, err := c1.DeleteNvmeSubsystem(ctx, &pb.DeleteNvmeSubsystemRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted NvmeSubsystem: %v -> %v", rs1, rs2)
	}
	return nil
}

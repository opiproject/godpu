// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"
	"fmt"
	"log"
	"path"

	"github.com/google/uuid"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// DoMiddleend executes the middle end code
func DoMiddleend(ctx context.Context, conn grpc.ClientConnInterface) error {
	encryption := pb.NewMiddleendEncryptionServiceClient(conn)
	qos := pb.NewMiddleendQosVolumeServiceClient(conn)
	err := executeEncryptedVolume(ctx, encryption)
	if err != nil {
		return err
	}
	err = executeQosVolume(ctx, qos)
	if err != nil {
		return err
	}
	return nil
}

func executeEncryptedVolume(ctx context.Context, c1 pb.MiddleendEncryptionServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewMiddleendEncryptionServiceClient")
	log.Printf("=======================================")

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"opi-encrypted-volume3", ""} {
		rs1, err := c1.CreateEncryptedVolume(ctx, &pb.CreateEncryptedVolumeRequest{
			EncryptedVolumeId: resourceID,
			EncryptedVolume: &pb.EncryptedVolume{
				VolumeNameRef: "Malloc1",
				Key:           []byte("0123456789abcdef0123456789abcdee"),
				Cipher:        pb.EncryptionType_ENCRYPTION_TYPE_AES_XTS_128,
			},
		})
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
		fullname := resourceIDToVolumeName(newResourceID)
		if rs1.Name != fullname {
			return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rs1.Name, fullname)
		}
		log.Printf("Added EncryptedVolume: %v", rs1)
		rs3, err := c1.UpdateEncryptedVolume(ctx, &pb.UpdateEncryptedVolumeRequest{
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			EncryptedVolume: &pb.EncryptedVolume{
				Name:          rs1.Name,
				VolumeNameRef: "Malloc1",
				Key:           []byte("0123456789abcdef0123456789abcdff"),
				Cipher:        pb.EncryptionType_ENCRYPTION_TYPE_AES_XTS_128,
			},
		})
		if err != nil {
			return err
		}
		log.Printf("Updated EncryptedVolume: %v", rs3)
		rs4, err := c1.ListEncryptedVolumes(ctx, &pb.ListEncryptedVolumesRequest{Parent: "todo"})
		if err != nil {
			return err
		}
		log.Printf("Listed EncryptedVolume: %v", rs4)
		rs5, err := c1.GetEncryptedVolume(ctx, &pb.GetEncryptedVolumeRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got EncryptedVolume: %s", rs5.Name)
		rs6, err := c1.StatsEncryptedVolume(ctx, &pb.StatsEncryptedVolumeRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Stats EncryptedVolume: %s", rs6.Stats)
		rs2, err := c1.DeleteEncryptedVolume(ctx, &pb.DeleteEncryptedVolumeRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted EncryptedVolume: %v -> %v", rs1, rs2)
	}
	return nil
}

func executeQosVolume(ctx context.Context, c2 pb.MiddleendQosVolumeServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewMiddleendQosVolumeServiceClient")
	log.Printf("=======================================")

	// testing with and without {resource}_id field
	for _, resourceID := range []string{"opi-qos-volume3", ""} {
		rs1, err := c2.CreateQosVolume(ctx, &pb.CreateQosVolumeRequest{
			QosVolumeId: resourceID,
			QosVolume: &pb.QosVolume{
				VolumeNameRef: "Malloc1",
				Limits: &pb.Limits{
					Max: &pb.QosLimit{
						RwBandwidthMbs: 2,
					},
				},
			},
		})
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
		fullname := resourceIDToVolumeName(newResourceID)
		if rs1.Name != fullname {
			return fmt.Errorf("server filled value '%s' is not matching user requested '%s'", rs1.Name, fullname)
		}
		log.Printf("Added QosVolume: %v", rs1)
		rs3, err := c2.UpdateQosVolume(ctx, &pb.UpdateQosVolumeRequest{
			UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"*"}},
			QosVolume: &pb.QosVolume{
				Name:          rs1.Name,
				VolumeNameRef: "Malloc1",
				Limits: &pb.Limits{
					Max: &pb.QosLimit{
						RdBandwidthMbs: 2,
					},
				},
			},
		})
		if err != nil {
			return err
		}
		log.Printf("Updated QosVolume: %v", rs3)
		rs4, err := c2.ListQosVolumes(ctx, &pb.ListQosVolumesRequest{Parent: "todo"})
		if err != nil {
			return err
		}
		log.Printf("Listed QosVolume: %v", rs4)
		rs5, err := c2.GetQosVolume(ctx, &pb.GetQosVolumeRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Got QosVolume: %v", rs5.Name)
		rs6, err := c2.StatsQosVolume(ctx, &pb.StatsQosVolumeRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Stats QosVolume: %v", rs6.Stats)
		rs2, err := c2.DeleteQosVolume(ctx, &pb.DeleteQosVolumeRequest{Name: rs1.Name})
		if err != nil {
			return err
		}
		log.Printf("Deleted QosVolume: %v -> %v", rs1, rs2)
	}
	return nil
}

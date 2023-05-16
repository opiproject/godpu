// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import (
	"context"
	"log"

	pc "github.com/opiproject/opi-api/common/v1/gen/go"
	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"google.golang.org/grpc"
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
	rs1, err := c1.CreateEncryptedVolume(ctx, &pb.CreateEncryptedVolumeRequest{
		EncryptedVolumeId: "OpiEncryptedVolume3",
		EncryptedVolume: &pb.EncryptedVolume{
			EncryptedVolumeId: &pc.ObjectKey{Value: "OpiEncryptedVolume3"},
			VolumeId:          &pc.ObjectKey{Value: "Malloc1"},
			Key:               []byte("0123456789abcdef0123456789abcdee"),
			Cipher:            pb.EncryptionType_ENCRYPTION_TYPE_AES_XTS_128,
		},
	})
	if err != nil {
		return err
	}
	log.Printf("Added EncryptedVolume: %v", rs1)
	rs3, err := c1.UpdateEncryptedVolume(ctx, &pb.UpdateEncryptedVolumeRequest{
		EncryptedVolume: &pb.EncryptedVolume{
			EncryptedVolumeId: &pc.ObjectKey{Value: "OpiEncryptedVolume3"},
			VolumeId:          &pc.ObjectKey{Value: "Malloc1"},
			Key:               []byte("0123456789abcdef0123456789abcdff"),
			Cipher:            pb.EncryptionType_ENCRYPTION_TYPE_AES_XTS_128,
		},
	})
	if err != nil {
		return err
	}
	log.Printf("Updated EncryptedVolume: %v", rs3)
	rs4, err := c1.ListEncryptedVolumes(ctx, &pb.ListEncryptedVolumesRequest{})
	if err != nil {
		return err
	}
	log.Printf("Listed EncryptedVolume: %v", rs4)
	rs5, err := c1.GetEncryptedVolume(ctx, &pb.GetEncryptedVolumeRequest{Name: "OpiEncryptedVolume3"})
	if err != nil {
		return err
	}
	log.Printf("Got EncryptedVolume: %s", rs5.EncryptedVolumeId.Value)
	rs6, err := c1.EncryptedVolumeStats(ctx, &pb.EncryptedVolumeStatsRequest{EncryptedVolumeId: &pc.ObjectKey{Value: "OpiEncryptedVolume3"}})
	if err != nil {
		return err
	}
	log.Printf("Stats EncryptedVolume: %s", rs6.Stats)
	rs2, err := c1.DeleteEncryptedVolume(ctx, &pb.DeleteEncryptedVolumeRequest{Name: "OpiEncryptedVolume3"})
	if err != nil {
		return err
	}
	log.Printf("Deleted EncryptedVolume: %v -> %v", rs1, rs2)
	return nil
}

func executeQosVolume(ctx context.Context, c2 pb.MiddleendQosVolumeServiceClient) error {
	log.Printf("=======================================")
	log.Printf("Testing NewMiddleendQosVolumeServiceClient")
	log.Printf("=======================================")
	rs1, err := c2.CreateQosVolume(ctx, &pb.CreateQosVolumeRequest{
		QosVolumeId: "OpiQosVolume3",
		QosVolume: &pb.QosVolume{
			QosVolumeId: &pc.ObjectKey{Value: "OpiQosVolume3"},
			VolumeId:    &pc.ObjectKey{Value: "Malloc1"},
			LimitMax: &pb.QosLimit{
				RwBandwidthMbs: 2,
			},
		},
	})
	if err != nil {
		return err
	}
	log.Printf("Added QosVolume: %v", rs1)
	rs3, err := c2.UpdateQosVolume(ctx, &pb.UpdateQosVolumeRequest{
		QosVolume: &pb.QosVolume{
			QosVolumeId: &pc.ObjectKey{Value: "OpiQosVolume3"},
			VolumeId:    &pc.ObjectKey{Value: "Malloc1"},
			LimitMax: &pb.QosLimit{
				RdBandwidthMbs: 2,
			},
		},
	})
	if err != nil {
		return err
	}
	log.Printf("Updated QosVolume: %v", rs3)
	rs4, err := c2.ListQosVolumes(ctx, &pb.ListQosVolumesRequest{})
	if err != nil {
		return err
	}
	log.Printf("Listed QosVolume: %v", rs4)
	rs5, err := c2.GetQosVolume(ctx, &pb.GetQosVolumeRequest{Name: "OpiQosVolume3"})
	if err != nil {
		return err
	}
	log.Printf("Got QosVolume: %v", rs5.QosVolumeId.Value)
	rs6, err := c2.QosVolumeStats(ctx, &pb.QosVolumeStatsRequest{VolumeId: &pc.ObjectKey{Value: "OpiQosVolume3"}})
	if err != nil {
		return err
	}
	log.Printf("Stats QosVolume: %v", rs6.Stats)
	rs2, err := c2.DeleteQosVolume(ctx, &pb.DeleteQosVolumeRequest{Name: "OpiQosVolume3"})
	if err != nil {
		return err
	}
	log.Printf("Deleted QosVolume: %v -> %v", rs1, rs2)
	return nil
}

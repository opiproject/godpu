// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

package goopicsi

import (
	"log"
	"testing"

	pb "github.com/opiproject/opi-api/storage/v1alpha1/gen/go"
	"github.com/stretchr/testify/assert"
)

func TestNVMeControllerConnect(t *testing.T) {
	resp, err := NVMeControllerConnect(&pb.NVMfRemoteController{
		Id:      12,
		Traddr:  "0.0.0.0", // Add a valid target address
		Subnqn:  "nqn",     // Add a valid NQN
		Trsvcid: 4420,
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
	assert.Error(t, err, "connection failed")
}

func TestNVMeControllerList(t *testing.T) {
	err := NVMeControllerList()
	if err != nil {
		log.Println(err)
	}
}

func TestNVMeControllerGet(t *testing.T) {
	err := NVMeControllerGet(12)
	if err != nil {
		log.Println(err)
	}
}

func TestNVMeControllerDisconnect(t *testing.T) {
	resp, err := NVMeControllerDisconnect(&pb.NVMfRemoteControllerDisconnectRequest{Id: 12})
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
	assert.Error(t, err, "disconnect failed")
}

// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022 Dell Inc, or its subsidiaries.

package goopicsi

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNVMeControllerConnect(t *testing.T) {
	err := NVMeControllerConnect(12, "", "", 44565)
	if err != nil {
		log.Println(err)
	}
	assert.Error(t, err)
}

func TestNVMeControllerList(t *testing.T) {
	resp, err := NVMeControllerList()
	if err != nil {
		log.Println(err)
	}
	log.Printf("NVMf Remote Connections: %v", resp)
}

func TestNVMeControllerGet(t *testing.T) {
	resp, err := NVMeControllerGet(12)
	if err != nil {
		log.Println(err)
	}
	log.Printf("NVMf remote connection corresponding to the ID: %v", resp)
}

func TestNVMeControllerDisconnect(t *testing.T) {
	err := NVMeControllerDisconnect(12)
	if err != nil {
		log.Println(err)
	}
}

func TestCreateNVMeNamespace(t *testing.T) {
	resp, err := CreateNVMeNamespace("1", "nqn", "nguid", 1)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

func TestDeleteNVMeNamespace(t *testing.T) {
	err := DeleteNVMeNamespace("1")
	if err != nil {
		log.Println(err)
	}
}

func TestGetSubSystemByID(t *testing.T) {
	resp, err := GetSubSystemByID("1")
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

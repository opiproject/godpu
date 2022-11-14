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
	err := NVMeControllerDisconnect(12)
	if err != nil {
		log.Println(err)
	}
}

func TestCreateNVMeNamespace(t *testing.T) {
	resp, err := CreateNVMeNamespace("1", "nqn", "opi", 1)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023 Intel Corporation

// Package storage implements the go library for OPI to be used in storage, for example, CSI drivers
package storage

import "go.einride.tech/aip/resourcename"

func resourceIDToVolumeName(resourceID string) string {
	return resourcename.Join(
		"//storage.opiproject.org/",
		"volumes", resourceID,
	)
}

func resourceIDToSubsystemName(resourceID string) string {
	return resourcename.Join(
		"//storage.opiproject.org/",
		"subsystems", resourceID,
	)
}

func resourceIDToNamespaceName(subsysResourceID, ctrlrResourceID string) string {
	return resourcename.Join(
		"//storage.opiproject.org/",
		"subsystems", subsysResourceID,
		"namespaces", ctrlrResourceID,
	)
}

func resourceIDToControllerName(subsysResourceID, ctrlrResourceID string) string {
	return resourcename.Join(
		"//storage.opiproject.org/",
		"subsystems", subsysResourceID,
		"controllers", ctrlrResourceID,
	)
}

func resourceIDToRemoteControllerName(resourceID string) string {
	return resourcename.Join(
		"//storage.opiproject.org/",
		"nvmeRemoteControllers", resourceID,
	)
}

func resourceIDToNvmePathName(ctrlrResourceID, pathResourceID string) string {
	return resourcename.Join(
		"//storage.opiproject.org/",
		"nvmeRemoteControllers", ctrlrResourceID,
		"nvmePaths", pathResourceID,
	)
}

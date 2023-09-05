// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network_test

import (
	"context"
	"errors"
	"testing"

	"github.com/opiproject/godpu/mocks"
	_go "github.com/opiproject/opi-api/network/evpn-gw/v1alpha1/gen/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func TestCreateLogicalBridgeSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedLogicalBridge := &_go.LogicalBridge{} // Create your expected response
	mockEvpnClient.On("CreateLogicalBridge", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(expectedLogicalBridge, nil)

	resultLogicalBridge, err := mockEvpnClient.CreateLogicalBridge(context.TODO(), "lb1", 1000, 1000, "10.10.10.10/16")

	assert.NoError(t, err)
	assert.Equal(t, expectedLogicalBridge, resultLogicalBridge)
	mockEvpnClient.AssertExpectations(t)
}

func TestCreateLogicalBridgeError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("CreateLogicalBridge", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultLogicalBridge, err := mockEvpnClient.CreateLogicalBridge(context.TODO(), "lb1", 1000, 1000, "10.10.10.10/16")

	assert.Error(t, err)
	assert.Nil(t, resultLogicalBridge)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestDeleteLogicalBridgeSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedResponse := &emptypb.Empty{} // Create your expected response
	mockEvpnClient.On("DeleteLogicalBridge", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedResponse, nil)

	resultResponse, err := mockEvpnClient.DeleteLogicalBridge(context.TODO(), "lb1", true)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, resultResponse)
	mockEvpnClient.AssertExpectations(t)
}

func TestDeleteLogicalBridgeError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("DeleteLogicalBridge", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultResponse, err := mockEvpnClient.DeleteLogicalBridge(context.TODO(), "lb1", true)

	assert.Error(t, err)
	assert.Nil(t, resultResponse)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestGetLogicalBridgeSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedLogicalBridge := &_go.LogicalBridge{} // Create your expected response
	mockEvpnClient.On("GetLogicalBridge", mock.Anything, mock.Anything).
		Return(expectedLogicalBridge, nil)

	resultLogicalBridge, err := mockEvpnClient.GetLogicalBridge(context.TODO(), "lb1")

	assert.NoError(t, err)
	assert.Equal(t, expectedLogicalBridge, resultLogicalBridge)
	mockEvpnClient.AssertExpectations(t)
}

func TestGetLogicalBridgeError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("GetLogicalBridge", mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultLogicalBridge, err := mockEvpnClient.GetLogicalBridge(context.TODO(), "lb1")

	assert.Error(t, err)
	assert.Nil(t, resultLogicalBridge)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestListLogicalBridgesSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedResponse := &_go.ListLogicalBridgesResponse{} // Create your expected response
	mockEvpnClient.On("ListLogicalBridges", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedResponse, nil)

	resultResponse, err := mockEvpnClient.ListLogicalBridges(context.TODO(), 10, "token")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, resultResponse)
	mockEvpnClient.AssertExpectations(t)
}

func TestListLogicalBridgesError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("ListLogicalBridges", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultResponse, err := mockEvpnClient.ListLogicalBridges(context.TODO(), 10, "token")

	assert.Error(t, err)
	assert.Nil(t, resultResponse)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestUpdateLogicalBridgeSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedLogicalBridge := &_go.LogicalBridge{} // Create your expected response
	mockEvpnClient.On("UpdateLogicalBridge", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedLogicalBridge, nil)

	resultLogicalBridge, err := mockEvpnClient.UpdateLogicalBridge(context.TODO(), "lb1", []string{"VlanId", "Vni"})

	assert.NoError(t, err)
	assert.Equal(t, expectedLogicalBridge, resultLogicalBridge)
	mockEvpnClient.AssertExpectations(t)
}

func TestUpdateLogicalBridgeError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("UpdateLogicalBridge", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultLogicalBridge, err := mockEvpnClient.UpdateLogicalBridge(context.TODO(), "name", []string{"VlanId", "Vni"})

	assert.Error(t, err)
	assert.Nil(t, resultLogicalBridge)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

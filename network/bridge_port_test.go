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
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBridgePortSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedBridgePort := &_go.BridgePort{} // Create your expected response
	mockEvpnClient.On("CreateBridgePort", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(expectedBridgePort, nil)

	resultBridgePort, err := mockEvpnClient.CreateBridgePort(context.TODO(), "bp1", "00:11:22:AA:BB:CC", "access", []string{"lb1", "lb2"})

	assert.NoError(t, err)
	assert.Equal(t, expectedBridgePort, resultBridgePort)
	mockEvpnClient.AssertExpectations(t)
}

func TestCreateBridgePortError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("CreateBridgePort", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultBridgePort, err := mockEvpnClient.CreateBridgePort(context.TODO(), "bp1", "00:11:22:AA:BB:CC", "access", []string{"lb1", "lb2"})

	assert.Error(t, err)
	assert.Nil(t, resultBridgePort)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestDeleteBridgePortSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedResponse := &emptypb.Empty{} // Create your expected response
	mockEvpnClient.On("DeleteBridgePort", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedResponse, nil)

	resultResponse, err := mockEvpnClient.DeleteBridgePort(context.TODO(), "bp1", true)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, resultResponse)
	mockEvpnClient.AssertExpectations(t)
}

func TestDeleteBridgePortError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("DeleteBridgePort", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultResponse, err := mockEvpnClient.DeleteBridgePort(context.TODO(), "bp1", true)

	assert.Error(t, err)
	assert.Nil(t, resultResponse)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestGetBridgePortSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedBridgePort := &_go.BridgePort{} // Create your expected response
	mockEvpnClient.On("GetBridgePort", mock.Anything, mock.Anything).
		Return(expectedBridgePort, nil)

	resultBridgePort, err := mockEvpnClient.GetBridgePort(context.TODO(), "bp1")

	assert.NoError(t, err)
	assert.Equal(t, expectedBridgePort, resultBridgePort)
	mockEvpnClient.AssertExpectations(t)
}

func TestGetBridgePortError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("GetBridgePort", mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultBridgePort, err := mockEvpnClient.GetBridgePort(context.TODO(), "bp1")

	assert.Error(t, err)
	assert.Nil(t, resultBridgePort)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestListBridgePortsSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedResponse := &_go.ListBridgePortsResponse{} // Create your expected response
	mockEvpnClient.On("ListBridgePorts", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedResponse, nil)

	resultResponse, err := mockEvpnClient.ListBridgePorts(context.TODO(), 10, "token")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, resultResponse)
	mockEvpnClient.AssertExpectations(t)
}

func TestListBridgePortsError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("ListBridgePorts", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultResponse, err := mockEvpnClient.ListBridgePorts(context.TODO(), 10, "token")

	assert.Error(t, err)
	assert.Nil(t, resultResponse)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestUpdateBridgePortSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedBridgePort := &_go.BridgePort{} // Create your expected response
	mockEvpnClient.On("UpdateBridgePort", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(expectedBridgePort, nil)

	resultBridgePort, err := mockEvpnClient.UpdateBridgePort(context.TODO(), "bp1", []string{"Ptype", "LogicalBridges"}, false)

	assert.NoError(t, err)
	assert.Equal(t, expectedBridgePort, resultBridgePort)
	mockEvpnClient.AssertExpectations(t)
}

func TestUpdateBridgePortError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("UpdateBridgePort", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultBridgePort, err := mockEvpnClient.UpdateBridgePort(context.TODO(), "bp1", []string{"Ptype", "LogicalBridges"}, false)

	assert.Error(t, err)
	assert.Nil(t, resultBridgePort)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

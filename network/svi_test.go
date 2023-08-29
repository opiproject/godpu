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

func TestCreateSviSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedSvi := &_go.Svi{} // Create your expected response
	mockEvpnClient.On("CreateSvi", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(expectedSvi, nil)

	resultSvi, err := mockEvpnClient.CreateSvi(context.TODO(), "svi1", "vrf1", "lb1", "00:11:22:AA:BB:CC", []string{"10.0.0.1/24, 20.0.0.1/24"}, true, 65000)

	assert.NoError(t, err)
	assert.Equal(t, expectedSvi, resultSvi)
	mockEvpnClient.AssertExpectations(t)
}

func TestCreateSviError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("CreateSvi", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultSvi, err := mockEvpnClient.CreateSvi(context.TODO(), "svi1", "vrf1", "lb1", "00:11:22:AA:BB:CC", []string{"10.0.0.1/24, 20.0.0.1/24"}, true, 65000)

	assert.Error(t, err)
	assert.Nil(t, resultSvi)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestDeleteSviSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedResponse := &emptypb.Empty{} // Create your expected response
	mockEvpnClient.On("DeleteSvi", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedResponse, nil)

	resultResponse, err := mockEvpnClient.DeleteSvi(context.TODO(), "svi1", true)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, resultResponse)
	mockEvpnClient.AssertExpectations(t)
}

func TestDeleteSviError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("DeleteSvi", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultResponse, err := mockEvpnClient.DeleteSvi(context.TODO(), "svi1", true)

	assert.Error(t, err)
	assert.Nil(t, resultResponse)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestGetSviSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedSvi := &_go.Svi{} // Create your expected response
	mockEvpnClient.On("GetSvi", mock.Anything, mock.Anything).
		Return(expectedSvi, nil)

	resultSvi, err := mockEvpnClient.GetSvi(context.TODO(), "svi1")

	assert.NoError(t, err)
	assert.Equal(t, expectedSvi, resultSvi)
	mockEvpnClient.AssertExpectations(t)
}

func TestGetSviError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("GetSvi", mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultSvi, err := mockEvpnClient.GetSvi(context.TODO(), "svi1")

	assert.Error(t, err)
	assert.Nil(t, resultSvi)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestListSvisSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedResponse := &_go.ListSvisResponse{} // Create your expected response
	mockEvpnClient.On("ListSvis", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedResponse, nil)

	resultResponse, err := mockEvpnClient.ListSvis(context.TODO(), 10, "token")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, resultResponse)
	mockEvpnClient.AssertExpectations(t)
}

func TestListSvisError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("ListSvis", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultResponse, err := mockEvpnClient.ListSvis(context.TODO(), 10, "token")

	assert.Error(t, err)
	assert.Nil(t, resultResponse)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestUpdateSviSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedSvi := &_go.Svi{} // Create your expected response
	mockEvpnClient.On("UpdateSvi", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(expectedSvi, nil)

	resultSvi, err := mockEvpnClient.UpdateSvi(context.TODO(), "svi1", []string{"field1", "field2"}, false)

	assert.NoError(t, err)
	assert.Equal(t, expectedSvi, resultSvi)
	mockEvpnClient.AssertExpectations(t)
}

func TestUpdateSviError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("UpdateSvi", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultSvi, err := mockEvpnClient.UpdateSvi(context.TODO(), "svi1", []string{"Vrf", "LogicalBridge"}, false)

	assert.Error(t, err)
	assert.Nil(t, resultSvi)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

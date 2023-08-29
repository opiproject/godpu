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

func TestCreateVrfSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedVrf := &_go.Vrf{} // Create your expected response
	mockEvpnClient.On("CreateVrf", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(expectedVrf, nil)

	resultVrf, err := mockEvpnClient.CreateVrf(context.TODO(), "vrf1", 1000, "10.10.10.10/16", "20.20.20.20/16")

	assert.NoError(t, err)
	assert.Equal(t, expectedVrf, resultVrf)
	mockEvpnClient.AssertExpectations(t)
}

func TestCreateVrfError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("CreateVrf", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultVrf, err := mockEvpnClient.CreateVrf(context.TODO(), "vrf1", 1000, "10.10.10.10/16", "20.20.20.20/16")

	assert.Error(t, err)
	assert.Nil(t, resultVrf)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestDeleteVrfSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedResponse := &emptypb.Empty{} // Create your expected response
	mockEvpnClient.On("DeleteVrf", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedResponse, nil)

	resultResponse, err := mockEvpnClient.DeleteVrf(context.TODO(), "vrf1", true)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, resultResponse)
	mockEvpnClient.AssertExpectations(t)
}

func TestDeleteVrfError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("DeleteVrf", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultResponse, err := mockEvpnClient.DeleteVrf(context.TODO(), "vrf1", true)

	assert.Error(t, err)
	assert.Nil(t, resultResponse)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestGetVrfSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedVrf := &_go.Vrf{} // Create your expected response
	mockEvpnClient.On("GetVrf", mock.Anything, mock.Anything).
		Return(expectedVrf, nil)

	resultVrf, err := mockEvpnClient.GetVrf(context.TODO(), "vrf1")

	assert.NoError(t, err)
	assert.Equal(t, expectedVrf, resultVrf)
	mockEvpnClient.AssertExpectations(t)
}

func TestGetVrfError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("GetVrf", mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultVrf, err := mockEvpnClient.GetVrf(context.TODO(), "vrf1")

	assert.Error(t, err)
	assert.Nil(t, resultVrf)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

func TestListVrfsSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedResponse := &_go.ListVrfsResponse{} // Create your expected response
	mockEvpnClient.On("ListVrfs", mock.Anything, mock.Anything, mock.Anything).
		Return(expectedResponse, nil)

	resultResponse, err := mockEvpnClient.ListVrfs(context.TODO(), 10, "token")

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, resultResponse)
	mockEvpnClient.AssertExpectations(t)
}

func TestListVrfsError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("ListVrfs", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultResponse, err := mockEvpnClient.ListVrfs(context.TODO(), 10, "token")

	assert.Error(t, err)
	assert.Nil(t, resultResponse)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}
func TestUpdateVrfSuccess(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedVrf := &_go.Vrf{} // Create your expected response
	mockEvpnClient.On("UpdateVrf", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(expectedVrf, nil)

	resultVrf, err := mockEvpnClient.UpdateVrf(context.TODO(), "vrf1", []string{"LoopbackIpPrefix", "VtepIpPrefix"}, false)

	assert.NoError(t, err)
	assert.Equal(t, expectedVrf, resultVrf)
	mockEvpnClient.AssertExpectations(t)
}

func TestUpdateVrfError(t *testing.T) {
	mockEvpnClient := &mocks.EvpnClient{}

	expectedError := errors.New("mocked error")
	mockEvpnClient.On("UpdateVrf", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, expectedError)

	resultVrf, err := mockEvpnClient.UpdateVrf(context.TODO(), "vrf1", []string{"LoopbackIpPrefix", "VtepIpPrefix"}, false)

	assert.Error(t, err)
	assert.Nil(t, resultVrf)
	assert.Equal(t, expectedError, err)
	mockEvpnClient.AssertExpectations(t)
}

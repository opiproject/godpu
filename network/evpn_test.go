// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.
// Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.

// Package network implements the go library for OPI to be used to establish networking
package network

import (
	"testing"
)

func TestNewLogicalBridge(t *testing.T) {
	tests := map[string]struct {
		address    string
		wantErr    bool
		wantClient bool
	}{
		"empty address": {
			address:    "",
			wantErr:    true,
			wantClient: false,
		},
		"non-empty address": {
			address:    "localhost:50051",
			wantErr:    false,
			wantClient: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			client, err := NewLogicalBridge(tt.address, "")
			if (err != nil) == !tt.wantErr {
				t.Errorf("expected err: %v, received: %v", tt.wantErr, err)
			}
			if (client != nil) == !tt.wantClient {
				t.Errorf("expected client: %v, received: %v", tt.wantClient, client)
			}
		})
	}
}

func TestNewBridgePort(t *testing.T) {
	tests := map[string]struct {
		address    string
		wantErr    bool
		wantClient bool
	}{
		"empty address": {
			address:    "",
			wantErr:    true,
			wantClient: false,
		},
		"non-empty address": {
			address:    "localhost:50051",
			wantErr:    false,
			wantClient: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			client, err := NewBridgePort(tt.address, "")
			if (err != nil) == !tt.wantErr {
				t.Errorf("expected err: %v, received: %v", tt.wantErr, err)
			}
			if (client != nil) == !tt.wantClient {
				t.Errorf("expected client: %v, received: %v", tt.wantClient, client)
			}
		})
	}
}

func TestNewVRF(t *testing.T) {
	tests := map[string]struct {
		address    string
		wantErr    bool
		wantClient bool
	}{
		"empty address": {
			address:    "",
			wantErr:    true,
			wantClient: false,
		},
		"non-empty address": {
			address:    "localhost:50051",
			wantErr:    false,
			wantClient: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			client, err := NewVRF(tt.address, "")
			if (err != nil) == !tt.wantErr {
				t.Errorf("expected err: %v, received: %v", tt.wantErr, err)
			}
			if (client != nil) == !tt.wantClient {
				t.Errorf("expected client: %v, received: %v", tt.wantClient, client)
			}
		})
	}
}

func TestNewSVI(t *testing.T) {
	tests := map[string]struct {
		address    string
		wantErr    bool
		wantClient bool
	}{
		"empty address": {
			address:    "",
			wantErr:    true,
			wantClient: false,
		},
		"non-empty address": {
			address:    "localhost:50051",
			wantErr:    false,
			wantClient: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			client, err := NewSVI(tt.address, "")
			if (err != nil) == !tt.wantErr {
				t.Errorf("expected err: %v, received: %v", tt.wantErr, err)
			}
			if (client != nil) == !tt.wantClient {
				t.Errorf("expected client: %v, received: %v", tt.wantClient, client)
			}
		})
	}
}

// SPDX-License-Identifier: Apache-2.0
// Copyright (C) 2023-2024 Intel Corporation

// Package frontend implements the go library for OPI frontend storage
package frontend

import (
	"testing"
)

func TestNewClient(t *testing.T) {
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
			client, err := New(tt.address, "")
			if (err != nil) == !tt.wantErr {
				t.Errorf("expected err: %v, received: %v", tt.wantErr, err)
			}
			if (client != nil) == !tt.wantClient {
				t.Errorf("expected client: %v, received: %v", tt.wantClient, client)
			}
		})
	}
}

# SPDX-License-Identifier: Apache-2.0
# Copyright (c) 2023 Dell Inc, or its subsidiaries.

ROOT_DIR='.'
PROJECTNAME=$(shell basename "$(PWD)")

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

go-compile: go-get go-build

go-build:
	@echo "  >  Building binaries..."
	@CGO_ENABLED=0 go build -o ${PROJECTNAME} .

go-get:
	@echo "  >  Checking if there are any missing dependencies..."
	@CGO_ENABLED=0 go get .

go-test:
	@echo "  >  Running ginkgo test suites..."
	# can replace with a recursive command ginkgo suites are defined for all packages
	ginkgo grpc inventory

mock-generate:
	@echo "  >  Starting mock code generation..."
	# Generate mocks for exported interfaces
	mockery --config=mocks/.mockery.yaml --name=Connector --dir=grpc
	mockery --config=mocks/.mockery.yaml --name=InvClient --dir=inventory

	# Generate mocks for imported protobuf clients too
	mockery --config=mocks/.mockery.yaml --name=InventorySvcClient --srcpkg=github.com/opiproject/opi-api/common/v1/gen/go
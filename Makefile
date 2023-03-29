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
	ginkgo common # can replace with a recursive command ginkgo suites are defined for all packages

## mock-generate: Generate the required mock files for interfaces.
mock-generate:
	@echo "  >  Starting mock code generation..."
	# Can replace with a single command for recursively creating mocks of all exported interfaces once all are implemented
	mockery --name=Client --dir=common
	mockery --name=Client --dir=inventory
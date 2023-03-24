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
	@echo "  >  Running tests..."
	go test ./...

## mock-generate: Generate the required mock files for interfaces.
mock-generate:
	@echo "  >  Starting mock code generation..."
	mockery --boilerplate-file=mocks/boilerplate.txt --name=Client --dir=common --inpackage
	mockery --boilerplate-file=mocks/boilerplate.txt --name=client --dir=inventory --inpackage
# SPDX-License-Identifier: Apache-2.0
# Copyright (c) 2023 Dell Inc, or its subsidiaries.

ROOT_DIR='.'
PROJECTNAME=$(shell basename "$(PWD)")

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

GOARCH ?= $(shell go env GOARCH) # detect automatically the underlying arch

go-compile: go-get go-build

tools:

	go get golang.org/x/tools/cmd/goimports
	go get github.com/kisielk/errcheck
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html
	go get github.com/tools/godep
	go get github.com/mitchellh/gox
	go get github.com/onsi/ginkgo
	go get -u golang.org/x/lint/golint
	go install golang.org/x/tools/cmd/goimports
	go install github.com/kisielk/errcheck
	go install github.com/vektra/mockery/v2@latest

go-build:
	@echo "  >  Building binaries..."
	@CGO_ENABLED=0 GOARCH=$(GOARCH) go build -o ${PROJECTNAME} .

go-get:
	@echo "  >  Checking if there are any missing dependencies..."
	@CGO_ENABLED=0 go get .

go-test:
	@echo "  >  Running ginkgo test suites..."
	# can replace with a recursive command ginkgo suites are defined for all packages
	ginkgo grpc inventory storage network

go-vet:
	@CGO_ENABLED=0 go vet -v ./...

go-errors:
	errcheck -ignoretests -blank ./...

go-lint:
	golint ./...

go-imports:
	goimports -l -w .

go-fmt:
	@CGO_ENABLED=0 go fmt ./...

mock-generate:
	@echo "  >  Starting mock code generation..."
	# Generate mocks for exported interfaces
	mockery --config=mocks/.mockery.yaml
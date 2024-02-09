# godpu

[![Linters](https://github.com/opiproject/godpu/actions/workflows/linters.yml/badge.svg)](https://github.com/opiproject/godpu/actions/workflows/linters.yml)
[![CodeQL](https://github.com/opiproject/godpu/actions/workflows/codeql.yml/badge.svg)](https://github.com/opiproject/godpu/actions/workflows/codeql.yml)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/opiproject/godpu/badge)](https://securityscorecards.dev/viewer/?platform=github.com&org=opiproject&repo=godpu)
[![Go](https://github.com/opiproject/godpu/actions/workflows/go.yml/badge.svg)](https://github.com/opiproject/godpu/actions/workflows/go.yml)
[![Docker](https://github.com/opiproject/godpu/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/opiproject/godpu/actions/workflows/docker-publish.yml)
[![License](https://img.shields.io/github/license/opiproject/godpu?style=flat-square&color=blue&label=License)](https://github.com/opiproject/godpu/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/opiproject/godpu/branch/main/graph/badge.svg)](https://codecov.io/gh/opiproject/godpu)
[![Go Report Card](https://goreportcard.com/badge/github.com/opiproject/godpu)](https://goreportcard.com/report/github.com/opiproject/godpu)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/opiproject/godpu)
[![Pulls](https://img.shields.io/docker/pulls/opiproject/godpu.svg?logo=docker&style=flat&label=Pulls)](https://hub.docker.com/r/opiproject/godpu)
[![Last Release](https://img.shields.io/github/v/release/opiproject/godpu?label=Latest&style=flat-square&logo=go)](https://github.com/opiproject/godpu/releases)
[![GitHub stars](https://img.shields.io/github/stars/opiproject/godpu.svg?style=flat-square&label=github%20stars)](https://github.com/opiproject/godpu)
[![GitHub Contributors](https://img.shields.io/github/contributors/opiproject/godpu.svg?style=flat-square)](https://github.com/opiproject/godpu/graphs/contributors)

Go library and cli to communicate with DPUs and IPUs.

## I Want To Contribute

This project welcomes contributions and suggestions.  We are happy to have the
Community involved via submission of **Issues and Pull Requests** (with
substantive content  or even just fixes). We are hoping for the documents,
test framework, etc. to become a community process with active engagement.
PRs can be reviewed by any number of people, and a maintainer may accept.

See [CONTRIBUTING](https://github.com/opiproject/opi/blob/main/CONTRIBUTING.md)
and [GitHub Basic Process](https://github.com/opiproject/opi/blob/main/doc-github-rules.md)
for more details.

## Installation

There are several ways of running this CLI.

### Docker

```sh
docker pull opiproject/godpu:<version>
```

You can specify a version like `0.1.1` or use `latest` to get the most up-to-date version.

Run latest version of the CLI in a container:

```sh
docker run --rm opiproject/godpu:latest --help
```

Replace `--help` with any `godpu` command, without `godpu` itself.

### Golang

```sh
go install -v github.com/opiproject/godpu@latest
```

or import

```go
import (
        "github.com/opiproject/godpu"
)
```

## Tests

Test your APIs even if unmerged using your private fork like this:

```bash
chmod a+w go.*
docker run --rm -it -v `pwd`:/app -w /app golang:alpine go mod edit -replace github.com/opiproject/opi-api@main=github.com/YOURUSERNAME/opi-api@main
docker run --rm -it -v `pwd`:/app -w /app golang:alpine go get -u github.com/YOURUSERNAME/opi-api/storage/v1alpha1/gen/go@a98ca449468a
docker run --rm -it -v `pwd`:/app -w /app golang:alpine go mod tidy
```

Generate mocks like this:

```bash
go install github.com/vektra/mockery/v2@latest
make mock-generate
```

## CLI

### Storage

```bash
alias dpu="docker run --rm --network host ghcr.io/opiproject/godpu:main"

# connect to remote nvme/tcp controller
nvmf0=$(dpu storage create backend nvme controller --id nvmf0 --multipath disable)
path0=$(dpu storage create backend nvme path tcp --controller "$nvmf0" --id path0 --ip "11.11.11.2" --port 4444 --nqn nqn.2016-06.io.spdk:cnode1 --hostnqn nqn.2014-08.org.nvmexpress:uuid:feb98abe-d51f-40c8-b348-2753f3571d3c)
dpu storage get backend nvme controller --name $nvmf0

# connect to local nvme/pcie ssd controller
nvmf1=$(dpu storage create backend nvme controller --id nvmf1 --multipath disable)
path1=$(dpu storage create backend nvme path pcie --controller "$nvmf1" --id path1 --bdf "0000:40:00.0")
dpu storage get backend nvme controller --name $nvmf1

# expose volume over nvme/tcp controller
ss0=$(dpu storage create frontend nvme subsystem --id subsys0 --nqn "nqn.2022-09.io.spdk:opitest0")
ns0=$(dpu storage create frontend nvme namespace --id namespace0 --volume "Malloc0" --subsystem "$ss0")
ctrl0=$(dpu storage create frontend nvme controller tcp --id ctrl0 --ip "127.0.0.1" --port 4420 --subsystem "$ss0")

# expose volume over emulated nvme/pcie controller
ss1=$(dpu storage create frontend nvme subsystem --id subsys1 --nqn "nqn.2022-09.io.spdk:opitest1")
ns1=$(dpu storage create frontend nvme namespace --id namespace1 --volume "Malloc1" --subsystem "$ss1")
ctrl1=$(dpu storage create frontend nvme controller pcie --id ctrl1 --port 0 --pf 0 --vf 0 --subsystem "$ss1")

# delete emulated nvme/pcie controller
dpu storage delete frontend nvme controller --name "$ctrl1"
dpu storage delete frontend nvme namespace --name "$ns1"
dpu storage delete frontend nvme subsystem --name "$ss1"

# delete nvme/tcp controller
dpu storage delete frontend nvme controller --name "$ctrl0"
dpu storage delete frontend nvme namespace --name "$ns0"
dpu storage delete frontend nvme subsystem --name "$ss0"

# disconnect from local nvme/pcie ssd controller
dpu storage delete backend nvme path --name "$path1"
dpu storage delete backend nvme controller --name "$nvmf1"

# disconnect from remote nvme/tcp controller
dpu storage delete backend nvme path --name "$path0"
dpu storage delete backend nvme controller --name "$nvmf0"
```

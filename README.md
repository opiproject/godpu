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

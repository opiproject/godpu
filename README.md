# godpu

[![Linters](https://github.com/opiproject/godpu/actions/workflows/linters.yml/badge.svg)](https://github.com/opiproject/godpu/actions/workflows/linters.yml)
[![Go](https://github.com/opiproject/godpu/actions/workflows/go.yml/badge.svg)](https://github.com/opiproject/godpu/actions/workflows/go.yml)
[![Docker](https://github.com/opiproject/godpu/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/opiproject/godpu/actions/workflows/docker-publish.yml)
[![License](https://img.shields.io/github/license/opiproject/godpu?style=flat-square&color=blue&label=License)](https://github.com/opiproject/godpu/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/opiproject/godpu/branch/main/graph/badge.svg)](https://codecov.io/gh/opiproject/godpu)
[![Go Report Card](https://goreportcard.com/badge/github.com/opiproject/godpu)](https://goreportcard.com/report/github.com/opiproject/godpu)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/opiproject/godpu)
[![Pulls](https://img.shields.io/docker/pulls/opiproject/godpu.svg?logo=docker&style=flat&label=Pulls)](https://hub.docker.com/r/opiproject/godpu)
[![Last Release](https://img.shields.io/github/v/release/opiproject/godpu?label=Latest&style=flat-square&logo=go)](https://github.com/opiproject/godpu/releases)

A Container Storage Interface (CSI) library, client, and other helpful utilities created with Go for OPI

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

# SPDX-License-Identifier: Apache-2.0
# Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
FROM docker.io/library/golang:1.23.4-alpine3.21 as builder

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

ENV CGO_ENABLED=0

# build an app
COPY . .
RUN go build -v -o /dpu .

FROM alpine:3.21

WORKDIR /

COPY --from=builder /dpu .
RUN chmod +x dpu

ENTRYPOINT [ "/dpu" ]

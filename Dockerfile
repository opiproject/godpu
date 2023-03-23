# SPDX-License-Identifier: Apache-2.0
# Copyright (c) 2022 Dell Inc, or its subsidiaries.

FROM docker.io/library/golang:1.20.2-alpine as builder

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# build an app
COPY . .
RUN CGO_ENABLED=0 go build -v -o ./dpu .

FROM alpine:3.17
WORKDIR /
COPY --from=builder /app/dpu .
RUN chmod +x dpu
ENTRYPOINT [ "/dpu" ]

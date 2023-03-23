# SPDX-License-Identifier: Apache-2.0
# Copyright (c) 2022-2023 Dell Inc, or its subsidiaries.
FROM docker.io/library/golang:1.20.2-alpine3.17 as builder

WORKDIR /build

# create a static binary
COPY . .
RUN go mod download && CGO_ENABLED=0 \
    go build -v -o ./dpu .

FROM alpine:3.17

WORKDIR /

COPY --from=builder /build/dpu .
RUN chmod +x dpu

ENTRYPOINT [ "/dpu" ]

module github.com/opiproject/godpu

go 1.19

require (
	github.com/golangci/golangci-lint v1.55.2
	github.com/PraserX/ipconv v1.2.0
	github.com/go-chi/chi v1.5.5
	github.com/go-ping/ping v1.1.0
	github.com/google/uuid v1.5.0
	github.com/lithammer/fuzzysearch v1.1.8
	github.com/onsi/ginkgo/v2 v2.14.0
	github.com/onsi/gomega v1.30.0
	github.com/opiproject/opi-api v0.0.0-20240304222410-5dba226aaa9e
	github.com/opiproject/opi-evpn-bridge v0.2.0
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.8.4
	go.einride.tech/aip v0.66.0
	golang.org/x/net v0.21.0
	golang.org/x/text v0.14.0
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.32.0
)

replace github.com/opiproject/opi-evpn-bridge => github.com/mardim91/opi-evpn-bridge v0.0.0-20241209100717-35ff8ff45934

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20210407192527-94a9f03dee38 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.1 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	google.golang.org/genproto v0.0.0-20240108191215-35c7eff3a6b1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240108191215-35c7eff3a6b1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240108191215-35c7eff3a6b1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

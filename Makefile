# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gqdc android ios gqdc-cross swarm evm all test clean
.PHONY: gqdc-linux gqdc-linux-386 gqdc-linux-amd64 gqdc-linux-mips64 gqdc-linux-mips64le
.PHONY: gqdc-linux-arm gqdc-linux-arm-5 gqdc-linux-arm-6 gqdc-linux-arm-7 gqdc-linux-arm64
.PHONY: gqdc-darwin gqdc-darwin-amd64
.PHONY: gqdc-windows gqdc-windows-386 gqdc-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

gqdc:
	build/env.sh go run build/ci.go install ./cmd/gqdc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gqdc\" to launch gqdc."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gqdc.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gqdc.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

lint: ## Run linters.
	build/env.sh go run build/ci.go lint

clean:
	./build/clean_go_build_cache.sh
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

swarm-devtools:
	env GOBIN= go install ./cmd/swarm/mimegen

# Cross Compilation Targets (xgo)

gqdc-cross: gqdc-linux gqdc-darwin gqdc-windows gqdc-android gqdc-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-*

gqdc-linux: gqdc-linux-386 gqdc-linux-amd64 gqdc-linux-arm gqdc-linux-mips64 gqdc-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-*

gqdc-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gqdc
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep 386

gqdc-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gqdc
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep amd64

gqdc-linux-arm: gqdc-linux-arm-5 gqdc-linux-arm-6 gqdc-linux-arm-7 gqdc-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep arm

gqdc-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gqdc
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep arm-5

gqdc-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gqdc
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep arm-6

gqdc-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gqdc
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep arm-7

gqdc-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gqdc
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep arm64

gqdc-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gqdc
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep mips

gqdc-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gqdc
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep mipsle

gqdc-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gqdc
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep mips64

gqdc-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gqdc
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-linux-* | grep mips64le

gqdc-darwin: gqdc-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-darwin-*

gqdc-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gqdc
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-darwin-* | grep 386

gqdc-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gqdc
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-darwin-* | grep amd64

gqdc-windows: gqdc-windows-386 gqdc-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-windows-*

gqdc-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gqdc
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-windows-* | grep 386

gqdc-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gqdc
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gqdc-windows-* | grep amd64

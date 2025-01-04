#!/bin/bash

export PKG_CONFIG_PATH=$PWD

build=$(date +%FT%T%z)
commit=$(git rev-parse --short HEAD)
version="$1"

ldflags="-s -w -X github.com/ai-flowx/flowx/config.Build=$build -X github.com/ai-flowx/flowx/config.Commit=$commit -X github.com/ai-flowx/flowx/config.Version=$version"
target="flowx"

go env -w GOPROXY=https://goproxy.cn,direct

CGO_ENABLED=1 GOARCH=$(go env GOARCH) GOOS=$(go env GOOS) go build -ldflags "$ldflags" -o bin/$target main.go

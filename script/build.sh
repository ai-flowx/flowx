#!/bin/bash

build=$(date +%FT%T%z)
commit=$(git rev-parse --short HEAD)
version="$1"

ldflags="-s -w -X github.com/ai-flowx/flowx/config.Build=$build -X github.com/ai-flowx/flowx/config.Commit=$commit -X github.com/ai-flowx/flowx/config.Version=$version"
target="flowx"

go env -w GOPROXY=https://goproxy.cn,direct

CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "$ldflags" -o bin/$target main.go
CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -ldflags "$ldflags" -o bin/$target.exe main.go

#!/bin/bash -e
cd "$(dirname "$0")"

unset GOPATH

go mod download
gofmt -l -w -s *.go
go vet ./...
go test -race ./...

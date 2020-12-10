#!/bin/bash -e
cd $(dirname $0)

unset GOPATH

go mod download

gofmt -l -w *.go
go vet ./...
go test -race ./...
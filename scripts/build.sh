#!/bin/sh
cd "$( dirname "$0"  )"
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -i -o ../bin/SerialTestServer ../cmd/main.go
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -i -o ../bin/SerialTestServer ../cmd/main.go
#!/usr/bin/env bash

### Code formatting

# gofmt
go install mvdan.cc/gofumpt@latest
gofumpt -w .

# golines
go install github.com/segmentio/golines@latest
golines -w .

### Import formatting

# dedupimport
go install github.com/nishanths/dedupimport@latest
dedupimport  -w .

# goimports
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .


### swaggo
go install github.com/swaggo/swag/cmd/swag@latest
swag fmt
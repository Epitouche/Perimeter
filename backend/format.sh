#!/usr/bin/env bash

### Code formatting

# gofmt
go install mvdan.cc/gofumpt@latest
gofumpt -w .

go install github.com/segmentio/golines@latest
golines -w .

### Import formatting

# dedupimport
go install github.com/nishanths/dedupimport@latest
dedupimport  -l .

# gci
go install github.com/daixiang0/gci@latest
gci write .

# goimports
go install golang.org/x/tools/cmd/goimports@latest
goimports -w .
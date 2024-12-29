#!/usr/bin/env bash

set -e
set -a
source ../.env
go test -cover -race -parallel 10 -coverprofile=cover.out ./...
go tool cover -html cover.out -o cover.html
go install github.com/nikolaydubina/go-cover-treemap@latest
go-cover-treemap -coverprofile cover.out >cover.svg
set +a

#!/usr/bin/env bash

# swaggo
go install github.com/swaggo/swag/cmd/swag@latest
swag init --parseDependency --parseInternal
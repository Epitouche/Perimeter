#!/usr/bin/env bash

echo "Generating Go documentation..."

# godoc
go install golang.org/x/tools/cmd/godoc@latest
# godoc-static
go install gitlab.com/tslocum/godoc-static@latest
mkdir -p godoc
godoc-static -link-index -site-name="AREA Documentation" -site-description-file=DevDoc.md -destination=godoc .

echo "Generating Swagger documentation..."

# swaggo
go install github.com/swaggo/swag/cmd/swag@latest
swag init --parseDependency --parseInternal

if [ "$(command -v npm)" ]; then
    npx @redocly/cli build-docs docs/swagger.yaml
    mv redoc-static.html godoc/redoc-static.html
fi
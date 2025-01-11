## BACKEND SERVER

FROM golang:1.23-alpine AS dependencies

WORKDIR /app

COPY ./go.mod .

COPY ./go.sum .

RUN go mod download


FROM dependencies AS build

COPY ./ .

RUN go build -o /app


FROM dependencies AS test

COPY ./ .

CMD ["go", "test", "-v"]

FROM alpine:3.21 AS run

WORKDIR /app

COPY --from=build /app/area .

EXPOSE 8080

CMD ["/app/area"]

## BACKEND DOCUMENTATION

FROM golang:1.23-alpine AS godoc-dependencies

WORKDIR /app

# godoc
RUN go install golang.org/x/tools/cmd/godoc@latest
# godoc-static
RUN go install gitlab.com/tslocum/godoc-static@latest
RUN mkdir -p godoc


COPY . .
RUN godoc-static -link-index -site-name="AREA Documentation" -site-description-file=DevDoc.md -destination=godoc .

FROM node:20-alpine AS redoc-dependencies

WORKDIR /app

RUN npm install @redocly/cli@latest -y

COPY docs/swagger.yaml swagger.yaml

RUN npx @redocly/cli build-docs swagger.yaml


FROM alpine:3.21 AS documentation

WORKDIR /app

# Install http-server in the final stage
RUN apk add --no-cache nodejs npm \
    && npm install -g http-server@latest

COPY --from=godoc-dependencies /app/godoc .

COPY --from=redoc-dependencies /app/redoc-static.html .


EXPOSE 6060

CMD ["http-server", "-p", "6060"]
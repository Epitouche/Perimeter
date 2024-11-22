FROM golang:1.22-alpine AS dependencies

WORKDIR /app

COPY project/go.mod .

COPY project/go.sum .

RUN go mod download


FROM dependencies AS build

COPY project/ .

RUN go build -o /app


FROM alpine:3.15 AS run

WORKDIR /app

COPY --from=build /app/Perimeter .

EXPOSE 8080

CMD ["/app/Perimeter"]
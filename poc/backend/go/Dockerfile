FROM golang:1.22-alpine AS dependencies

WORKDIR /app

COPY backend/go/project/go.mod .

COPY backend/go/project/go.sum .

RUN go mod download


FROM dependencies AS build

COPY backend/go/project/ .

RUN go build -o /app


FROM dependencies AS test

COPY backend/go/project/ .

CMD ["go", "test", "-v"]

FROM alpine:3.15 AS run

WORKDIR /app

COPY --from=build /app/Perimeter .

EXPOSE 8080

CMD ["/app/Perimeter"]
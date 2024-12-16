FROM golang:1.23-alpine AS build

WORKDIR /app

RUN apk add --no-cache make
RUN go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.* .
RUN go mod download

COPY Makefile .
COPY migrations ./migrations/

COPY cmd ./cmd/
COPY internal ./internal/

RUN make build

SHELL ["/bin/sh", "-c"]
CMD make mig && ./main

EXPOSE 8080

# Default to Go 1.18
ARG GO_VERSION=1.18

# Start from golang v1.18 base image
FROM golang:${GO_VERSION}-alpine

WORKDIR /usr/src/app

COPY . .
RUN go mod download

# Installing cosmtrek/air for live reload
RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT [ "air" ]
# Default to Go 1.19
ARG GO_VERSION=1.19

# Start from golang v1.18 base image
FROM golang:${GO_VERSION}-alpine AS builder

# 'upx' is a utility used to compress executable files to reduce their size.
RUN apk add --no-cache upx

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Import the code from the context.
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOFLAGS=-mod=mod GOOS=linux  go build -ldflags="-w -s" -o  /app cmd/api/main.go

# Compress the build
RUN upx --best --lzma /app

######## Start a new stage from scratch #######
# Final stage: the running container.
FROM alpine:latest AS final

# Import the compiled executable from the first stage.
COPY --from=builder /app /app
COPY --from=builder /src/internal/config/.env ./internal/config/.env

EXPOSE 8080

# Run the compiled binary.
ENTRYPOINT ["/app"]
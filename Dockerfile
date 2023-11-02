
# Start from the latest golang base image
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0
# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

ENTRYPOINT CompileDaemon -build="go build -a -installsuffix cgo -o /build/main ./cmd/main.go" -command="/build/main"

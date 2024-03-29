# Start from golang base image
FROM golang:1.17.3-alpine

# Update Alpine
RUN apk update

# Install utils packages.
RUN apk add --no-cache git

RUN apk add --no-cache build-base

RUN git --version
# Set the current working directory inside the container
WORKDIR /go/src/github.com/jordanlanch/ionix-test

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.7.0

RUN go get -v github.com/canthefason/go-watcher && go install github.com/canthefason/go-watcher/cmd/watcher

ENV GIN_MODE="debug"

ENTRYPOINT ["watcher", "-watch", "github.com/jordanlanch/ionix-test", "-buildvcs=false"]

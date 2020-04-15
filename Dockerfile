FROM golang:alpine as builder
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o /envault .

## copy only build file
FROM node:12-alpine

LABEL maintainer="razzkumar <razzkumar.dev@gmail.com>"
LABEL version="0.1.0"
LABEL repository="https://github.com/pratishshr/envault"


# install build utilities and aws sdk
RUN apk update && \
      apk add curl unzip python && \
      curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "awscli-bundle.zip" && \
      unzip awscli-bundle.zip && \
      ./awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws && \
      apk del curl unzip && \
      rm -rf awscli-bundle*

COPY --from=builder /envault /usr/local/bin/envault

# Command to run when starting the container
CMD ["envault","run"]

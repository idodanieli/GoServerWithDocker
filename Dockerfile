# Use the official golang image as a base image
FROM golang:1.22.2-alpine3.19 AS base

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY src/ ./src

# Build the Go app
RUN go build -o idos_server ./src/main/

# Start a new stage from the minimal Alpine Linux image for the final runtime environment
FROM alpine:3.19

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy only the executable file from the previous stage
COPY --from=base /app/idos_server .

# Copy the configuration files
COPY config/ ./config

# Expose port 90 to the outside world
EXPOSE 90

# Install pkl
RUN apk update
RUN apk add curl
RUN curl -L -o pkl https://github.com/apple/pkl/releases/download/0.25.3/pkl-alpine-linux-amd64
RUN chmod +x pkl
RUN mv pkl /bin/

# Command to run the executable
ENTRYPOINT  ["/app/idos_server"]

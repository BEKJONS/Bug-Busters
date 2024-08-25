# Stage 1: Build the Go application
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Copy the environment file
COPY .env .

# Build the Go application
# Adding GOFLAGS=-mod=mod to disable module vendoring
RUN GOFLAGS=-mod=mod CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Download the wait-for-it.sh script
RUN curl -o /wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Stage 2: Create the final runtime image
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates bash

WORKDIR /app

# Copy the built Go binary and other files from the builder stage
COPY --from=builder /app .
COPY --from=builder /app/main .
COPY --from=builder /wait-for-it.sh /wait-for-it.sh

# Create a directory for logs
RUN mkdir -p pkg/logs

# Expose the application's port
EXPOSE 8080

# Ensure the main binary is executable
RUN chmod +x ./main

# Define the command to run the application with wait-for-it.sh
CMD ["/bin/bash", "/wait-for-it.sh", "--", "./main"]

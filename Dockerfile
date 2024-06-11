# Use the official golang image as a base image
FROM golang:alpine AS builder

# Install git and other dependencies
RUN apk update && apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Clone your git repository
RUN git clone https://github.com/pranjal123662/Subscription-API.git 

# Navigate into the cloned directory
WORKDIR /app/Subscription-API

#Take the pull
RUN git pull origin master

# Build the Go application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

# Stage 2: Create a minimal image for running the Go application
FROM alpine:latest

# Install necessary libraries (if any)
RUN apk add --no-cache ca-certificates

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/Subscription-API/main /app/main

# Set the working directory inside the container
WORKDIR /app

# Expose port 7689 to the outside world
EXPOSE 7689

# Command to run the executable
CMD ["./main"]
# Use an official Go image based on Alpine
FROM golang:1.23.4-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install dependencies required for Go and for building the binary
RUN apk add --no-cache git

# Copy go.mod and go.sum files to cache dependencies
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal Alpine image for the runtime
FROM alpine:latest

# Set the working directory inside the runtime image
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy templates (if applicable)
COPY templates ./templates

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]

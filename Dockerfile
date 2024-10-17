# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Install git (required for Go modules)
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod tidy

# Copy the rest of the source code to the container
COPY . .

# Build the Go binary
RUN go build -o main .

# Stage 2: Build the final lightweight image
FROM alpine:latest

# Set a working directory in the final image
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port on which the app will run
EXPOSE 5600

# Run the Go binary
CMD ["./main"]
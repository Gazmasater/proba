# Stage 1: Build the Go binary
FROM golang:1.22.2 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o main .

# Stage 2: Create a minimal image to run the Go binary
FROM alpine:latest

# Install necessary runtime dependencies including gcompat for glibc compatibility
RUN apk --no-cache add ca-certificates gcompat

# Set the working directory
WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the application's port
EXPOSE 8080

# Run the Go binary
CMD ["./main"]

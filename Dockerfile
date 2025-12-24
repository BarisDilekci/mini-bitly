# Stage 1: Build the application
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download dependencies (if any)
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# -o server: output binary name
# ./cmd/server: entry point
RUN go build -o server ./cmd/server

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./server"]

# Build Stage
FROM golang:latest AS builder

# Set working directory
WORKDIR /app

# Copy the application source code
COPY . .

# Build the Go binary from the ./cmd directory
RUN CGO_ENABLED=0 GOOS=linux go build -o app-binary ./cmd

# Minimal Final Stage
FROM scratch

# Copy the binary from the builder stage to the root of the scratch image
COPY --from=builder /app/app-binary /app-binary

# Set the entry point to the Go binary
CMD ["/app-binary"]

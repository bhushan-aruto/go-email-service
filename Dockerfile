FROM golang:1.23-bullseye as builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    SERVER_MODE=prod

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application
RUN go build -o main ./cmd


# ---------- STAGE 2: Final (slim Debian) ----------
FROM debian:bullseye-slim

# Add a non-root user
RUN useradd -m -s /bin/bash asp

# Set working directory
WORKDIR /app

# Copy the built binary from the builder
COPY --from=builder /app/main .

# Change ownership
RUN chown -R asp:asp /app

# Switch to non-root user
USER asp

# Run the app
ENTRYPOINT ["./main"]
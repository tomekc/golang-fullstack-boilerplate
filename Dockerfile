# Stage 1: Build Go binary
FROM golang:1.24-alpine AS go-builder

WORKDIR /app

# Copy go module files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy server source
COPY server/ server/
COPY main.go .

# Build Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Stage 2: Final runtime image
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=go-builder /app/server .

# Create /data directory
RUN mkdir -p /data

# Expose port
EXPOSE 3000

# Run with -docker flag
ENTRYPOINT ["./server"]
CMD ["-docker"]

# Stage 1: Build frontend
FROM node:20-alpine AS frontend-builder

WORKDIR /app

# Copy package files
COPY package.json package-lock.json ./
COPY frontend/webapp/package.json frontend/webapp/

# Install dependencies
RUN npm ci

# Copy frontend source
COPY frontend/ frontend/

# Build frontend
RUN npm run build

# Stage 2: Build Go binary
FROM golang:1.24-alpine AS go-builder

WORKDIR /app

# Copy go module files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy backend source
COPY backend/ backend/
COPY main.go .

# Copy built frontend from previous stage
COPY --from=frontend-builder /app/frontend/webapp/build ./frontend/webapp/build

# Build Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Stage 3: Final runtime image
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

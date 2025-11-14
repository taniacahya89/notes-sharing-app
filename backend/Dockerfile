# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download dependencies (this will create go.sum automatically)
RUN go mod download

# Copy source code
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Copy .env file
COPY --from=builder /app/.env .

# Create directories for uploads and logs
RUN mkdir -p uploads logs

# Expose port
EXPOSE 8080

# Run application
CMD ["./main"]
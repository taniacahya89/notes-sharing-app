# Build stage
FROM golang:1.21-alpine AS builder

# Set Go proxy
ENV GOPROXY=https://goproxy.io,https://proxy.golang.org,direct
ENV CGO_ENABLED=0

# Install dependencies
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy go files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source code
COPY . .

# Build
RUN go build -ldflags="-w -s" -o main ./cmd/main.go

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary
COPY --from=builder /app/main .

# Copy .env (kalau ada, kalau tidak ada skip saja)
COPY --from=builder /app/.env .env

# Create directories
RUN mkdir -p uploads logs

EXPOSE 8080

CMD ["./main"]
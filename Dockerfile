# Build stage
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

# Install git for go mod if needed
RUN apk add --no-cache git

# Copy go.mod and go.sum first (layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source
COPY . .

# Build the binary from cmd/
RUN go build -o /bin/app ./cmd

# Runtime stage
FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /bin/app .

CMD ["./app"]

# -------------------------
# Build stage
# -------------------------
FROM golang:1.25-alpine AS builder

# Install build deps
RUN apk add --no-cache git

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o kd-gen .

# -------------------------
# Runtime stage
# -------------------------
FROM alpine:3.20

WORKDIR /app

# Add non-root user
RUN adduser -D kdgen
USER kdgen

# Copy binary
COPY --from=builder /app/kd-gen /usr/local/bin/kd-gen

# Default entrypoint
ENTRYPOINT ["kd-gen"]
CMD ["--help"]

# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Download dependencies first (caching layers)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary statically
# CGO_ENABLED=0 ensures no C libraries are linked dynamically
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api.go ./cmd/main.go

# Stage 2: Run in a minimal "distroless" image
FROM gcr.io/distroless/static-debian12

WORKDIR /

# Copy only the binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 8080

USER nonroot:nonroot

CMD ["./main"]
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Generate swagger docs
RUN go install github.com/go-swagger/go-swagger/cmd/swagger@latest
RUN swagger generate server -A tsp-payout-service -f services/tsp-service-payout/api/swagger.yaml -t services/tsp-service-payout/generated

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o tsp-payout-service ./services/tsp-service-payout/cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/tsp-payout-service .

# Run the application
CMD ["./tsp-payout-service"] 
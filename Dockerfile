# Use official Golang image
FROM golang:1.23.5 as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Use a minimal base image
FROM alpine:latest  

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy the seed.sql file from the build stage
COPY --from=builder /app/seed.sql .

# Expose API port
EXPOSE 8080

# Start the application
CMD ["./main"]
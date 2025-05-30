# Stage 1: Build the app
FROM golang:1.23.0 as builder
WORKDIR /app

# Copy go.mod and go.sum and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy all files and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Use Alpine as base image for the final build
FROM alpine:latest
WORKDIR /app

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Copy the built binary from the previous stage
COPY --from=builder /app/main .

# Expose port for the app
EXPOSE 8080

# Run the app
CMD ["./main"]
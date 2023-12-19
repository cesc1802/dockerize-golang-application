# Start by building the application in a builder image
# Stage 1: Build the application
FROM golang:1.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Build a small image with only the compiled application
FROM alpine:latest

# Add necessary dependencies for a minimal runtime
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose necessary port (e.g., if your app is a web server)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

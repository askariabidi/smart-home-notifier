# Start from the official Go image
FROM golang:1.21-alpine

# Install build tools and SQLite driver requirements
RUN apk add --no-cache gcc musl-dev sqlite

# Set working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the app
COPY . .

# Build the Go binary
RUN go build -o smartapp ./cmd

# Expose the port your Go server runs on
EXPOSE 8080

# Start the app
CMD ["./smartapp"]

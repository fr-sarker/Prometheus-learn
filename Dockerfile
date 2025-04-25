# Use Go 1.23 image (as required by your go.mod)
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Set environment variables for Go modules
ENV GO111MODULE=on

# Copy go.mod and go.sum files to leverage Docker cache for dependencies
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod tidy && go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Run the Go binary
CMD ["./main", "start"]

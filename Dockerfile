# Use the official Golang image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o http_demo main.go

# Expose the ports
EXPOSE 9001
EXPOSE 9002

# Command to run the application
CMD ["./http_demo"]

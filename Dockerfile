# Use golang:1.23 as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the application code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o http_demo .

# Set the entrypoint to run the application
ENTRYPOINT ["./http_demo"]

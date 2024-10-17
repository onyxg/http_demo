# HTTP 1.1 vs HTTP 2 Demo

A simple demo to show the difference between HTTP 1.1 and HTTP 2.
Go's default net/http package supports HTTP 2 out of the box, however will only use HTTP 2 if the server has TLS enabled.

## Usage

Start HTTP 1.1 server on port [localhost:9001](https://localhost:9001)
```bash
go run main.go
```

Start HTTP 2 server on port [localhost:9002](https://localhost:9002)
```bash
HTTP2=true go run main.go
```
_Uses a self-signed certificate._

Run the following command to the see HTTP 1.1 / Established connections
```bash
netstat -an | grep 9001
```

Run the following command to the see the HTTP 2 / Established connections
```bash
netstat -an | grep 9002
```

## Running Tests

To run the tests, use the following command:
```bash
go test ./...
```

## Using the Makefile

A `Makefile` is provided to automate common tasks. The following targets are available:
- `make run`: Run the server
- `make test`: Run the tests

## Using the Dockerfile

A `Dockerfile` is provided to containerize the application. To build the Docker image, use the following command:
```bash
docker build -t http_demo .
```

To run the Docker container, use the following command:
```bash
docker run -p 9001:9001 -p 9002:9002 http_demo
```

## CI/CD Pipeline

A CI/CD pipeline is configured to automate testing and deployment. The pipeline will run tests and build the Docker image on each push to the repository.

# HTTP 1.1 vs HTTP 2 Demo

A simple demo to show the difference between HTTP 1.1 and HTTP 2.
Go's default net/http package supports HTTP 2 out of the box, however will only use HTTP 2 if the server has TLS enabled.

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

## Running with Docker Compose

To run the project using Docker Compose, use the following commands:

Start HTTP 1.1 server on port [localhost:9001](https://localhost:9001)
```bash
docker-compose up http1
```

Start HTTP 2 server on port [localhost:9002](https://localhost:9002)
```bash
docker-compose up http2
```

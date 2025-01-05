# HTTP 1.1 vs HTTP 2 Demo

A simple demo to show the difference between HTTP 1.1 and HTTP 2.
Go's default net/http package supports HTTP 2 out of the box, however will only use HTTP 2 if the server has TLS enabled.

## Prerequisites

- Go 1.23 or later
- Docker
- Azure CLI

## Running the application locally

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

## Building the Docker container

To build the Docker container, run the following command:
```bash
docker build -t http_demo .
```

## Deploying the container to Azure

To deploy the container to Azure, follow these steps:

1. Log in to your Azure account:
```bash
az login
```

2. Create a resource group:
```bash
az group create --name myResourceGroup --location eastus
```

3. Create a container registry:
```bash
az acr create --resource-group myResourceGroup --name myContainerRegistry --sku Basic
```

4. Log in to the container registry:
```bash
az acr login --name myContainerRegistry
```

5. Tag the Docker image:
```bash
docker tag http_demo myContainerRegistry.azurecr.io/http_demo:latest
```

6. Push the Docker image to the container registry:
```bash
docker push myContainerRegistry.azurecr.io/http_demo:latest
```

7. Deploy the container to Azure Container Instances:
```bash
az container create --resource-group myResourceGroup --name http-demo --image myContainerRegistry.azurecr.io/http_demo:latest --ports 9001 9002 --environment-variables HTTP_PORT=9001 HTTP2_PORT=9002
```

8. Check the status of the deployment:
```bash
az container show --resource-group myResourceGroup --name http-demo --query "{FQDN:ipAddress.fqdn, ProvisioningState:provisioningState}"
```

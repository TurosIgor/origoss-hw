# origoss-hw
Homework - HTTP server and infrastructure

## Server
This simple HTTP server written in **Go** responds with *"Hello, World!"*, when accessed at the root ("/") URL. The server uses the **"net/http"** package to handle the incoming request and the response, and exposes the port *8080*.

## Docker
The Dockerfile includes the building process of the server. After downloading all the necessary packages, the application gets built, then copies the generated files into a light-weight alpine image, which is then responsible for starting and running the server.

## CI/CD pipeline
The Github Actions pipeline is responsible for building and pushing the Docker image to a cloud registry
- The first stage checks out the code, using **actions/checkout@v4** action.
- The second step logs into Dockerhub using the Repository Secrets and the **docker/login-action@v3** action.
- The third and final step builds and pushes the image to the user's registry using the **docker/build-push-action@v5** action.

## Kubernetes
The server is deployed in a Kubernetes Cluster from a file named ***server-deployment.yml***. This file includes the configuration of the server's deployment, such as the container image, resource limits, exposed ports and number of replicas. The exposed port is accessed through a LoadBalancer service, deployed from the file ***server-service.yml***. This service maps the server's port to port *80*.

## Terraform
The Terraform configuration uses the Kubernetes provider to parse the Kubernetes YML files and deploy them into the locally run *docker-desktop* cluster.
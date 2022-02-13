# go-skaffold
Sample project to run golang app locally with Skaffold. It spins up two deployments in your local K8S cluster. One for the app, and one MongoDB deployment.
The database is not yet used from the app.

Skaffold is watching for code changes, and automatically rebuilds and rebuilds your containers if you change your golang source code. 
It will also automatically apply any changes you make to the .yaml files.

## Prerequisites
This sample assumes you are running Docker Desktop. To adapt to minikube, just change the context in **init.sh**. Be careful about running in the 
right context if you have any kind of production context set up in your **kubectl**.

## Getting started


```bash
# Install ingress-nginx and create namespace
./init.sh

# Build and run in dev mode
skaffold dev
```




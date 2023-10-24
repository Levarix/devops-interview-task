# devops-interview

## Preparation:
Before the interview, please ensure that you have the following software/tools installed on your laptop:

1. **Docker:** Ensure Docker is installed and running on your laptop. You should be able to build and run Docker containers.

2. **Kubectl:** Install and configure kubectl to interact with the Kubernetes cluster.

3. **Minikube or Kind:** Choose between Minikube or Kind (Kubernetes in Docker) to run a local Kubernetes cluster. Make sure you have one of these tools installed.


## Step 1: Docker image
1. Prepare Dockerfile to build and run the application in root directory `main.go`
2. Build image with `1.0.0` tag and `secret-value-app` name

**Hints:**

Go version: `1.21.2`

Command to build Go application: `go build -o my-app main.go`

How to run Go application: `./my-app`

## Step 2: Prepare Kubernetes YAML Files for Deployment of prepared application

You should prepare the necessary Kubernetes YAML files. These should include:

- Deployment configuration for the application with environment variable with name `secret` vith value from kubernetes secret.
- Service configuration if required.
- A secret configuration containing sensitive information which is: `sUp3rS3creTVa1u3#`

## Step 3: Create Kubernetes Resources

1. Create a `staging` namespace for the application:
2. Apply the secret to the namespace
3. Apply the Deployment configuration
4. Verify that the Deployment pods are running as expected

## Step 4: Test the Application

You should test the application to ensure that it can access the secret data as expected.

1. Forward a local port to the Deployment pod to access the application:

>kubectl port-forward -n your-namespace deployment-pod-name local-port:container-port

2. Access the application in a web browser or using a tool like curl to ensure it's working as expected.

*If all steps was proceeded correctly you should see green screen with `Success!` text.*

## Additional Tasks (Optional):

1. Scaling the Deployment up or down.
2. Update the deployment
3. Check logs and events

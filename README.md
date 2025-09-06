# 🎂 Birthday CLI Lambda Authorizer

This repository contains a simple AWS Lambda Authorizer written in Go for securing the **Birthday CLI**. The authorizer validates incoming requests by comparing a header value with a predefined environment variable.

---

## 🔐 How It Works

The Lambda authorizer implements a basic authentication mechanism:

- Reads a secret token from the environment variable `AUTH_TOKEN`.
- Expects the incoming request to include an `Authorization` header.
- Compares the `Authorization` header value with the `AUTH_TOKEN`.
- If they match, the request is **authorized**.
- If not, the request is **denied**.

This logic makes it ideal for simple CLI-based or internal tool authentication without relying on external identity providers.

---

## 🛠 Build Instructions

To compile the Lambda function for deployment, use the following steps:

```bash
# Set the target architecture and OS for AWS Lambda
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bootstrap main.go

# Package the binary for Lambda
zip bootstrap.zip bootstrap

The build process has been automated in the `run.sh` script.

To build the project, simply run:

./run.sh
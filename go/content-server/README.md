# Multimedia

This directory contains the implementation of a basic multimedia service using various gRPC methods with Golang. The project follows hexagonal architecture principles to separate business logic from external interfaces.

## Project Structure

- **cmd/**: Application startup code.
  - **server/**: Server configuration and startup.
    - **bucket/**: Client management for bucket service connection.
    - **database/**: Client management for database connection.
- **config/**: Application configuration.
- **iac/**: Infrastructure as code (Terraform and Docker).
  - **deployment.tf, main.tf, variables.tf**: Terraform files.
  - **Dockerfile**: Docker container.
- **internal/**: Internal and business logic.
  - **domain/**: Business logic.
    - **content/**: Multimedia content logic.
  - **adapter/**: Adapters for external interfaces.
    - **grpc/**: gRPC implementation.
    - **repository/**: Data repositories.
  - **shared/**: Shared components.
- **pkg/**: Shared packages.
- **proto/**: Protocol buffer definitions for gRPC.

## Requirements

- Golang 1.23+
- Protoc 3.21+
- MinIO (for object storage)
- MongoDB

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/edwynrrangel/music
    cd music/go/content-server
    ```

2. Install dependencies:
    ```bash
    go mod download
    ```

3. Rename the `.env.template` file to `.env` and set the appropriate values for each environment variable.

## Usage

To run the service:
```bash
go run cmd/main.go
```

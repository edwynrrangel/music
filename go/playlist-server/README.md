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
    - **playlist/**: Playlist logic.
  - **ports/**: Adapters and ports.
    - **grpc/**: gRPC implementation.
    - **repository/**: Data repositories.
- **pkg/**: Shared packages.
- **proto/**: Protocol buffer definitions for gRPC.

## Requirements

- Golang 1.21+
- Protoc 3.21+
- MinIO (for object storage)
- MongoDB

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/edwynrrangel/grpc
    cd grpc/go/multimedia
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

```
playlist-server
├─ .env
├─ .env.template
├─ .vscode
│  └─ launch.json
├─ Makefile
├─ README.md
├─ cmd
│  ├─ main.go
│  └─ server
│     ├─ database
│     │  └─ mongodb.go
│     └─ server.go
├─ config
│  └─ config.go
├─ go.mod
├─ go.sum
├─ iac
│  ├─ Dockerfile
│  ├─ deployment.tf
│  ├─ main.tf
│  ├─ outputs.tf
│  ├─ service.tf
│  └─ variables.tf
├─ internal
│  ├─ domain
│  │  ├─ content
│  │  │  ├─ dto.go
│  │  │  └─ interface.go
│  │  └─ playlist
│  │     ├─ dto.go
│  │     ├─ entity.go
│  │     ├─ interface.go
│  │     └─ usecase.go
│  └─ port
│     ├─ grpc
│     │  ├─ client
│     │  │  └─ client.go
│     │  ├─ content
│     │  │  ├─ adapter.go
│     │  │  ├─ content.pb.go
│     │  │  ├─ content_grpc.pb.go
│     │  │  └─ mapper.go
│     │  └─ playlist
│     │     ├─ adapter.go
│     │     ├─ dto.go
│     │     ├─ mapper.go
│     │     ├─ playlist.pb.go
│     │     ├─ playlist_grpc.pb.go
│     │     └─ register.go
│     └─ repository
│        └─ playlist
│           └─ repository.go
├─ pkg
│  └─ mongodb
│     └─ mongodb.go
├─ proto
│  ├─ content.proto
│  └─ playlist.proto
└─ scripts
   └─ generate-stubs.sh

```
# Multimedia

This directory contains the implementation of a basic playback service using various gRPC methods with Golang. The project follows hexagonal architecture principles to separate business logic from external interfaces.

## Project Structure

- **cmd/**: Application startup code.
  - **server/**: Server configuration and startup.
    - **bucket/**: Client management for bucket service connection.
- **config/**: Application configuration.
- **iac/**: Infrastructure as code (Terraform and Docker).
  - **deployment.tf, main.tf, variables.tf**: Terraform files.
  - **Dockerfile**: Docker container.
  - **values.yaml**: Configuration values for deployment.
- **internal/**: Internal and business logic.
  - **adapter/**: Adapters and ports.
    - **grpc/**: gRPC implementation.
      - **content/**: gRPC content service.
      - **playback/**: gRPC playback service.
  - **domain/**: Business logic.
    - **playback/**: Playback logic.
- **proto/**: Protocol buffer definitions for gRPC.
- **scripts/**: Utility scripts.
  - **generate-stubs.sh**: Script to generate gRPC stubs.

## Requirements

- Golang 1.21+
- Protoc 3.21+
- MinIO (for object storage)
- MongoDB

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/edwynrrangel/music
    cd grpc/go/playback-server
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

## License

This project is licensed under the **GNU Affero General Public License, version 3 (AGPLv3)**.

This means that:
1. You can freely use, copy, modify, and distribute this software.
2. If you make modifications and use this software as part of a public service (e.g., an API, web application, or SaaS), you are required to:
   - Release the source code of the modifications made.
   - Ensure that the modifications are available under the same terms of the AGPLv3.

You can find the full text of the license in the [LICENSE](./LICENSE) file or consult it on the official GNU site:  
[https://www.gnu.org/licenses/agpl-3.0.html](https://www.gnu.org/licenses/agpl-3.0.html).
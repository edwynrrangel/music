.PHONY: run run-server run-front build build-front clean generate-stubs help

run-server:
	@echo "Running the Go server..."
	@trap 'echo "Execution interrupted. Exiting..."; exit 0' INT; \
		go run cmd/main.go;

run-front:
	@echo "Starting React application..."
	@cd resource && npm install && trap 'echo "Execution interrupted. Exiting..."; exit 0' INT; \
		npm start

run-all:
	@echo "Running the application..."
	@trap 'echo "Execution interrupted. Exiting..."; exit 0' INT; \
		make run-server & make run-front

build: build-front
	@echo "Building the application..."
	@go build -o bin/main cmd/main.go

build-front:
	@echo "Building React application..."
	@cd resource && npm install && npm run build

clean:
	@echo "Cleaning up..."
	@cd resource && npm run clean
	@go clean

generate-stubs:
	@echo "Generating stubs..."
	@sh scripts/generate-stubs.sh

help:
	@echo "Available targets:"
	@echo "  run-server     Run the Go server"
	@echo "  run-front      Start the React application"
	@echo "  run-all        Run the application (Go and frontend)"
	@echo "  build          Build the application (Go and frontend)"
	@echo "  build-front    Build the frontend application"
	@echo "  clean          Clean up built files"
	@echo "  generate-stubs Generate stubs for the gRPC services"
	@echo "  help           Show this help message"

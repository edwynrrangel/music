.PHONY: run start-react build build-react

run: start-react
	@echo "Running the application..."
	@trap 'echo "Execution interrupted. Exiting..."; exit 0' INT; \
		{ go run cmd/main.go; }

start-react:
	@echo "Starting React application..."
	@cd resource && npm install && npm start &

build: build-react
	@echo "Building the application..."
	@go build -o bin/main cmd/main.go

build-react:
	@echo "Building React application..."
	@cd resource && npm install && npm run build

clean:
	@echo "Cleaning up..."
	@cd resource && npm run clean
	@go clean

help:
	@echo "Available targets:"
	@echo "  run         Run the application (Go and React)"
	@echo "  start-react Start the React application"
	@echo "  build       Build the application (Go and React)"
	@echo "  build-react Build the React application"
	@echo "  clean       Clean up built files"
	@echo "  help        Show this help message"

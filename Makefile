BINARY_NAME := retfsm

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building the application..."
	go build -o $(BINARY_NAME)

.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)

.PHONY: test
test:
	@echo "Running tests..."
	go test ./lexer ./parser
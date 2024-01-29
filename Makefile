.PHONY: start run run-tests

start:
	@echo "Starting project..."
	go run cmd/main.go cmd/wire_gen.go

run: start

run-tests:
	go test ./... -v

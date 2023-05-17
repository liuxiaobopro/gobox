tidy:
	@echo "Synchronizing..."
	@go mod tidy
	@echo "Done! $(shell date "+%Y-%m-%d %H:%M:%S")"

check:
	@echo "Checking..."
	@go fmt
	@golangci-lint run
	@echo "Done! $(shell date "+%Y-%m-%d %H:%M:%S")"
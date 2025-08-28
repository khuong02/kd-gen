.PHONY: fmt lint sec

sec:
	@echo "Running lint"
	@gosec ./...
	@echo "Done."

lint:
	@echo "Running lint"
	@golangci-lint run ./...
	@echo "Done."

fmt:
	@echo "Running gofmt..."
	gofmt -w .
	@echo "Done."
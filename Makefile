install: # Install dependencies
	@go mod tidy

sec: # Run security tests
	@if [ ! -f "$(GOPATH)/bin/gosec" ]; then \
		echo "Gosec not found. Installing Gosec..."; \
		go install github.com/securego/gosec/v2/cmd/gosec@latest; \
	fi
	@gosec ./...

test: # Run all unit tests
	@go test ./... -timeout 5s -cover -coverprofile=cover.out
	@go tool cover -html=cover.out -o cover.html
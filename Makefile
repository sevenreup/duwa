build:
	@go build -C src -o ../bin/chewa

dev:
	@go run src/main.go

run: build
	@./bin/chewa

lint:
	@golangci-lint run -c ./golangci.yml ./...

test:
	@go test ./... -v --cover

test-report:
	@go test ./... -v --cover -coverprofile=coverage.out
	@go tool cover -html=coverage.out
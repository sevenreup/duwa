build:
	@go build -C src -o ../bin/chewa

dev:
	@go run src/main.go

run: build
	@./bin/chewa

test:
	@go test -v ./...
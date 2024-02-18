build:
	@go build -C src -o ../bin/chewa

run: build
	@./bin/chewa

test:
	@go test -v ./...
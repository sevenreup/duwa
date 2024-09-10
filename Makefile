build:
	@go build -C ./src/cmd/duwa -o ../../../bin/duwa

build-wasm:
	@GOOS=js GOARCH=wasm tinygo build -o ../chewa-site/public/duwa.wasm -opt 1 ./src/cmd/wasm/duwa.go

build-all: build build-wasm

dev:
	@go run ./src/cmd/duwa/duwa.go

run: build
	@./bin/duwa

lint:
	@golangci-lint run -c ./golangci.yml ./...

test:
	@go test ./... -v --cover

test-report:
	@go test ./... -v --cover -coverprofile=coverage.out
	@go tool cover -html=coverage.out
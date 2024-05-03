build:
	@go build -o bin/tagl

run: build
	@./bin/tagl

test:
	@go test -v ./...

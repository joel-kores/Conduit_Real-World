build:
	@go build -o bin/conduit ./cmd/main.go

run: build
	@./bin/conduit

test:
	@go test -v ./...
build:
	@go build -o bin/conduit ./cmd/main.go

run: build
	@./bin/conduit

test:
	@go test -v ./...

gen: 
	@oapi-codegen -generate types,chi-server,spec -o internal/generated/api.gen.go -package generated api/openapi.yml
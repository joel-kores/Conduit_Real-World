build:
	@go build -o bin/conduit ./cmd/main.go

run: build
	@./bin/conduit

dev:
	@air

debug-build:
	@go build -gcflags="all=-N -l" -o bin/conduit ./cmd/main.go

debug: debug-build
	@dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./bin/conduit

debug-dev:
	@air -- dlv debug --listen=:2345 --headless=true --api-version=2 --accept-multiclient --continue ./cmd/main.go

test:
	@go test -v ./...

gen:
	@oapi-codegen -generate types,chi-server,spec -o internal/generated/api.gen.go -package generated api/openapi.yml

.PHONY: build run dev debug-build debug debug-dev test gen

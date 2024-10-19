all:
	@go build -o bin/rgol cmd/main.go

run: all
	@./bin/rgol

test:
	@go test ./...

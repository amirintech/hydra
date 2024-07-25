build:
	@go build -o ./bin/hydra main.go

run: build
	@./bin/hydra

test:
	@go test -v ./...

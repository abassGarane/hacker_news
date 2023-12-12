.PHONY: run, build

build:
	@go build -o bin/HackerNews -ldflags "-s -w" cmd/web/main.go

run: build
	./bin/HackerNews

test:
	go test --cover -v ./...

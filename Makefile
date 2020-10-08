.PHONY: build
build:
	mkdir -p ./bin
	go build -o ./bin/go2ts ./cmd/go2ts

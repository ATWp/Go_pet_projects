.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -time 30s ./ ...

	
.DEFAULT_GOAL := build
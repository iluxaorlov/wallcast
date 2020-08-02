.PHONY: build
build:
	go build -v ./cmd/wallcast

.DEFAULT_GOAL := build
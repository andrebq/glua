.PHONY: build example

build:
	go build .

example: build
	./glua example.lua

install:
	go install .

fmt:
	go fmt

.PHONY: build run

build:
	go build -o build/go-todo


run: build
	./build/go-todo
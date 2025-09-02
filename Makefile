BINARY := go-ebpf-example

all: run

generate:
	@mkdir -p gen
	go generate ./...

build: generate
	@mkdir -p build
	go build -o build/$(BINARY) .

run: build
	sudo ./build/$(BINARY)

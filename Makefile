.PHONY: build

default: build

build: fmt clean
	go build .

fmt:
	gofmt -w .

clean:
	go clean .
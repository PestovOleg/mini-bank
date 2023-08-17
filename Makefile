.PHONY: build

lint:
	golangci-lint run ./...

build:
	mkdir -p build/; go build -o build/minibank cmd/main.go

run:
	rm -rf build/ && mkdir -p build/; go build -o build/minibank cmd/main.go && build/minibank

clean:
	rm -rf build/

docker:
	docker build -t minibank:0.0.1 .
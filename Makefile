.PHONY: build

lint:
	golangci-lint run ./...

build:
	mkdir -p build/; go build -o build/minibank cmd/main.go

run:
	rm -rf build/ && mkdir -p build/
	go build -o build/minibank cmd/main.go
	CONFIG_PATH=./config/local.yaml build/minibank

clean:
	rm -rf build/

docker:
	docker build -t minibank:0.0.1 .
.PHONY: help
help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	
.PHONY: build

lint: ## lint all
	golangci-lint run ./...

test: ## run tests
	go test ./... -cover
build: ## build service
	go build -o build/minibank cmd/main.go

clean: ## clean build dir
	rm -rf backend/build/

todockerhub: ##make image 
	docker build -t pistollo/minibank:latest .

gitlog: ## show git log
	git log --pretty=format:"%H [%cd]: %an - %s" --graph --date=format:%c


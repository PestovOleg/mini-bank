.PHONY: help
help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	
.PHONY: build

lint:
	cd backend && golangci-lint run ./...

build:
	cd backend && go build -o build/minibank cmd/main.go

run:
	rm -rf backend/build/ && mkdir -p backend/build/
	cd backend && go build -o build/minibank cmd/main.go
	CONFIG_PATH=./config/local.yaml backend/build/minibank

clean:
	rm -rf backend/build/

docker:
	docker build -t minibank:0.1.0 .

todockerhub:
	docker build -t pistollo/minibank:latest .

gitlog:
	git log --pretty=format:"%H [%cd]: %an - %s" --graph --date=format:%c

migrateup:
	docker compose up -d migrate

migratedown:
	docker run -v ./backend/internal/migrations:/migrations --network mini-bank_minibank_net migrate/migrate \
    -path=/migrations/ -database postgres://${MINIBANK_USER}:${MINIBANK_PASSWORD}@db:5432/${MINIBANK_DB}?sslmode=disable up 2 
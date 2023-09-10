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
	docker build -t minibank:0.1.0 .

gitlog:
	git log --pretty=format:"%H [%cd]: %an - %s" --graph --date=format:%c

createdb:
	docker exec -it db createdb --username=postgres --owner=postgres p2pexchange

dropdb:
	docker exec -it db dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/p2pexchange?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/p2pexchange?sslmode=disable" -verbose down
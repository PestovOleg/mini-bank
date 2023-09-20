.PHONY: build

lint:
	golangci-lint run ./...

build:
	go build -o backend/build/minibank backend/cmd/main.go

run:
	rm -rf backend/build/ && mkdir -p backend/build/
	go build -o backend/build/minibank backend/cmd/main.go
	CONFIG_PATH=./config/local.yaml backend/build/minibank

clean:
	rm -rf backend/build/

docker:
	docker build -t minibank:0.1.0 .

gitlog:
	git log --pretty=format:"%H [%cd]: %an - %s" --graph --date=format:%c

createdb:
	docker exec -it db createdb --username=postgres --owner=postgres p2pexchange

dropdb:
	docker exec -it db dropdb simple_bank

migrateup:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/p2pexchange?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/p2pexchange?sslmode=disable" -verbose down
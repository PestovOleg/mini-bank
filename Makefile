.PHONY: help
help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	
.PHONY: build

auth: ## Deploy and run auth service
	export SERVICE=auth-minibank
	MINIBANK_DB=$(grep AUTH_MINIBANK_DB .env | cut -d '=' -f2)
	MINIBANK_USER=$(grep AUTH_MINIBANK_USER .env | cut -d '=' -f2)
	MINIBANK_PASSWORD=$(grep AUTH_MINIBANK_PASSWORD .env | cut -d '=' -f2)
	/bin/bash deploy.sh

user: ## Deploy and run user service
	export SERVICE=user-minibank
	MINIBANK_DB=$(grep USER_MINIBANK_DB .env | cut -d '=' -f2)
	MINIBANK_USER=$(grep USER_MINIBANK_USER .env | cut -d '=' -f2)
	MINIBANK_PASSWORD=$(grep USER_MINIBANK_PASSWORD .env | cut -d '=' -f2)
	/bin/bash deploy.sh

account: ## Deploy and run account service
	export SERVICE=account-minibank
	MINIBANK_DB=$(grep ACCOUNT_MINIBANK_DB .env | cut -d '=' -f2)
	MINIBANK_USER=$(grep ACCOUNT_MINIBANK_USER .env | cut -d '=' -f2)
	MINIBANK_PASSWORD=$(grep ACCOUNT_MINIBANK_PASSWORD .env | cut -d '=' -f2)
	/bin/bash deploy.sh

mgmt: ## Deploy and run mgmt service
	export SERVICE=mgmt-minibank
	/bin/bash deploy.sh

gitlog: ## Output git log
	git log --pretty=format:"%H [%cd]: %an - %s" --graph --date=format:%c

migrateup: ## Migrate DB with current vars: MINIBANK_USER,MINIBANK_PASSWORD,MINIBANK_DB
	docker compose up -d migrate

migratedown:
	docker run -v ./backend/internal/migrations:/migrations --network mini-bank_minibank_net migrate/migrate \
    -path=/migrations/ -database postgres://${MINIBANK_USER}:${MINIBANK_PASSWORD}@db:5432/${MINIBANK_DB}?sslmode=disable up 2 
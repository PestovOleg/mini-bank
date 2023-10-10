.PHONY: help
help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	
.PHONY: build

auth: ## Deploy and run auth service
	export SERVICE=auth-minibank 
	export MIGRATE=YES
	export MINIBANK_DB=$(grep AUTH_MINIBANK_DB .env | cut -d '=' -f2) 
	export MINIBANK_USER=$(grep AUTH_MINIBANK_USER .env | cut -d '=' -f2) 
	export MINIBANK_PASSWORD=$(grep AUTH_MINIBANK_PASSWORD .env | cut -d '=' -f2)
	/bin/bash deploy.sh

user: ## Deploy and run user service
	export SERVICE=user-minibank
	export MIGRATE=YES
	export MINIBANK_DB=$(grep USER_MINIBANK_DB .env | cut -d '=' -f2)
	export MINIBANK_USER=$(grep USER_MINIBANK_USER .env | cut -d '=' -f2)
	export MINIBANK_PASSWORD=$(grep USER_MINIBANK_PASSWORD .env | cut -d '=' -f2)
	/bin/bash deploy.sh

account: ## Deploy and run account service
	export SERVICE=account-minibank
	export MIGRATE=YES
	export MINIBANK_DB=$(grep ACCOUNT_MINIBANK_DB .env | cut -d '=' -f2)
	export MINIBANK_USER=$(grep ACCOUNT_MINIBANK_USER .env | cut -d '=' -f2)
	export MINIBANK_PASSWORD=$(grep ACCOUNT_MINIBANK_PASSWORD .env | cut -d '=' -f2)
	/bin/bash deploy.sh

mgmt: ## Deploy and run mgmt service
	export SERVICE=mgmt-minibank
	export MIGRATE=NO
	/bin/bash deploy.sh

backendbuild: ## Up all backend
	docker compose build account-minibank-green user-minibank-green auth-minibank-green mgmt-minibank-green

backendup: ## Up all backend
	docker compose up -d account-minibank-green user-minibank-green auth-minibank-green mgmt-minibank-green

backenddown: ## Down all backend
	docker compose down account-minibank-green user-minibank-green auth-minibank-green mgmt-minibank-green

swag: ##Generate Swagger documentation
	swag init --pd -g ./backend/services/auth/cmd/main.go  ./backend/services/user/cmd/main.go ./backend/services/mgmt/cmd/main.go ./backend/services/account/cmd/main.go 

gitlog: ## Output git log
	git log --pretty=format:"%H [%cd]: %an - %s" --graph --date=format:%c

migrateup: ## Migrate DB with current vars: MINIBANK_USER,MINIBANK_PASSWORD,MINIBANK_DB
	docker compose up -d migrate

migratedown:
	docker run -v ./backend/internal/migrations:/migrations --network mini-bank_minibank_net migrate/migrate \
    -path=/migrations/ -database postgres://${MINIBANK_USER}:${MINIBANK_PASSWORD}@db:5432/${MINIBANK_DB}?sslmode=disable up 2 
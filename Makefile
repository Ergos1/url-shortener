include .env

ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

MIGRATION_FOLDER=$(CURDIR)/internal/app/migrations

ifeq ($(POSTGRES_URI),)
	POSTGRES_URI := user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) sslmode=disable
endif

up-db:
	docker-compose -f deployments/psql-db/docker-compose.yml up -d 
down-db:
	docker-compose -f deployments/psql-db/docker-compose.yml up -d 

migration-create:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

migration-up:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_URI)" up

migration-down:
	$(GOPATH)/bin/goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_URI)" down
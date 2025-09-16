# Makefile for managing Goose database migrations (PostgreSQL)
GOOSE=goose
DB_DRIVER=postgres

# Load environment variables from .env file
include .env

DB_STRING="postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"
MIGRATION_DIR=./pkg/database/migrations

.PHONY: migrate up down status create

migrate:
	$(GOOSE) -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_STRING) up

up:
	$(GOOSE) -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_STRING) up

down:
	$(GOOSE) -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_STRING) down

status:
	$(GOOSE) -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_STRING) status

create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make create name=your_migration_name"; \
		exit 1; \
	fi
	$(GOOSE) -dir $(MIGRATION_DIR) create $(name) sql
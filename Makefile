include .env

CURRENT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
SHELL = /bin/sh

.PHONY: help
help:        ## Print available targets.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: run
run:         ## Start development web server.
	@make start-containers

.PHONY: stop
stop:        ## Stop development web server.
	@docker-compose --env-file .env down

.PHONY: clean
clean:       ## Clearing existing data.
	@echo "Clearing existing data"
	@docker-compose down --volumes --env-file .env up

.PHONY: start-containers
start-containers:
	@echo "Starting app containers"
	@docker-compose --env-file .env up
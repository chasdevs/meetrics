SHELL := /bin/bash

OK    := $(shell printf "\e[2D\e[32m✅ ")
WARN  := $(shell printf "\e[2D\e[33m⚠️ \e[1m")
INFO  := $(shell printf "\e[2D\e[36mℹ️ ")
ERROR := $(shell printf "\e[2D\e[31m❗ ")
END   := $(shell printf "\e[0m")

.PHONY: init image database backfill

init:
	docker-compose up -d mysql
	go run db_init.go
	# $(OK) init $(END)

image:
	docker build -t chasdevs/meetrics .
	# $(OK) image $(END)

database:
	go run db_reset.go
	# $(OK) database $(END)

backfill:
	go run backfill.go
	# $(OK) backfill $(END)

#!make
include .env

$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

dev:
	GIN_MODE=release GO111MODULE=on go get -d github.com/codegangsta/gin
	gin --appPort 5001 --port 5000

mongo:
	docker run \
		-p 27017:27017 \
		-e MONGO_INITDB_ROOT_USERNAME=${MONGO_INITDB_ROOT_USERNAME} \
		-e MONGO_INITDB_ROOT_PASSWORD=${MONGO_INITDB_ROOT_PASSWORD} \
		-d mongo

install:
	go get -v

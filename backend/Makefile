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
		--name aliashmongo \
		-d mongo

dump:
	# We extract the database
	docker exec \
		-i aliashmongo /usr/bin/mongodump \
		--username ${MONGO_INITDB_ROOT_USERNAME} \
		--password ${MONGO_INITDB_ROOT_PASSWORD} \
		--authenticationDatabase admin \
		--db ${DATABASE} --out /dump
	# then we copy in our dump folder
	docker cp aliashmongo:/dump ./dump

restore:
	docker cp ./dump aliashmongo:/dump 
	docker exec \
		-i aliashmongo /usr/bin/mongorestore \
		--username ${MONGO_INITDB_ROOT_USERNAME} \
		--password ${MONGO_INITDB_ROOT_PASSWORD} \
		--authenticationDatabase admin \
		--db ${DATABASE} --out /dump/${DATABASE}

install:
	go get -v

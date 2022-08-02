#!make
include .env

$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

dev:
	GO111MODULE=on go get -d github.com/codegangsta/gin
	gin --appPort 5001 --port 5000

install:
	go get -v

GO111MODULE:="on"
export GO111MODULE

dev:
	go get -d github.com/codegangsta/gin
	gin --appPort 5001 --port 5000

install:
	go get -v

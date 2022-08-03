package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// the common Printf but as a log
func Println(something ...interface{}) {

	log.Println(something...)
}

func GetRequest(requestURL string) []byte {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	MustOrExit(err)

	res, err := http.DefaultClient.Do(req)
	MustOrExit(err)

	resBody, err := ioutil.ReadAll(res.Body)
	MustOrExit(err)

	return resBody
}

// Must just to raise an error if there is one
func Must(err error) {

	if err != nil {
		log.Panic(err)
	}
}

// Must or Exit
func MustOrExit(err error) {

	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
}

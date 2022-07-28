package main

import "log"

// the common Printf but as a log
func Println(something ...interface{}) {

	log.Println(something...)
}

// Must just to raise an error if there is one
func Must(err error) {

	if err != nil {
		log.Panic(err)
	}
}

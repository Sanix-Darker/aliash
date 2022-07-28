package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// routes
	router.HEAD("/", EchoHandler)
	router.GET("/", InstallHandler)
	router.HEAD("/c/:id", EchoHandler)
	router.GET("/c/:id", GetHandler)

	router.Run("localhost:5001")
}

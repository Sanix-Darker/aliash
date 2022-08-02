package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// routes
	router.GET("/", InstallHandler)
	router.GET("/:id", GetHandler)

	router.Run("localhost:5001")
}

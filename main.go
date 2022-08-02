package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// routes
	router.GET("/", InstallHandler)
	router.GET("/:id", GetHandler)

	router.Run("localhost:" + os.Getenv("PORT"))
}

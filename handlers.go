package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// EmptyHandler ideal for HEAD calls
func EchoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "aliash is up and running !",
		"version": "1.0.1",
	})
}

func InstallHandler(c *gin.Context) {

	installScript, err := os.ReadFile("./install.sh")
	Must(err)

	c.Data(http.StatusOK, "text/plain", installScript)
}

func CreateAliasesHandler(c *gin.Context) {
	var as Aliases

	Must(c.BindJSON(&as))

	as.CreatedAt = time.Now()
	as.UpdatedAt = time.Now()

	if err := createAliases(&as); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":         as.Uid,
		"url":         as.Url,
		"description": as.Description,
		"created_at":  as.CreatedAt,
	})
}

func GetHandler(c *gin.Context) {

	echoScript, err := os.ReadFile("./echo.sh")
	Must(err)

	c.Data(http.StatusOK, "text/plain", echoScript)
}

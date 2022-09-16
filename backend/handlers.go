package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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

	if len(as.Content) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your content should not be empty",
		})
		return
	}

	// We should check here if the user is logged, then allow content <= 1000
	if len(as.Content) > 300 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your content is > 300",
		})
		return
	}

	as.Title = TruncateText(as.Title, 70)
	// we truncate to 300 characters for the markdown description of the
	as.Description = TruncateText(as.Description, 300)

	// We should search for the same title and the content in the database
	// and refuse to save if there is already something similar in the database
	filter := bson.M{"$and": []interface{}{bson.M{"title": as.Title}, bson.M{"content": as.Content}}}

	aliases, _ := filterAliasesBy(filter)

	if len(aliases) > 0 {
		c.JSON(http.StatusAlreadyReported, gin.H{
			"error": "An alias already exist for this title and this content",
		})
		return
	}

	as.Hash512 = ShaIt(as.Title + as.Content)
	as.Uid = TruncateText(slug.Make(as.Title), 2) + "-" + TruncateText(as.Hash512, 7)
	as.CreatedAt = time.Now()
	as.UpdatedAt = time.Now()

	if err := createAliases(&as); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"uid":        as.Uid,
		"title":      as.Title,
		"hash":       TruncateText(as.Hash512, 4),
		"created_at": as.CreatedAt,
	})
}

func SearchHandler(c *gin.Context) {
	searchText, status := c.GetQuery("q")

	if !status {
		c.JSON(http.StatusOK, gin.H{
			"aliases": []*Aliases{},
		})
		return
	}

	aliases, err := searchAliases("title", searchText)
	Must(err)
	if aliases == nil {
		aliases = []*Aliases{}
	}

	c.JSON(http.StatusOK, aliases)
}

func GetAllHandler(c *gin.Context) {
	aliases := getAllAliases()

	c.JSON(http.StatusOK, aliases)
}

func GetHandler(c *gin.Context) {
	uid := c.Param("uid")

	if strings.TrimSpace(uid) == "" {
		log.Printf("[x] uid is empty !")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You need to provide the uid parameter !",
		})
		return
	}

	filter := bson.D{{"uid", uid}}
	aliases, err := filterAliasesBy(filter)
	if err != nil || len(aliases) == 0 {
		log.Printf("%s", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No aliases found for this uid !",
		})
		return
	}

	c.Data(http.StatusOK, "text/plain", []byte(aliases[0].Content))
}

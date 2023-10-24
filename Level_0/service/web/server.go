package server

import (
	"fmt"
	"log"
	"os"

	"service/web/handler"

	"service/pkg/memorycache"

	"github.com/gin-gonic/gin"
)

func RunServer(cache *memorycache.Cache) *gin.Engine {
	filename := "../data/htmltemplate.html"
	htmlTemplate, err := readFileToString(filename)
	if err != nil {
		fmt.Println("Error reading htmltemplate file:", err)
		return nil
	}

	router := gin.Default()

	router.GET("/", handler.HelloWorldHandler)
	router.GET("/:id", func(c *gin.Context) {
		handler.GetDataHandler(c, cache, htmlTemplate)
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error while starting Gin server: %s", err.Error())
	}

	return router
}

func readFileToString(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// http://127.0.0.1:8080
// http://127.0.0.1:8080/1

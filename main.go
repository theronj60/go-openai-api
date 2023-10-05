package main

import (
	"log"
	"net/http"

	"github.com/theronj60/go-openai-api/internal/Controllers"

	"github.com/gin-gonic/gin"
	goenv "github.com/joho/godotenv"
)

func main() {

	err := goenv.Load()
	if err != nil {
		log.Fatalf("Could not load env. Err: %s", err)
	}

	router := gin.Default()

	// router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"title": "Index",
	// 	})
		c.String(200, "Success")

		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome home",
		})
	})
	// router.GET("/about", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "about.html", gin.H{
	// 		"title": "About",
	// 	})
	// })

	api := router.Group("/api")
	{
		api.POST("/financial-writer", Controllers.GetFinancialHandler)
		// api.POST("/financial-writer", Controllers.GetFinancialHandler)
	}

	http.ListenAndServe(":8885", router)
}


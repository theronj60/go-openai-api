package main

import (
	"log"
	"os"

	"github.com/theronj60/go-openai-api/internal/Controllers"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	goenv "github.com/joho/godotenv"
)

func main() {

	err := goenv.Load()
	if err != nil {
		log.Fatalf("Could not load env. Err: %s", err)
	}

	router := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = false
	config.AllowedOrigins = []string{os.Getenv("URL_ORIGIN")} // Allow only this origin or you can allow all origins with `config.AllowAllOrigins = true`
	config.AllowedMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowedHeaders = []string{"Origin", "Content-Type", "Authorization"} // Adjust headers as necessary

	router.Use(cors.New(config))


	// router.LoadHTMLGlob("templates/*")
	router.GET("/", Controllers.HomeHandler)
	// router.GET("/about", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "about.html", gin.H{
	// 		"title": "About",
	// 	})
	// })

	api := router.Group("/api")
	{
		api.POST("/financial-assistant", Controllers.GetFinancialHandler)
		// api.POST("/financial-writer", Controllers.GetFinancialHandler)
	}
	router.Run(os.Getenv("PORT"))

	// http.ListenAndServe(":8885", router)
}


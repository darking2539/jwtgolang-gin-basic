package main

import (
	"fmt"
	"go_stock/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {

	// If production Disable Here
	if os.Getenv("PROD") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		fmt.Println("Load ENV....")
	}

	router := gin.Default()
	port := os.Getenv("PORT")
	router.Static("/images", "./uploaded/images")
	router.Use(cors.Default())
	api.Setup(router)
	router.Run(":" + port)

}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Gabs-Leo/FATEC-Go-Bike-Api/controllers"
	"github.com/Gabs-Leo/FATEC-Go-Bike-Api/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//Dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	models.InitDatabase()
	router := gin.Default()

	// Routes
	router.GET("/bikes", controllers.GetBikes)
	router.GET("/bikes/:id", controllers.GetBike)
	router.POST("/bikes", controllers.CreateBike)
	router.PUT("/bikes/:id", controllers.UpdateBike)
	router.DELETE("/bikes/:id", controllers.DeleteBike)

	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}

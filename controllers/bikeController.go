package controllers

import (
	"net/http"
	"strconv"

	"github.com/Gabs-Leo/FATEC-Go-Bike-Api/models"
	"github.com/gin-gonic/gin"
)

func GetBikes(context *gin.Context) {
	var bikes []models.Bike
	models.DB.Find(&bikes)
	context.JSON(http.StatusOK, bikes)
}

func GetBike(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var bike models.Bike
	if err := models.DB.First(&bike, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Bike not found"})
		return
	}
	context.JSON(http.StatusOK, bike)
}

func CreateBike(context *gin.Context) {
	var bike models.Bike
	if err := context.ShouldBindJSON(&bike); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if bike.Price < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Price should be greater than 0."})
		return
	}
	if bike.Model == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Model can't be empty."})
		return
	}
	models.DB.Create(&bike)
	context.JSON(http.StatusCreated, bike)
}

func UpdateBike(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var bike models.Bike
	if err := models.DB.First(&bike, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Bike not found"})
		return
	}

	if err := context.ShouldBindJSON(&bike); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Save(&bike)
	context.JSON(http.StatusOK, bike)
}

func DeleteBike(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var bike models.Bike
	if err := models.DB.First(&bike, id).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Bike not found"})
		return
	}
	models.DB.Delete(&bike)
	context.JSON(http.StatusOK, gin.H{"message": "Bike deleted"})
}

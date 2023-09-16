package model

import (
	"github.com/gin-gonic/gin"
	"golobe/structure"
	"gorm.io/gorm"
	"net/http"
)

type HotelScheme struct {
	gorm.Model
	DB    *gorm.DB
	Hotel structure.Hotel
}

func (scheme *HotelScheme) GetHotelById(ctx *gin.Context) {

	var hotel structure.Hotel

	if err := scheme.DB.Preload("Rooms").Find(&hotel, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error getting hotel by id": err.Error()})
		return
	}

	if hotel.Id == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Hotel not found"})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, hotel)
}

func (scheme *HotelScheme) GetHotels(ctx *gin.Context) {

	var hotels []structure.Hotel

	if err := scheme.DB.Preload("Rooms").Find(&hotels).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error getting all hotels": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, hotels)
}

func (scheme *HotelScheme) CreateHotel(ctx *gin.Context) {

	if err := ctx.ShouldBindJSON(&scheme.Hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := scheme.DB.Create(&scheme.Hotel).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, scheme.Hotel)
}

func (scheme *HotelScheme) UpdateHotel(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateData map[string]interface{}

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := scheme.DB.First(&scheme.Hotel, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	// Используем функцию Updates для выполнения частичного обновления
	if err := scheme.DB.Model(&scheme.Hotel).Updates(updateData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &scheme.Hotel)

}

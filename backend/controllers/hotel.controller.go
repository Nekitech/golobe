package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golobe/model"
	"gorm.io/gorm"
	"net/http"
)

type HotelScheme struct {
	gorm.Model
	DB    *sql.DB
	Hotel model.Hotel
}

//func (scheme *HotelScheme) GetHotelById(ctx *gin.Context) {
//
//	var hotel model.Hotel
//
//	if err := scheme.DB.Preload("Rooms").First(&hotel, ctx.Param("id")).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error getting hotel by id": err.Error()})
//		return
//	}
//
//	ctx.IndentedJSON(http.StatusCreated, hotel)
//}

func (scheme *HotelScheme) GetHotels(ctx *gin.Context) {

	var hotels []model.Hotel

	rows, err := scheme.DB.Query("SELECT * FROM hotels")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var hotel model.Hotel
		if err := rows.Scan(
			&hotel.Id,
			&hotel.Name,
			&hotel.Stars,
			&hotel.Rating,
			&hotel.CountReviews,
			&hotel.Amenities,
			&hotel.Address,
			&hotel.PricePerNight,
			&hotel.UrlImage,
			&hotel.Freebies); err != nil {
			panic(err)
		}
		hotels = append(hotels, hotel)
	}

	//if err := scheme.DB.Preload("Rooms").Find(&hotels).Error; err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error getting all hotels": err.Error()})
	//	return
	//}

	ctx.IndentedJSON(http.StatusCreated, &hotels)
}

//
//func (scheme *HotelScheme) CreateHotel(ctx *gin.Context) {
//
//	var hotel model.Hotel
//
//	if err := ctx.ShouldBindJSON(&hotel); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	if err := scheme.DB.Create(&hotel).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": hotel})
//		return
//	}
//
//	ctx.IndentedJSON(http.StatusCreated, &hotel)
//}
//
//func (scheme *HotelScheme) UpdateHotel(ctx *gin.Context) {
//	id := ctx.Param("id")
//
//	var updateData map[string]interface{}
//
//	if err := ctx.ShouldBindJSON(&updateData); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	if err := scheme.DB.First(&scheme.Hotel, id).Error; err != nil {
//		ctx.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
//		return
//	}
//
//	if err := scheme.DB.Model(&scheme.Hotel).Updates(updateData).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, &scheme.Hotel)
//
//}
//
//func (scheme *HotelScheme) DeleteHotel(ctx *gin.Context) {
//
//	id := ctx.Param("id")
//
//	if err := scheme.DB.First(&scheme.Hotel, id).Error; err != nil {
//		ctx.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
//		return
//	}
//
//	if err := scheme.DB.Delete(&scheme.Hotel, id).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})
//}

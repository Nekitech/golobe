package controllers

import (
	"github.com/gin-gonic/gin"
	"golobe/model"
	"gorm.io/gorm"
	"net/http"
)

type RoomScheme struct {
	gorm.Model
	DB   *gorm.DB
	Room model.Room
}

func (scheme *RoomScheme) CreateRoom(ctx *gin.Context) {

	var room model.Room

	if err := ctx.ShouldBindJSON(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	if err := scheme.DB.Create(&room).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Not found hotel with this id"})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, &room)
}

func (scheme *RoomScheme) UpdateRoom(ctx *gin.Context) {

	id := ctx.Param("id")

	var updateRoom map[string]interface{}

	if err := ctx.ShouldBindJSON(&updateRoom); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := scheme.DB.First(&scheme.Room, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	if err := scheme.DB.Model(&scheme.Room).Updates(updateRoom).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, &scheme.Room)
}

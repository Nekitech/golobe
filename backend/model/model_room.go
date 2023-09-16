package model

import (
	"github.com/gin-gonic/gin"
	"golobe/structure"
	"gorm.io/gorm"
	"net/http"
)

type RoomScheme struct {
	gorm.Model
	DB   *gorm.DB
	Room structure.Room
}

func (scheme *RoomScheme) CreateRoom(ctx *gin.Context) {

	if err := ctx.ShouldBindJSON(&scheme.Room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := scheme.DB.Create(&scheme.Room).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, scheme.Room)
}

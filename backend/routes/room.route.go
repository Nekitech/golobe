package routes

import (
	"github.com/gin-gonic/gin"
	"golobe/controllers"
	"golobe/model"
	"gorm.io/gorm"
)

func RoomRoute(DB *gorm.DB, router *gin.Engine) {
	roomMethods := controllers.RoomScheme{
		Model: gorm.Model{},
		DB:    DB,
		Room:  model.Room{},
	}

	router.POST("/room", roomMethods.CreateRoom)
	router.PATCH("/room/:id", roomMethods.UpdateRoom)
}

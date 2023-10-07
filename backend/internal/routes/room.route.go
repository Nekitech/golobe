package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golobe/internal/controllers"
	"golobe/internal/model"
)

func RoomRoute(DB *sql.DB, router *gin.Engine) {
	roomMethods := controllers.RoomScheme{
		DB:   DB,
		Room: model.Room{},
	}

	router.POST("/room", roomMethods.CreateRoom)
	router.PATCH("/room/:id", roomMethods.UpdateRoom)
}

package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golobe/internal/database/model"
	"golobe/internal/services"
)

func RoomRoute(DB *sql.DB, router *gin.Engine) {
	var roomMethods services.Room = &services.RoomScheme{
		DB:   DB,
		Room: model.Room{},
	}

	router.POST("/room", roomMethods.CreateRoom)
	router.PATCH("/room/:id", roomMethods.UpdateRoom)
}

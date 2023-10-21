package handler

import (
	"github.com/gin-gonic/gin"
	"golobe/internal/services"
)

type Handlers struct {
	services *services.Services
}

func InitHandlers(services *services.Services) *Handlers {
	return &Handlers{services: services}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
	}

	api := router.Group("/api")
	{
		hotels := api.Group("/hotel")
		{
			hotels.GET("/", h.GetHotels)
			hotels.GET("/:hotelID", h.GetHotelById)
			hotels.PATCH("/:hotelID", h.UpdateHotel)
			hotels.POST("/", h.CreateHotels)
			hotels.DELETE("/:hotelID", h.DeleteHotel)

			rooms := hotels.Group(":hotelID")
			{
				rooms.POST("/room", h.CreateRoom)
				rooms.PATCH("/room/:roomID", h.RoomUpdate)
			}
		}

	}

	return router

}

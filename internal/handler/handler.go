package handler

import (
	"github.com/gin-contrib/cors"
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
	//corsConfig := cors.Config{
	//	//AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "X-Requested-With"},
	//	AllowAllOrigins:  true,
	//	AllowCredentials: true,
	//}
	router.Use(cors.Default())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp, h.CreateUserHistoryBooking)
	}

	api := router.Group("/api")
	{
		users := api.Group("/user")
		{
			users.PATCH("/:id", h.UserInfoUpdate)
		}

		hotels := api.Group("/hotel")
		{
			hotels.GET("", h.GetHotels)
			hotels.GET("/:hotelID", h.GetHotelById)
			hotels.PATCH("/:hotelID", h.UpdateHotel)
			hotels.POST("", h.CreateHotels)
			hotels.DELETE("/:hotelID", h.DeleteHotel)

			rooms := hotels.Group(":hotelID")
			{
				rooms.POST("/room", h.CreateRoom)
				rooms.PATCH("/room/:roomID", h.RoomUpdate)
			}
		}

		booking := api.Group("/booking")
		{
			booking.POST("/create", h.CreateBooking)
		}

	}

	return router

}

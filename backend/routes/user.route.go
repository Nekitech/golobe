package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golobe/controllers"
	"golobe/model"
)

func UserRoute(DB *sql.DB, router *gin.Engine) {
	userMethods := controllers.UserScheme{
		DB:   DB,
		User: model.User{},
	}

	bookingMethods := controllers.BookingScheme{
		DB:      DB,
		Booking: model.Booking{},
	}

	router.POST("/user", userMethods.CreateUser, bookingMethods.CreateUserHistoryBooking)
	router.PATCH("/user/:id", userMethods.UpdateUser)
}

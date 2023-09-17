package routes

import (
	"github.com/gin-gonic/gin"
	"golobe/controllers"
	"golobe/model"
	"gorm.io/gorm"
)

func UserRoute(DB *gorm.DB, router *gin.Engine) {
	userMethods := controllers.UserScheme{
		Model: gorm.Model{},
		DB:    DB,
		User:  model.User{},
	}

	router.POST("/user", userMethods.CreateUser)
	router.PATCH("/user/:id", userMethods.UpdateUser)
}

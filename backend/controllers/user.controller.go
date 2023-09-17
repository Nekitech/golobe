package controllers

import (
	"github.com/gin-gonic/gin"
	"golobe/model"
	"gorm.io/gorm"
)

type UserScheme struct {
	gorm.Model
	DB   *gorm.DB
	User model.User
}

func (user *UserScheme) CreateUser(ctx *gin.Context) {

	if err := ctx.ShouldBindJSON(&user.User); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := user.DB.Create(&user.User).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(201, user.User)
}

func (user *UserScheme) UpdateUser(ctx *gin.Context) {

	id := ctx.Param("id")

	var updateUser map[string]interface{}

	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := user.DB.First(&user.User, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := user.DB.Model(&user.User).Updates(updateUser).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(201, user.User)

}

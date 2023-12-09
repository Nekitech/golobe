package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) UserInfoUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateUser *map[string]interface{}

	if err := ctx.ShouldBindJSON(&updateUser); err != nil {
		ErrorHandleResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedInfo, err := h.services.User.UpdateUserInfo(id, updateUser)

	if err != nil {
		ErrorHandleResponse(ctx, http.StatusInternalServerError, "Failed to update user")
		return
	}
	ctx.IndentedJSON(201, &updatedInfo)

}

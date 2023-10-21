package handler

import (
	"github.com/gin-gonic/gin"
	"golobe/internal/database/model"
	"net/http"
)

func (h *Handlers) SignUp(ctx *gin.Context) {
	var user model.UserSignUp

	if err := ctx.BindJSON(&user); err != nil {
		ErrorHandleResponse(ctx, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Authorization.SignUpUser(&user)

	if err != nil {
		ErrorHandleResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

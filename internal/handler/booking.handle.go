package handler

import (
	"github.com/gin-gonic/gin"
	"golobe/internal/database/model"
	"net/http"
)

func (h *Handlers) CreateBooking(ctx *gin.Context) {
	var booking *model.Booking

	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ErrorHandleResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if &booking == nil {
		//ErrorHandleResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	newBooking, err := h.services.Booking.CreateBooking(booking)

	if err != nil {
		ErrorHandleResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.IndentedJSON(201, &newBooking)
}

func (h *Handlers) CreateUserHistoryBooking(ctx *gin.Context) {

	user_id, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found in context"})
		return
	}

	err := h.services.CreateUserHistoryBooking(user_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create history booking user"})
		return
	}

	ctx.IndentedJSON(201, gin.H{"success": "history success created!"})

}

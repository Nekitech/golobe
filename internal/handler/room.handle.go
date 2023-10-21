package handler

import (
	"github.com/gin-gonic/gin"
	"golobe/internal/database/model"
	"net/http"
)

func (h *Handlers) CreateRoom(ctx *gin.Context) {
	var room *model.Room

	if err := ctx.ShouldBindJSON(&room); err != nil {
		ErrorHandleResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	room, err := h.services.Rooms.CreateRoom(room)

	if err != nil {
		ErrorHandleResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &room)
}

func (h *Handlers) RoomUpdate(ctx *gin.Context) {
	id := ctx.Param("roomID")

	var room *map[string]interface{}
	if err := ctx.ShouldBindJSON(&room); err != nil {
		ErrorHandleResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedRoom, err := h.services.Rooms.UpdateRoom(id, room)

	if err != nil {
		ErrorHandleResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &updatedRoom)
}

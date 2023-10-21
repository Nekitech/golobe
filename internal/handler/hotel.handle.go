package handler

import (
	"github.com/gin-gonic/gin"
	"golobe/internal/database/model"
	"net/http"
)

func (h *Handlers) GetHotels(ctx *gin.Context) {

	hotels, err := h.services.Hotels.GetHotels()

	if err != nil {
		ErrorHandleResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.IndentedJSON(http.StatusCreated, &hotels)

}

func (h *Handlers) GetHotelById(ctx *gin.Context) {
	var id = ctx.Param("hotelID")

	hotel, err := h.services.Hotels.GetHotelById(id)

	if err != nil {
		ErrorHandleResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	ctx.IndentedJSON(http.StatusCreated, &hotel)

}

func (h *Handlers) CreateHotels(ctx *gin.Context) {
	var hotel *model.Hotel

	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ErrorHandleResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hotel, err := h.services.Hotels.CreateHotel(hotel)

	if err != nil {
		ErrorHandleResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, hotel)

}

func (h *Handlers) DeleteHotel(ctx *gin.Context) {
	id := ctx.Param("hotelID")

	err := h.services.Hotels.DeleteHotel(id)

	if err != nil {
		ErrorHandleResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})

}

func (h *Handlers) UpdateHotel(ctx *gin.Context) {

	id := ctx.Param("hotelID")

	var hotel *map[string]interface{}

	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ErrorHandleResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedHotel, err := h.services.Hotels.UpdateHotel(id, hotel)

	if err != nil {
		ErrorHandleResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, updatedHotel)

}

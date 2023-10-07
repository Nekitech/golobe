package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golobe/internal/model"
	"net/http"
)

type RoomScheme struct {
	DB   *sql.DB
	Room model.Room
}

func (scheme *RoomScheme) CreateRoom(ctx *gin.Context) {

	var room model.Room

	if err := ctx.ShouldBindJSON(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO rooms (hotel_id, type, is_view_on_city, amount_double_beds, amount_single_beds, cost_per_night) 
		values ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err := scheme.DB.QueryRow(query,
		room.HotelId,
		room.Type,
		room.IsViewOnCity,
		room.AmountDoubleBeds,
		room.AmountSingleBeds,
		room.CostPerNight,
	).Scan(&room.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert room"})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, &room)
}

func (scheme *RoomScheme) UpdateRoom(ctx *gin.Context) {

	id := ctx.Param("id")

	var room map[string]interface{}
	if err := ctx.ShouldBindJSON(&room); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE rooms SET"
	params := []interface{}{}

	i := 1

	for field, value := range room {
		if value != nil {
			query += " " + field + "=$" + fmt.Sprint(i) + ","
			params = append(params, value)
			i++
		}
	}

	fmt.Println(query)

	if query[len(query)-1] == ',' {
		query = query[:len(query)-1]
	}

	query += " WHERE id=$" + fmt.Sprint(i)
	params = append(params, id)

	_, err := scheme.DB.Exec(query, params...)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update room"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "room success updated!"})

	ctx.IndentedJSON(http.StatusCreated, &scheme.Room)
}

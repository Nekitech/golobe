package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golobe/internal/model"
	"net/http"
	"time"
)

type BookingScheme struct {
	DB      *sql.DB
	Booking model.Booking
}

func (scheme *BookingScheme) CreateBooking(ctx *gin.Context) {
	var booking model.Booking

	if err := ctx.ShouldBindJSON(&booking); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO bookings (booking_id, room_id, check_in_time, check_out_time)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	if &booking == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Booking pointer is nil"})
		return
	}

	checkInTime, _ := time.Parse(time.RFC3339, booking.CheckInTime)
	checkOutTime, _ := time.Parse(time.RFC3339, booking.CheckOutTime)

	fmt.Println(checkInTime, checkOutTime)

	err := scheme.DB.QueryRow(query,
		booking.BookingID,
		booking.RoomID,
		checkInTime,
		checkOutTime).Scan(&booking.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	ctx.IndentedJSON(201, &booking)

}

func (scheme *BookingScheme) CreateUserHistoryBooking(ctx *gin.Context) {

	var id_booking_history int

	user_id, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user_id not found in context"})
		return
	}

	query := `
		INSERT INTO history_bookings (user_id)
		VALUES ($1)
		returning id`

	err := scheme.DB.QueryRow(query,
		user_id,
	).Scan(&id_booking_history)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create history booking user"})
		return
	}

	ctx.IndentedJSON(201, gin.H{"success": &id_booking_history})
}

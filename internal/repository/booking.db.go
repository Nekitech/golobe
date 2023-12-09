package repository

import (
	"database/sql"
	"golobe/internal/database/model"
	"time"
)

type BookingDB struct {
	db *sql.DB
}

func InitBookingDB(db *sql.DB) *BookingDB {
	return &BookingDB{db: db}
}

func (repo *BookingDB) CreateBooking(booking *model.Booking) (*model.Booking, error) {

	query := `
		INSERT INTO bookings (booking_id, room_id, check_in_time, check_out_time)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	checkInTime, _ := time.Parse(time.RFC3339, booking.CheckInTime)
	checkOutTime, _ := time.Parse(time.RFC3339, booking.CheckOutTime)

	err := repo.db.QueryRow(query,
		booking.BookingID,
		booking.RoomID,
		checkInTime,
		checkOutTime).Scan(&booking.Id)

	if err != nil {
		return booking, err
	}
	return booking, err
}

func (repo *BookingDB) CreateUserHistoryBooking(userId any) error {
	var id_booking_history int

	query := `
		INSERT INTO history_bookings (user_id)
		VALUES ($1)
		returning id`

	err := repo.db.QueryRow(query,
		userId,
	).Scan(&id_booking_history)

	return err
}

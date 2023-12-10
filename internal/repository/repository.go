package repository

import (
	"database/sql"
	"golobe/internal/database/model"
)

type Authorization interface {
	CreateUser(user *model.UserSignUp) (uint, error)
}
type User interface {
	UpdateUser(id string, user *map[string]interface{}) (*model.User, error)
}

type Hotels interface {
	GetHotels(filter *map[string]interface{}) (*[]model.Hotel, error)
	GetHotelById(id string) (*model.Hotel, error)
	CreateHotel(hotel *model.Hotel) (*model.Hotel, error)
	UpdateHotel(id string, hotel *map[string]interface{}) (*model.Hotel, error)
	DeleteHotel(id string) error
}

type Rooms interface {
	CreateRoom(room *model.Room) (*model.Room, error)
	UpdateRoom(id string, room *map[string]interface{}) (*model.Room, error)
}

type Booking interface {
	CreateBooking(booking *model.Booking) (*model.Booking, error)
	CreateUserHistoryBooking(userId any) error
}

type Repositories struct {
	Authorization
	Hotels
	Rooms
	Booking
	User
}

func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Authorization: InitAuthDB(db),
		Hotels:        InitHotelDB(db),
		Rooms:         InitRoomDB(db),
		Booking:       InitBookingDB(db),
		User:          InitUserDB(db),
	}
}

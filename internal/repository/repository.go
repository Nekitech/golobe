package repository

import (
	"database/sql"
	"golobe/internal/database/model"
)

type Authorization interface {
	CreateUser(user *model.UserSignUp) (uint, error)
}

type Hotels interface {
	GetHotels() (*[]model.Hotel, error)
	GetHotelById(id string) (*model.Hotel, error)
	CreateHotel(hotel *model.Hotel) (*model.Hotel, error)
	UpdateHotel(id string, hotel *map[string]interface{}) (*model.Hotel, error)
	DeleteHotel(id string) error
}

type Rooms interface {
	CreateRoom(room *model.Room) (*model.Room, error)
	UpdateRoom(id string, room *map[string]interface{}) (*model.Room, error)
}

type Repositories struct {
	Authorization
	Hotels
	Rooms
}

func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Authorization: InitAuthDB(db),
		Hotels:        InitHotelDB(db),
		Rooms:         InitRoomDB(db),
	}
}

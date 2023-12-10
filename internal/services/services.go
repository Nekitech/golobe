package services

import (
	"golobe/internal/database/model"
	"golobe/internal/repository"
)

type Authorization interface {
	SignUpUser(user *model.UserSignUp) (uint, error)
}

type User interface {
	UpdateUserInfo(id string, user *map[string]interface{}) (*model.User, error)
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

type Services struct {
	Authorization
	Hotels
	Rooms
	Booking
	User
}

func InitServices(repos *repository.Repositories) *Services {
	return &Services{
		Authorization: InitAuthService(repos.Authorization),
		Hotels:        InitHotelService(repos.Hotels),
		Rooms:         InitRoomService(repos.Rooms),
		Booking:       InitBookingService(repos.Booking),
		User:          InitUserService(repos.User),
	}
}

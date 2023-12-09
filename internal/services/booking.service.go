package services

import (
	"golobe/internal/database/model"
	"golobe/internal/repository"
)

type BookingService struct {
	repo repository.Booking
}

func InitBookingService(repo repository.Booking) *BookingService {
	return &BookingService{repo: repo}
}

func (service *BookingService) CreateBooking(booking *model.Booking) (*model.Booking, error) {
	return service.repo.CreateBooking(booking)
}

func (service *BookingService) CreateUserHistoryBooking(userId any) error {
	return service.repo.CreateUserHistoryBooking(userId)
}

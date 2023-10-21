package services

import (
	"golobe/internal/database/model"
	"golobe/internal/repository"
)

type HotelService struct {
	repo repository.Hotels
}

func InitHotelService(repo repository.Hotels) *HotelService {
	return &HotelService{repo: repo}
}

// -- Functions --

func (service *HotelService) UpdateHotel(id string, hotel *map[string]interface{}) (*model.Hotel, error) {
	return service.repo.UpdateHotel(id, hotel)
}

func (service *HotelService) GetHotels() (*[]model.Hotel, error) {
	return service.repo.GetHotels()
}

func (service *HotelService) GetHotelById(id string) (*model.Hotel, error) {
	return service.repo.GetHotelById(id)
}

func (service *HotelService) CreateHotel(hotel *model.Hotel) (*model.Hotel, error) {
	return service.repo.CreateHotel(hotel)
}

func (service *HotelService) DeleteHotel(id string) error {
	return service.repo.DeleteHotel(id)
}

package services

import (
	"golobe/internal/database/model"
	"golobe/internal/repository"
)

type RoomService struct {
	repo repository.Rooms
}

func InitRoomService(repo repository.Rooms) *RoomService {
	return &RoomService{repo: repo}
}

func (service *RoomService) CreateRoom(room *model.Room) (*model.Room, error) {
	return service.repo.CreateRoom(room)
}

func (service *RoomService) UpdateRoom(id string, room *map[string]interface{}) (*model.Room, error) {
	return service.repo.UpdateRoom(id, room)
}

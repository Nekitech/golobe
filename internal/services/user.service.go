package services

import (
	"golobe/internal/database/model"
	"golobe/internal/repository"
)

type UserService struct {
	repo repository.User
}

func InitUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) UpdateUserInfo(id string, user *map[string]interface{}) (*model.User, error) {
	return service.repo.UpdateUser(id, user)
}

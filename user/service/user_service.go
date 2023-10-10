package user_service

import (
	user_models "github.com/Cheveo/recruiting/user/models"
	user_storage"github.com/Cheveo/recruiting/user/storage"
)

type UserService struct {
	UserStorage user_storage.Storage
}

func NewUserService(userStorage user_storage.Storage) *UserService {
	return &UserService{
		UserStorage: userStorage,
	}
}

func (service *UserService) GetUsers() ([]*user_models.User, error) {
	return service.UserStorage.GetUsers()
}
func (service *UserService) CreateUser(user *user_models.User) error {
	return service.UserStorage.CreateUser(user)
}
func (service *UserService) UpdateUser(user *user_models.User) error {
	return service.UserStorage.UpdateUser(user)
}
func (service *UserService) GetUserById(id int) (*user_models.User, error) {
	return service.UserStorage.GetUserById(id)
}
func (service *UserService) DeleteUser(id int) error {
	return service.UserStorage.DeleteUser(id)
}

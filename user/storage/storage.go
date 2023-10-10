package user_storage

import user_models "github.com/Cheveo/recruiting/user/models"

type Storage interface {
	GetUsers() ([]*user_models.User, error)
	GetUserById(id int) (*user_models.User, error)
	UpdateUser(user *user_models.User) error
	CreateUser(user *user_models.User) error
	DeleteUser(id int) error
}

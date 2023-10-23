package user_storage

import (
	user_models "github.com/Cheveo/recruiting/user/models"
	"gorm.io/gorm"
)

type UserStorage struct {
	DBPool *gorm.DB
}

func NewUserStorage(db *gorm.DB) *UserStorage {
	db.AutoMigrate(&user_models.User{}, &user_models.Profile{}, &user_models.Skill{}, &user_models.Project{})

	return &UserStorage{
		DBPool: db,
	}
}

func (storage *UserStorage) GetUsers() ([]*user_models.User, error) {
	users := []*user_models.User{}

	res := storage.DBPool.Preload("Profile.Projects").Preload("Profile.Tools").Preload("Profile.ProgrammingSkills").Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}
func (storage *UserStorage) CreateUser(user *user_models.User) error {
	res := storage.DBPool.Create(user)

	return res.Error
}
func (storage *UserStorage) UpdateUser(user *user_models.User) error {
	res := storage.DBPool.Save(user)

	return res.Error
}
func (storage *UserStorage) GetUserById(id int) (*user_models.User, error) {
	user := new(user_models.User)

	res := storage.DBPool.Find(&user, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
func (storage *UserStorage) DeleteUser(id int) error {
	res := storage.DBPool.Delete(&user_models.User{}, id)

	return res.Error
}

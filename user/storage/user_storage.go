package user_storage

import (
	"fmt"
	"os"

	user_models "github.com/Cheveo/recruiting/user/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserStorage struct {
	DBPool *gorm.DB
}

func NewUserStorage() *UserStorage {
	godotenv.Load(".env")
	databaseUrl := os.Getenv("DATABASE_URL")
	fmt.Printf("trying to connect to DB: %s", databaseUrl)
	dbpool, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err.Error())
		os.Exit(1)
	}

	dbpool.AutoMigrate(&user_models.User{})

	return &UserStorage{
		DBPool: dbpool,
	}
}

func (storage *UserStorage) GetUsers() ([]*user_models.User, error) {
	users := []*user_models.User{}

	res := storage.DBPool.Find(&users)
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

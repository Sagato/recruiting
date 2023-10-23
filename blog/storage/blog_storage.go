package blog_storage

import (
	"net/http"

	blog_models "github.com/Cheveo/recruiting/blog/models"
	"github.com/Cheveo/recruiting/errors"
	"gorm.io/gorm"
)

type blogStorage struct {
	DBPool *gorm.DB
}

func NewBlogStorage(db *gorm.DB) *blogStorage {
	db.AutoMigrate(&blog_models.Blog{})

	return &blogStorage{
		DBPool: db,
	}
}

func (storage *blogStorage) GetBlogs() ([]*blog_models.Blog, error) {
	blogs := []*blog_models.Blog{}

	res := storage.DBPool.Find(&blogs)
	if res.Error != nil {
		return nil, res.Error
	}

	return blogs, nil
}
func (storage *blogStorage) CreateBlog(blog *blog_models.Blog) error {
	res := storage.DBPool.Create(blog)

	if res.Error != nil {
		return errors.NewHttpError("malformed client request", http.StatusBadRequest)
	}

	return res.Error
}
func (storage *blogStorage) UpdateBlog(blog *blog_models.Blog) error {
	res := storage.DBPool.Save(blog)

	if res.RowsAffected == 0 {
		return errors.NewHttpError("Resource not found", http.StatusNotFound)
	}

	if res.Error != nil {
		return errors.NewHttpError("Malformed client request", http.StatusBadRequest)
	}

	return nil 
}
func (storage *blogStorage) GetBlogById(id int) (*blog_models.Blog, error) {
	blog := new(blog_models.Blog)

	res := storage.DBPool.Find(&blog, id)

	if res.RowsAffected == 0 {
		return nil, errors.NewHttpError("Resource not found", http.StatusNotFound)
	}

	if res.Error != nil {
		return nil, errors.NewHttpError("Malformed client request", http.StatusBadRequest)
	}

	return blog, nil
}
func (storage *blogStorage) DeleteBlog(id int) error {
	res := storage.DBPool.Delete(&blog_models.Blog{}, id)

	if res.RowsAffected == 0 {
		return errors.NewHttpError("Resource not found", http.StatusNotFound)
	}

	if res.Error != nil {
		return errors.NewHttpError("Couldn't delete object", http.StatusBadRequest)
	}

	return nil
}

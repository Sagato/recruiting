package blog_service

import (
	blog_models "github.com/Cheveo/recruiting/blog/models"
	blog_storage"github.com/Cheveo/recruiting/blog/storage"
)

type BlogService struct {
	BlogStorage blog_storage.Storage
}

func NewBlogService(blogStorage blog_storage.Storage) *BlogService {
	return &BlogService{
		BlogStorage: blogStorage,
	}
}

func (service *BlogService) GetBlogs() ([]*blog_models.Blog, error) {
	return service.BlogStorage.GetBlogs()
}
func (service *BlogService) CreateBlog(blog *blog_models.Blog) error {
	return service.BlogStorage.CreateBlog(blog)
}
func (service *BlogService) UpdateBlog(blog *blog_models.Blog) error {
	return service.BlogStorage.UpdateBlog(blog)
}
func (service *BlogService) GetBlogById(id int) (*blog_models.Blog, error) {
	return service.BlogStorage.GetBlogById(id)
}
func (service *BlogService) DeleteBlog(id int) error {
	return service.BlogStorage.DeleteBlog(id)
}

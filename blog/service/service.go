package blog_service

import blog_models "github.com/Cheveo/recruiting/blog/models"

type Service interface {
	GetBlogs() ([]*blog_models.Blog, error)
	GetBlogById(id int) (*blog_models.Blog, error)
	UpdateBlog(blog *blog_models.Blog) error
	CreateBlog(blog *blog_models.Blog) error
	DeleteBlog(id int) error
}

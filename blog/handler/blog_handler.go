package blog_handler

import (
	"net/http"
	"strconv"

	blog_models "github.com/Cheveo/recruiting/blog/models"
	blog_service "github.com/Cheveo/recruiting/blog/service"
	blog_storage "github.com/Cheveo/recruiting/blog/storage"
	"github.com/Cheveo/recruiting/errors"
	"github.com/Cheveo/recruiting/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type blogHandler struct {
	BlogService blog_service.Service
}

func NewBlogHandler(db *gorm.DB) *blogHandler {
	blogStorage := blog_storage.NewBlogStorage(db)
	blogService := blog_service.NewBlogService(blogStorage)
	return &blogHandler{
		BlogService: blogService,
	}
}

func (handler *blogHandler) GetBlogs(c *gin.Context) {
	blogs, err := handler.BlogService.GetBlogs()
	if err != nil {
		httperror := errors.NewHttpError("Test error", http.StatusBadRequest)
		c.Error(httperror)
		return
	}

	response := responses.NewHttpResponse(http.StatusOK, blogs)

	c.JSON(http.StatusOK, response)
}

func (handler *blogHandler) CreateBlog(c *gin.Context) {
	var blog blog_models.Blog

	err := c.BindJSON(&blog)
	if err != nil {
		http_error := errors.NewHttpError("Object is malformed", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	err = handler.BlogService.CreateBlog(&blog)
	if err != nil {
		http_error := errors.NewHttpError("Could not create resource object", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (handler *blogHandler) UpdateBlog(c *gin.Context) {
	var blog blog_models.Blog

	err := c.BindJSON(&blog)
	if err != nil {
		http_error := errors.NewHttpError("Object is malformed", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	err = handler.BlogService.UpdateBlog(&blog)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (handler *blogHandler) GetBlogById(c *gin.Context) {
	idStr := c.Param("id")
	id, conversionErr := strconv.Atoi(idStr)
	if conversionErr != nil {
		http_error := errors.NewHttpError("path param malformed", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	blog, err := handler.BlogService.GetBlogById(id)
	if err != nil {
		c.Error(err)
		return
	}

	response := responses.NewHttpResponse(http.StatusOK, blog)
	c.JSON(http.StatusOK, response)
}

func (handler *blogHandler) DeleteBlog(c *gin.Context) {
	idStr := c.Param("id")
	id, conversionErr := strconv.Atoi(idStr)
	if conversionErr != nil {
		http_error := errors.NewHttpError("path param malformed", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	err := handler.BlogService.DeleteBlog(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (handler *blogHandler) SetupRouter(r *gin.Engine) *gin.Engine {
	r.GET("/blogs", handler.GetBlogs)
	r.POST("/blogs", handler.CreateBlog)
	r.GET("/blogs/:id", handler.GetBlogById)
	r.PUT("/blogs/:id", handler.UpdateBlog)
	r.DELETE("/blogs/:id", handler.DeleteBlog)

	return r
}

package user_handler

import (
	"net/http"
	"strconv"

	"github.com/Cheveo/recruiting/errors"
	"github.com/Cheveo/recruiting/responses"
	user_models "github.com/Cheveo/recruiting/user/models"
	user_service "github.com/Cheveo/recruiting/user/service"
	user_storage "github.com/Cheveo/recruiting/user/storage"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService user_service.Service
}

func NewUserHandler() *UserHandler {
	userStorage := user_storage.NewUserStorage()
	userService := user_service.NewUserService(userStorage)
	return &UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) GetUsers(c *gin.Context) {
	users, err := handler.UserService.GetUsers()
	if err != nil {
		httperror := errors.NewHttpError("Test error", http.StatusBadRequest)
		c.Error(httperror)
		return
	}

	response := responses.NewHttpResponse(http.StatusOK, users)

	c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var user user_models.User

	err := c.BindJSON(&user)
	if err != nil {
		http_error := errors.NewHttpError("Object is malformed", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	err = handler.UserService.CreateUser(&user)
	if err != nil {
		http_error := errors.NewHttpError("Could not create resource object", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (handler *UserHandler) UpdateUser(c *gin.Context) {
	var user user_models.User

	err := c.BindJSON(&user)
	if err != nil {
		http_error := errors.NewHttpError("Object is malformed", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	err = handler.UserService.UpdateUser(&user)
	if err != nil {
		http_error := errors.NewHttpError("Could not update resource", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (handler *UserHandler) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, conversionErr := strconv.Atoi(idStr)
	if conversionErr != nil {
		http_error := errors.NewHttpError("path param malformed", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	user, err := handler.UserService.GetUserById(id)
	if err != nil {
		http_error := errors.NewHttpError("Object is malformed", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	if err != nil {
		http_error := errors.NewHttpError("Object is malformed", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	response := responses.NewHttpResponse(http.StatusOK, user)
	c.JSON(http.StatusOK, response)
}

func (handler *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, conversionErr := strconv.Atoi(idStr)
	if conversionErr != nil {
		http_error := errors.NewHttpError("path param malformed", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	err := handler.UserService.DeleteUser(id)
	if err != nil {
		http_error := errors.NewHttpError("Couldn't delete object", http.StatusBadRequest)
		c.Error(http_error)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (handler *UserHandler) SetupRouter(r *gin.Engine) *gin.Engine {
	r.GET("/users", handler.GetUsers)
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUserById)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	return r
}

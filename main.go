package main

import (
	"github.com/Cheveo/recruiting/middlewares"
	user_handler "github.com/Cheveo/recruiting/user/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middlewares.ErrorHandler())
	uh := user_handler.NewUserHandler()
	uh.SetupRouter(r)

	r.Run(":3001")
}

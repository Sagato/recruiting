package main

import (
	"fmt"
	"os"

	"github.com/Cheveo/recruiting/contact"
	"github.com/Cheveo/recruiting/email"
	"github.com/Cheveo/recruiting/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.ErrorHandler())

	mailService := email.NewGoMailService()
	ch := contact.NewContactUsHandler(mailService)
	ch.SetupRouter(r)

	return r
}
func main() {
	r := SetupRouter()
	r.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")))
}

package main

import (
	"fmt"
	"os"

	blog_handler "github.com/Cheveo/recruiting/blog/handler"
	"github.com/Cheveo/recruiting/contact"
	"github.com/Cheveo/recruiting/db"
	"github.com/Cheveo/recruiting/email"
	"github.com/Cheveo/recruiting/middlewares"
	user_handler "github.com/Cheveo/recruiting/user/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter() *gin.Engine {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("environment could not be loaded: ", err.Error())
	}
	databaseUrl := os.Getenv("DATABASE_URL")
	database := db.NewDatabase()
	db := database.CreateDB(databaseUrl)

	r := gin.Default()
	r.Use(middlewares.ErrorHandler())

	uh := user_handler.NewUserHandler(db)
	uh.SetupRouter(r)

	bh := blog_handler.NewBlogHandler(db)
	bh.SetupRouter(r)

	mailService := email.NewGoMailService()
	ch := contact.NewContactUsHandler(mailService)
	ch.SetupRouter(r)

	return r
}
func main() {
	r := SetupRouter()
	r.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")))
}

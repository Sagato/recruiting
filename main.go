package main

import (
	"fmt"
	"os"

	"github.com/Cheveo/recruiting/contact"
	"github.com/Cheveo/recruiting/db"
	"github.com/Cheveo/recruiting/email"
	"github.com/Cheveo/recruiting/middlewares"
	user_handler "github.com/Cheveo/recruiting/user/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	appPort := os.Getenv("PORT")
	godotenv.Load(".env")
	databaseUrl := os.Getenv("DATABASE_URL")
	database := db.NewDatabase()
	db := database.CreateDB(databaseUrl)

	r := gin.Default()
	r.Use(middlewares.ErrorHandler())

	uh := user_handler.NewUserHandler(db)
	uh.SetupRouter(r)


	mailService := email.NewGoMailService()
	ch := contact.NewContactUsHandler(mailService)
	ch.SetupRouter(r)

	r.Run(fmt.Sprintf("0.0.0.0:%s", appPort))
}

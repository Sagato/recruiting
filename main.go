package main

import (
	"os"

	"github.com/Cheveo/recruiting/db"
	"github.com/Cheveo/recruiting/middlewares"
	user_handler "github.com/Cheveo/recruiting/user/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	databaseUrl := os.Getenv("DATABASE_URL")
	database := db.NewDatabase()
	db := database.CreateDB(databaseUrl)

	r := gin.Default()
	r.Use(middlewares.ErrorHandler())
	uh := user_handler.NewUserHandler(db)
	uh.SetupRouter(r)

	r.Run(":3001")
}

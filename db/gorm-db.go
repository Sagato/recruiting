package db 

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) CreateDB(databaseUrl string) *gorm.DB {
		fmt.Printf("trying to connect to DB: %s", databaseUrl)
	database, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err.Error())
		os.Exit(1)
	}

	return database 
}

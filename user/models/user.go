package user_models

import (
	"fmt"

	"github.com/jackc/pgx/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

func (user *User) ScanRow(row pgx.Row) error {
	err := row.Scan(&user.ID)
	if err != nil {
		fmt.Printf("Scan error single row: %s", err.Error())
	}
	return err
}
func (user *User) ScanRows(rows pgx.Rows) error {
	err := rows.Scan(&user.ID)
	if err != nil {
		fmt.Printf("Scan User error: %s", err.Error())
	}

	return err
}

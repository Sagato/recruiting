package blog_models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Content string `json:"content"`
}

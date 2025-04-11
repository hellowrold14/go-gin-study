package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (User) TableName() string {
	return "user"
}

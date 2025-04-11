package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Age   int    `json:"age" form:"age"`
	Email string `json:"email" form:"email"`
}

func (User) TableName() string {
	return "user"
}

package v2

import (
	"github.com/cody/go-demo1/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUsers(c *gin.Context) {
	users := []User{
		{ID: 1, Name: "Alice", Age: 66},
		{ID: 2, Name: "Tom", Age: 88},
	}
	utils.OK(users, c)
}

package v1

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
		{ID: 1, Name: "Alice", Age: 25},
		{ID: 2, Name: "Tom", Age: 25},
	}
	utils.OK(users, c)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	utils.OK(id, c)
}

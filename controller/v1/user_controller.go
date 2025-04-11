package v1

import (
	"github.com/cody/go-demo1/common"
	"github.com/cody/go-demo1/model"
	"github.com/cody/go-demo1/service"
	"github.com/cody/go-demo1/utils"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	user := &model.User{}
	err := c.ShouldBindQuery(user)
	if err != nil {
		return
	}
	pageParam := common.NewPageParam()
	err = c.ShouldBindQuery(pageParam)
	if err != nil {
		return
	}
	pageResult := service.NewUserService().GetAllUsersByGPlus(pageParam, user)
	utils.OK(pageResult, c)
}

func CreateUser(c *gin.Context) {
	user := &model.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		utils.Fail(400, "Invalid request", c)
		return
	}
	userService := service.NewUserService()
	err := userService.CreateUser(user)
	if err != nil {
		utils.Fail(500, "Failed to create user", c)
		return
	}
	utils.OK(user.ID, c)
}

func DeleteUser(c *gin.Context) {
	service.NewUserService().DeleteUserById(c.Param("id"))
	utils.OK(nil, c)
}

package router

import (
	v1 "github.com/cody/go-demo1/controller/v1"
	v2 "github.com/cody/go-demo1/controller/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/ping", ping)
	groupV1 := r.Group("/v1")
	{
		groupV1.GET("/users", v1.GetUsers)
		groupV1.DELETE("/users/:id", v1.DeleteUser)
	}

	groupV2 := r.Group("/v2")
	{
		groupV2.GET("/users", v2.GetUsers)
	}
}

func ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}

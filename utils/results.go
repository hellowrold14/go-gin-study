package utils

import "github.com/gin-gonic/gin"

func JSON(code int, data any, message string, c *gin.Context) {
	c.JSON(code, gin.H{"code": code, "data": data, "message": message})
}

func OK(data any, c *gin.Context) {
	JSON(0, data, "success", c)
}

func Fail(code int, message string, c *gin.Context) {
	JSON(code, nil, message, c)
}

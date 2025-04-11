package main

import (
	"github.com/cody/go-demo1/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

var rpcAddr string

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	rpcAddr = os.Getenv("RPC_ADDR")
}

func main() {
	gin.SetMode(gin.DebugMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	router.InitRouter(engine)
	// 设置路由
	err := engine.Run()
	if err != nil {
		return
	}
}

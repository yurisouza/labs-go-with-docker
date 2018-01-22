package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yurisouza/labs-go-with-docker/webapi"
)

func main() {
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	webapi.Setup(router)
	router.Run(":3000")
}

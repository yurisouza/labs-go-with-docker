package webapi

import (
	"github.com/gin-gonic/gin"
)

//Setup http controllers from Gin Router
func Setup(engine *gin.Engine) {
	serveUsuarioResource(engine)
}

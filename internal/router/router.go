package router

import (
	"github.com/gin-gonic/gin"
)

func New(engine *gin.Engine) *gin.Engine {
	engine.GET(EnvRoute, GetEnv)
	return engine
}

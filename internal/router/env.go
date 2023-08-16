package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetEnv(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, os.Environ())
}

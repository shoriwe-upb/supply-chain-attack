package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnyEcho(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/plain", []byte(ctx.RemoteIP()))
}

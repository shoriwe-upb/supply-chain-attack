package router

import (
	"fmt"
	"net"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	// NOTE: Never do this on real test, this is just for smaller stdout
	gin.SetMode(gin.ReleaseMode)
}

func newExpect(t *testing.T) (expect *httpexpect.Expect, l net.Listener) {
	var err error
	l, err = net.Listen("tcp", "127.0.0.1:0")
	assert.Nil(t, err)

	engine := New(gin.New())
	engine.Use(gin.Recovery())
	go engine.RunListener(l)

	expect = httpexpect.Default(t, fmt.Sprintf("http://%s", l.Addr()))

	return expect, l
}

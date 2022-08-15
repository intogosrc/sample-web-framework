package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/intogosrc/sample-web-framework/conf"
)

var (
	webEngine *gin.Engine
)

func StartServer() error {
	webEngine = gin.Default()
	webEngine.RedirectFixedPath = true

	e := webEngine.Run(fmt.Sprintf("%s:%d", conf.ConfigInst.Http.Host, conf.ConfigInst.Http.Port))
	return e
}

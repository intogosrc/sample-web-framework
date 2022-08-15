package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/intogosrc/sample-web-framework/conf"
)

var (
	webEngine *gin.Engine
)

func StartServer() error {
	m := new(Middleware)
	apiCtrl := new(ApiController)

	webEngine = gin.Default()
	webEngine.RedirectFixedPath = true

	webEngine.Use(m.Cors)

	// group
	authGroup := webEngine.Group("auth")
	authGroup.Use(m.CheckAuth)
	authGroup.Handle(http.MethodGet, "test", apiCtrl.Test)

	e := webEngine.Run(fmt.Sprintf("%s:%d", conf.ConfigInst.Http.Host, conf.ConfigInst.Http.Port))
	return e
}

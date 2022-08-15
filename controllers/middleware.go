package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/intogosrc/sample-web-framework/conf"
	"github.com/intogosrc/sample-web-framework/tools/rpchelper"
)

type Middleware struct {
	BaseController
}

func (ctrl *Middleware) CheckAuth(c *gin.Context) {
	authToken := rpchelper.RequestParameterString(c, "auth_token")
	if strings.Compare(authToken, conf.ConfigInst.AuthToken) != 0 {
		ctrl.JsonError(c, conf.ApiCodeNoAuth, "wrong token")
		c.Abort()
		return
	}

	c.Next()
}

func (ctrl *Middleware) Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Next()
}

package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/intogosrc/sample-web-framework/conf"
	"github.com/intogosrc/sample-web-framework/models"
	"github.com/intogosrc/sample-web-framework/tools/log"
)

type ApiController struct {
	BaseController // compose base BaseController, so we can use basic functions (JsonSuccess...) of BaseController
}

func (ctrl *ApiController) Test(c *gin.Context) {
	test := new(models.TestModel)
	test.Name = fmt.Sprintf("name_%d", time.Now().Unix())
	err := test.Save()
	if err != nil {
		log.LogInfo("save test failed %s", err.Error)
		ctrl.JsonError(c, conf.ApiCodeErrMsg, "inner error 1000")
		return
	}

	allData, err := models.FindAllTest()
	if err != nil {
		log.LogInfo("find all failed %s", err.Error())
		ctrl.JsonError(c, conf.ApiCodeErrMsg, "inner error 1001")
		return
	}

	oneData, err := models.FindTest(1)
	if err != nil {
		log.LogInfo("find one failed %s", err.Error())
		ctrl.JsonError(c, conf.ApiCodeErrMsg, "inner error 1002")
		return
	}

	ctrl.JsonSuccess(c, map[string]interface{}{
		"all": allData,
		"one": oneData,
	})
}

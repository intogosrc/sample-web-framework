package main

import (
	"github.com/intogosrc/sample-web-framework/controllers"
	"github.com/intogosrc/sample-web-framework/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// main 函数是整个程序的入口，程序将从这里执行
func main() {
	var err error

	// 首先初始化 ORM 句柄，连接到数据库
	err = models.InitOrm()
	if err != nil {
		panic(err)
	}

	// 启动 WEB 服务
	err = controllers.StartServer()
	if err != nil {
		panic(err)
	}

}

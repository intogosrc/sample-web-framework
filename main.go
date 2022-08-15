package main

import (
	"fmt"
	"os"

	"github.com/intogosrc/sample-web-framework/conf"
	"github.com/intogosrc/sample-web-framework/controllers"
	"github.com/intogosrc/sample-web-framework/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// main 函数是整个程序的入口，程序将从这里执行
func main() {
	var err error

	if len(os.Args) < 2 {
		fmt.Println("usage: ./build/FrameworkTest ./build/conf.json")
		return
	}

	// 解析配置文件
	err = conf.InitConfig(os.Args[1])
	if err != nil {
		panic(err)
	}

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

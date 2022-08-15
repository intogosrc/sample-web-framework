package main

import (
	"fmt"
	"os"

	"github.com/intogosrc/sample-web-framework/conf"
	"github.com/intogosrc/sample-web-framework/controllers"
	"github.com/intogosrc/sample-web-framework/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// application begins at here
func main() {
	var err error

	if len(os.Args) < 2 {
		fmt.Println("usage: ./build/FrameworkTest ./build/conf.json")
		return
	}

	// parse config file
	err = conf.InitConfig(os.Args[1])
	if err != nil {
		panic(err)
	}

	// init ORM handle
	err = models.InitOrm()
	if err != nil {
		panic(err)
	}

	// start WEB server
	err = controllers.StartServer()
	if err != nil {
		panic(err)
	}

}

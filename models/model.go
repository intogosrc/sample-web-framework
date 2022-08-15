package models

import (
	"fmt"
	"strconv"

	"github.com/intogosrc/sample-web-framework/conf"
	"github.com/jinzhu/gorm"
)

var (
	// global ORM handle, it effect the package
	dbHandle *gorm.DB
)

func InitOrm() error {
	//"user:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local"
	myHost := conf.ConfigInst.DB.Host
	myUser := conf.ConfigInst.DB.User
	myPwd := conf.ConfigInst.DB.Password
	myDb := conf.ConfigInst.DB.Name
	myPort := strconv.FormatInt(int64(conf.ConfigInst.DB.Port), 10)

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", myUser, myPwd, myHost, myPort, myDb)

	var err error
	dbHandle, err = gorm.Open("mysql", connStr)
	if err != nil {
		return err
	}

	return nil
}

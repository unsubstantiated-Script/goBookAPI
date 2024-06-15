package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goBookAPI/pkg/utils"
)

var (
	db *gorm.DB
)

func Connect() {

	mysqlUser := utils.GoDotEnvVariable("MYSQL_USER")
	mysqlPassword := utils.GoDotEnvVariable("MYSQL_PASSWORD")
	mysqlDB := utils.GoDotEnvVariable("MYSQL_DB")

	args := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlDB)

	d, err := gorm.Open("mysql", args)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

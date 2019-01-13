package database

import (
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var Db *gorm.DB

func Connect() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("数据库连接失败")
	}
	db.SingularTable(true)

	Db = db

}

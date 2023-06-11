package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//docker run -d -p 33060:3306 --name mysql-db -e MYSQL_ROOT_PASSWORD=secret mysql

func ConnectionMysql() {
	dsn := "root:@tcp(127.0.0.1:3306)/tlr?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Mysql Connected")
	}
}

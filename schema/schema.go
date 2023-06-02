package schema

import (
	"fmt"

	"github.com/Mau005/KrayManagerUpdate/tools"
)

type MysqlAttr struct {
	Host     string
	Port     string
	User     string
	Password string
	DataBase string
}

func GetDataMysql(path string) {
	data := tools.ConvertString(tools.OpenJson(path), "MYSQL")
	var mysql MysqlAttr
	mysql.Host = data["Host"]
	mysql.Port = data["Port"]
	mysql.User = data["User"]
	mysql.Password = data["Password"]
	mysql.DataBase = data["DataBase"]
	fmt.Println(mysql)

}

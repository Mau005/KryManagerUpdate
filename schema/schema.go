package schema

import (
	"database/sql"
	"fmt"

	"github.com/Mau005/KryManagerUpdate/tools"
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	Host     string
	Port     string
	User     string
	Password string
	DataBase string
}

/* connect bd, return conection */
func (m *mysql) connectBD() (db *sql.DB, e error) {
	host := fmt.Sprintf("tcp(%s:%s)", m.Host, m.Port)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", m.User, m.Password, host, m.DataBase))
	if err != nil {
		return nil, err
	}
	return db, nil
}

/* Preparing execute sentencies, return err */
func (m *mysql) Execute(query string) error {
	conn, err := m.connectBD()
	if err != nil {
		return err
	}
	defer conn.Close()
	err = conn.Ping()
	if err != nil {
		return err
	}

	execute_query, err := conn.Prepare(query)
	if err != nil {
		return err
	}
	defer execute_query.Close()

	_, err = execute_query.Exec()
	if err != nil {
		return err
	}

	return nil
}

/*Generate attribute for mysql */
func getDataMysql(path string) *mysql {
	data := tools.ConvertString(tools.OpenJson(path), "MYSQL")
	var mysql mysql
	mysql.Host = data["Host"]
	mysql.Port = data["Port"]
	mysql.User = data["User"]
	mysql.Password = data["Password"]
	mysql.DataBase = data["DataBase"]
	return &mysql
}

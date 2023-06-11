package main

import (
	"fmt"
	"net/http"

	"github.com/Mau005/KryManagerUpdate/db"
	"github.com/Mau005/KryManagerUpdate/models"
	"github.com/Mau005/KryManagerUpdate/routes"
	"github.com/gorilla/mux"
)

var appName = "KrayManagerUpdate"
var version = "1.0"

func NewServerWeb() {
	db.ConnectionMysql()
	//db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.Accounts{})
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/accounts", routes.GetAccountsHandler).Methods("GET")
	r.HandleFunc("/account/{id}", routes.GetAccountHandler).Methods("GET")
	r.HandleFunc("/account", routes.PostAccountHandler).Methods("POST")
	r.HandleFunc("/account/{id}", routes.DeleteAccountHandler).Methods("DELETE")

	//taks
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks", routes.CreateTasksHandler).Methods("POST")

	http.ListenAndServe("192.168.80.1:8000", r)
}
func WelComeTerminal() {
	var welcome string

	welcome = `
	_     _  ______ _______ __   __
	|____/  |_____/ |_____|   \_/  
	|    \_ |    \_ |     |    |   
								   
	_______ _______ __   _ _______  ______ _______  ______
	|  |  | |_____| | \  | |_____| |  ____ |______ |_____/
	|  |  | |     | |  \_| |     | |_____| |______ |    \_
														  
	Website: https://github.com/mau005 
	%s %s %s
	Created by Mau
	`
	fmt.Printf(welcome, appName, version, "\n")
}

func main() {
	WelComeTerminal()
	NewServerWeb()

}

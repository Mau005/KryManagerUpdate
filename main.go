package main

import (
	"fmt"
	"net/http"

	controllers "github.com/Mau005/KryManagerUpdate/crontrollers"
	"github.com/Mau005/KryManagerUpdate/db"
	"github.com/Mau005/KryManagerUpdate/middleware"
	"github.com/Mau005/KryManagerUpdate/models"
	"github.com/Mau005/KryManagerUpdate/routes"
	"github.com/Mau005/KryManagerUpdate/utils/auth"
	"github.com/gorilla/mux"
)

var appName = "KrayManagerUpdate"
var version = "1.0"

func NewServerWeb() {
	db.ConnectionMysql()
	//db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.Accounts{})
	r := mux.NewRouter()
	r.Use(middleware.CommonMiddleware)
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	s := r.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	s.HandleFunc("/admin", routes.AdminHandler)
	s.HandleFunc("/accounts", routes.GetAccountsHandler).Methods("GET")
	s.HandleFunc("/account/{id}", routes.GetAccountHandler).Methods("GET")
	s.HandleFunc("/account", routes.PostAccountHandler).Methods("POST")
	s.HandleFunc("/account/{id}", routes.DeleteAccountHandler).Methods("DELETE")

	//taks
	s.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	s.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	s.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	s.HandleFunc("/tasks", routes.CreateTasksHandler).Methods("POST")

	http.ListenAndServe(":8000", r)
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

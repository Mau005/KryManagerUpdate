package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/Mau005/pkg/server"
)

var appName = "KrayManagerUpdate"
var version = "1.0"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Saludos, bienvenidos a %s!", r.URL.Path[1:])
}

func NewServerWeb() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router))
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
	NewServerWeb()
}

func ExecuteServer() {
	var path string
	path = "data/Server/theforgottenserver-x64.exe"

	go func() {
		out, err := exec.Command(path).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}()
	time.Sleep(10 * time.Second)
}

func main() {
	WelComeTerminal()

}

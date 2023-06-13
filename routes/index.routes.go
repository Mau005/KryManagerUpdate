package routes

import (
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	template.Execute(w, nil)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/admin.html")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	template.Execute(w, nil)
}

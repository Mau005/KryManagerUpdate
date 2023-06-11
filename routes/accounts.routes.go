package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Mau005/KryManagerUpdate/db"
	"github.com/Mau005/KryManagerUpdate/models"
	"github.com/gorilla/mux"
)

func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	var account []models.Accounts
	db.DB.Find(&account)
	json.NewEncoder(w).Encode(&account)
}

func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account models.Accounts
	params := mux.Vars(r)
	db.DB.First(&account, params["id"])

	if account.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}
	db.DB.Model(&account).Association("Tasks").Find(&account.ID)
	json.NewEncoder(w).Encode(&account)

}

func PostAccountHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Accounts

	json.NewDecoder(r.Body).Decode(&user) //son los datos que envia el cliente
	createUser := db.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	var user models.Accounts
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Found"))
		return
	}

	db.DB.Delete(&user)
	w.WriteHeader(http.StatusOK)
}

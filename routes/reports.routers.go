package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Mau005/KryManagerUpdate/db"
	"github.com/Mau005/KryManagerUpdate/models"
	"github.com/gorilla/mux"
)

func GetReportsHandler(w http.ResponseWriter, r *http.Request) {
	var reports []models.Reports
	db.DB.Find(&reports)
	json.NewEncoder(w).Encode(&reports)
}

func GetReportHandler(w http.ResponseWriter, r *http.Request) {
	var reports models.Reports
	params := mux.Vars(r)
	db.DB.First(&reports, params["id"])

	if reports.ID == 0 {
		w.WriteHeader(http.StatusNotFound)

		w.Write([]byte("No se encuentra Reporte con ese indice"))
		return
	}
	json.NewEncoder(w).Encode(reports)
}

func PostReportHandler(w http.ResponseWriter, r *http.Request) {
	var reports models.Accounts

	json.NewDecoder(r.Body).Decode(&reports) //son los datos que envia el cliente
	createReports := db.DB.Create(&reports)
	err := createReports.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&reports)
}

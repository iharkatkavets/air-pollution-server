package main

import (
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"sensor/cmd/api/service"
)

func (app *application) CreateMeasurement(w http.ResponseWriter, r *http.Request) {
	var req service.CreateMeasurementReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		app.errorLog.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := app.service.CreateMeasurement(&req); err != nil {
		app.errorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) GetAllMeasurements(w http.ResponseWriter, r *http.Request) {
	readings, err := app.service.GetAllMeasurements()
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(readings)
	if err != nil {
		app.errorLog.Println("Failed to encode JSON:", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

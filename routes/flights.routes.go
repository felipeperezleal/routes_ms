package routes

import (
	"encoding/json"
	"net/http"

	"github.com/felipeperezleal/routes_ms/db"
	"github.com/felipeperezleal/routes_ms/models"
	"github.com/gorilla/mux"
)

func GetFlightsHandler(w http.ResponseWriter, r *http.Request) {
	var flights []models.Flight
	db.DB.Find(&flights)
	json.NewEncoder(w).Encode(flights)
}
func GetFlightHandler(w http.ResponseWriter, r *http.Request) {
	var flight models.Flight
	params := mux.Vars(r)
	db.DB.First(&flight, params["id"])
	if flight.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}
	json.NewEncoder(w).Encode(params["id"])
}
func PostFlightHandler(w http.ResponseWriter, r *http.Request) {
	var flight models.Flight
	json.NewDecoder(r.Body).Decode(&flight)

	createdUser := db.DB.Create(&flight)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(flight)
}
func DeleteFlightHandler(w http.ResponseWriter, r *http.Request) {
	var flight models.Flight
	params := mux.Vars(r)
	db.DB.First(&flight, params["id"])

	if flight.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	db.DB.Unscoped().Delete(&flight)
	w.WriteHeader(http.StatusOK)
}

func UpdateFlightHandler(w http.ResponseWriter, r *http.Request) {
    var flight models.Flight
    params := mux.Vars(r)
    db.DB.First(&flight, params["id"])

    if flight.ID == 0 {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("Flight not found"))
        return
    }
    updatedFlight := models.Flight{}
    if err := json.NewDecoder(r.Body).Decode(&updatedFlight); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid request data"))
        return
    }
    flight.Origin = updatedFlight.Origin
    flight.Destination = updatedFlight.Destination
    flight.Duration = updatedFlight.Duration
    flight.Price = updatedFlight.Price

    db.DB.Save(&flight)
    json.NewEncoder(w).Encode(flight)
}

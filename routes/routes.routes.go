package routes

import (
	"encoding/json"
	"net/http"

	"github.com/felipeperezleal/routes_ms/db"
	"github.com/felipeperezleal/routes_ms/models"
	"github.com/gorilla/mux"
)

func GetRoutesHandler(w http.ResponseWriter, r *http.Request) {
	var routes []models.Routes
	db.DB.Find(&routes)
	json.NewEncoder(w).Encode(&routes)
}

func PostRouteHandler(w http.ResponseWriter, r *http.Request) {
	var route models.Routes
	json.NewDecoder(r.Body).Decode(&route)
	createdRoute := db.DB.Create(&route)
	err := createdRoute.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&route)
}

func GetRouteHandler(w http.ResponseWriter, r *http.Request) {
	var route models.Routes
	params := mux.Vars(r)

	db.DB.First(&route, params["id"])

	if route.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No se encontró la ruta"))
		return
	}

	json.NewEncoder(w).Encode(&route)

}

func DeleteRoutesHandler(w http.ResponseWriter, r *http.Request) {
	var route models.Routes
	params := mux.Vars(r)

	db.DB.First(&route, params["id"])

	if route.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No se encontró la ruta"))
		return
	}

	db.DB.Unscoped().Delete(&route)
	w.WriteHeader(http.StatusNoContent)
}

func UpdateRouteHandler(w http.ResponseWriter, r *http.Request) {
	var route models.Routes
	params := mux.Vars(r)
	db.DB.First(&route, params["id"])

	if route.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No se encontró la ruta"))
		return
	}

	updatedRoute := models.Routes{}
	if err := json.NewDecoder(r.Body).Decode(&updatedRoute); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Datos de solicitud inválidos"))
		return
	}

	route.NumNodes = updatedRoute.NumNodes
	route.Ordering = updatedRoute.Ordering

	db.DB.Save(&route)
	json.NewEncoder(w).Encode(&route)
}

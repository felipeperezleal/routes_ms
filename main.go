package main

import (
	"net/http"

	"github.com/felipeperezleal/routes_ms/db"
	"github.com/felipeperezleal/routes_ms/models"
	"github.com/felipeperezleal/routes_ms/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Routes{})

	startServer()
}

func startServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/routes", routes.GetRoutesHandler).Methods("GET")
	r.HandleFunc("/routes", routes.PostRouteHandler).Methods("POST")
	r.HandleFunc("/routes/{id}", routes.UpdateRouteHandler).Methods("PUT")
	r.HandleFunc("/routes/{id}", routes.GetRouteHandler).Methods("GET")
	r.HandleFunc("/routes/{id}", routes.DeleteRoutesHandler).Methods("DELETE")

	http.ListenAndServe(":8081", r)
}

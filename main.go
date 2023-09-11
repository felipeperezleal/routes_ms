package main

import (
	"fmt"
	"net/http"

	"github.com/felipeperezleal/routes_ms/db"
	"github.com/felipeperezleal/routes_ms/models"
	"github.com/felipeperezleal/routes_ms/routes"
	"github.com/felipeperezleal/routes_ms/src"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Flight{})
	db.DB.AutoMigrate(models.RouteGraph{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/flights", routes.GetFlightsHandler).Methods("GET")
	r.HandleFunc("/flights", routes.GetFlightsHandler).Methods("POST")
	r.HandleFunc("/flights/{id}", routes.PostFlightHandler).Methods("GET")
	r.HandleFunc("/flights/{id}", routes.DeleteFlightHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

	Example()
}

func Example() {
	graph := src.NewRoutesGraph(6)
	graph.AddEdge(5, 2, src.NewFlight("Minnesota", "San Francisco", 100, 600, 100))
	graph.AddEdge(5, 0, src.NewFlight("Minnesota", "Las Vegas", 50, 300, 50))
	graph.AddEdge(4, 0, src.NewFlight("New York", "Las Vegas", 200, 1200, 200))
	graph.AddEdge(4, 1, src.NewFlight("New York", "Seattle", 300, 1500, 300))
	graph.AddEdge(3, 1, src.NewFlight("Los Ángeles", "Seattle", 400, 2400, 400))
	graph.AddEdge(2, 3, src.NewFlight("San Francisco", "Los Ángeles", 400, 2400, 400))

	topoSort := graph.TopoSort()
	fmt.Println(topoSort)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/felipeperezleal/routes_ms/db"
	"github.com/felipeperezleal/routes_ms/models"
	"github.com/felipeperezleal/routes_ms/routes"
	"github.com/felipeperezleal/routes_ms/src"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Routes{})

	message := "Buscando la mejor ruta, por favor espere"
	publishToRabbitMQ(message)
	fmt.Println("Mensaje enviado a RabbitMQ: ", message)

	startAlgorithm()
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

	http.ListenAndServe(":8080", r)
}

func startAlgorithm() {
	flightsData, err := src.FetchFlights()
	if err != nil {
		log.Fatal("Error al obtener datos de vuelos:", err)
	}

	var flights []models.Flight
	if err := json.Unmarshal(flightsData, &flights); err != nil {
		fmt.Println("Error al deserializar datos de vuelos:", err)
		return
	}
	fmt.Println("Datos de vuelos obtenidos:", flights)

	nodes := len(flights)
	route, dbRoute := src.NewRoute(nodes)

	for i, flight := range flights {
		route.AddEdge(i, i+1, &flight)
	}

	topoSort := route.TopoSort(dbRoute)
	fmt.Println("Orden topológico:", topoSort)
}

func publishToRabbitMQ(message string) {
	conn, err := amqp.Dial("amqp://guest:guest@tripster-mq:5672/")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	queueName := "tripster_queue"

	err = ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Mensaje enviado a RabbitMQ: %s", message)
}

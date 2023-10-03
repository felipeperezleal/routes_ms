package main

import (
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

	db.DB.AutoMigrate(models.Flight{})
	db.DB.AutoMigrate(models.Routes{})

	Example()
	message := "Buscando la mejor ruta, por favor espere"
	publishToRabbitMQ(message)
	fmt.Println("Mensaje enviado a RabbitMQ: ", message)

	startServer()
}

func startServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/flights", routes.GetFlightsHandler).Methods("GET")
	r.HandleFunc("/flights", routes.PostFlightHandler).Methods("POST")
	r.HandleFunc("/flights/{id}", routes.UpdateFlightHandler).Methods("PUT")
	r.HandleFunc("/flights/{id}", routes.GetFlightHandler).Methods("GET")
	r.HandleFunc("/flights/{id}", routes.DeleteFlightHandler).Methods("DELETE")

	r.HandleFunc("/routes", routes.GetRoutesHandler).Methods("GET")
	r.HandleFunc("/routes", routes.PostRouteHandler).Methods("POST")
	r.HandleFunc("/flights/{id}", routes.UpdateRouteHandler).Methods("PUT")
	r.HandleFunc("/routes/{id}", routes.GetRouteHandler).Methods("GET")
	r.HandleFunc("/routes/{id}", routes.DeleteRoutesHandler).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

func Example() {
	nodes := 6
	graph := src.NewRoute(nodes)

	graph.AddEdge(5, 2, src.NewFlight("Minnesota", "San Francisco", 100, 100))
	graph.AddEdge(5, 0, src.NewFlight("Minnesota", "Las Vegas", 50, 50))
	graph.AddEdge(4, 0, src.NewFlight("New York", "Las Vegas", 200, 200))
	graph.AddEdge(4, 1, src.NewFlight("New York", "Seattle", 300, 300))
	graph.AddEdge(3, 1, src.NewFlight("Los Ángeles", "Seattle", 400, 400))
	graph.AddEdge(2, 3, src.NewFlight("San Francisco", "Los Ángeles", 400, 400))
	topoSort := graph.TopoSort()
	fmt.Println(topoSort)
}

func publishToRabbitMQ(message string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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

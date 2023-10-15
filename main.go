package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

	startAlgorithm("A", "D")
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

func startAlgorithm(origin, destiny string) {

	// Uncomment for Fetching flights from flights_ms
	// flightData, err := src.FetchFlights()
	// if err != nil {
	// 	log.Printf("Error al obtener los datos de vuelo desde el API: %v", err)
	// 	return
	// }

	publishToRabbitMQ("Estamos calculando tu ruta...")

	flightData := []models.FlightData{
		{
			AirportOriginName:    "A",
			AirportDestinoName:   "B",
			FlightDepartureTime:  time.Date(2023, 9, 16, 5, 12, 0, 0, time.UTC),
			FlightArrivalTime:    time.Date(2023, 9, 16, 6, 47, 0, 0, time.UTC),
			FlightAirline:        "Latam",
			FlightSeatClass:      "Basic",
			FlightEscalas:        []string{},
			FlightAvailableSeats: 5,
			FlightTicketPrice:    511004.0,
		},
		{
			AirportOriginName:    "B",
			AirportDestinoName:   "C",
			FlightDepartureTime:  time.Date(2023, 9, 16, 5, 12, 0, 0, time.UTC),
			FlightArrivalTime:    time.Date(2023, 9, 16, 6, 47, 0, 0, time.UTC),
			FlightAirline:        "Latam",
			FlightSeatClass:      "Basic",
			FlightEscalas:        []string{},
			FlightAvailableSeats: 5,
			FlightTicketPrice:    345673.0,
		},
		{
			AirportOriginName:    "C",
			AirportDestinoName:   "D",
			FlightDepartureTime:  time.Date(2023, 9, 16, 5, 12, 0, 0, time.UTC),
			FlightArrivalTime:    time.Date(2023, 9, 16, 6, 47, 0, 0, time.UTC),
			FlightAirline:        "Latam",
			FlightSeatClass:      "Basic",
			FlightEscalas:        []string{},
			FlightAvailableSeats: 5,
			FlightTicketPrice:    654345.0,
		},
		{
			AirportOriginName:    "A",
			AirportDestinoName:   "D",
			FlightDepartureTime:  time.Date(2023, 9, 16, 5, 12, 0, 0, time.UTC),
			FlightArrivalTime:    time.Date(2023, 9, 16, 6, 47, 0, 0, time.UTC),
			FlightAirline:        "Latam",
			FlightSeatClass:      "Basic",
			FlightEscalas:        []string{},
			FlightAvailableSeats: 5,
			FlightTicketPrice:    234565.0,
		},
	}

	graph := src.NewGraph()

	for _, flight := range flightData {
		graph.AddRoute(flight.AirportOriginName, flight.AirportDestinoName)
	}
	graph.TopologicalSort(origin)

	route := models.Routes{
		Origin:   origin,
		Destiny:  destiny,
		NumNodes: len(graph.Sorted),
		Ordering: fmt.Sprintf("%v", graph.Sorted),
	}

	if err := db.DB.Create(&route).Error; err != nil {
		fmt.Printf("Error al crear el registro en la base de datos: %v\n", err)
	}

	publishToRabbitMQ("Terminamos de calcular tu ruta!")
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

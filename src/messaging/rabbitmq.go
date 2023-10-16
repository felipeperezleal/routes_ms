package messaging

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/felipeperezleal/routes_ms/db"
	"github.com/felipeperezleal/routes_ms/models"
	"github.com/felipeperezleal/routes_ms/src"
	"github.com/streadway/amqp"
)

func ExecuteAlgorithm(route *models.Routes) {

	// Uncomment for Fetching flights from flights_ms
	// flightData, err := src.FetchFlights()
	// if err != nil {
	// 	log.Printf("Error al obtener los datos de vuelo desde el API: %v", err)
	// 	return
	// }

	PublishToRabbitMQ("Estamos calculando tu ruta...")

	flightData := []models.FlightData{
		{
			AirportOriginName:    "El Dorado International Airport (BOG)",
			AirportDestinoName:   "Rafael Núñez International Airport (CTG)",
			FlightDepartureTime:  time.Date(2023, 9, 16, 8, 0, 0, 0, time.UTC),
			FlightArrivalTime:    time.Date(2023, 9, 16, 9, 15, 0, 0, time.UTC),
			FlightAirline:        "Avianca",
			FlightSeatClass:      "Economy",
			FlightEscalas:        []string{"No escalas"},
			FlightAvailableSeats: 120,
			FlightTicketPrice:    450000.0,
		},
		{
			AirportOriginName:    "José María Córdova International Airport (MDE)",
			AirportDestinoName:   "Alfonso Bonilla Aragón International Airport (CLO)",
			FlightDepartureTime:  time.Date(2023, 9, 16, 10, 30, 0, 0, time.UTC),
			FlightArrivalTime:    time.Date(2023, 9, 16, 12, 0, 0, 0, time.UTC),
			FlightAirline:        "LATAM",
			FlightSeatClass:      "Business",
			FlightEscalas:        []string{"Una escala en Bogotá (BOG)"},
			FlightAvailableSeats: 24,
			FlightTicketPrice:    890000.0,
		},
		{
			AirportOriginName:    "Gustavo Rojas Pinilla International Airport (ADZ)",
			AirportDestinoName:   "Simón Bolívar International Airport (SMR)",
			FlightDepartureTime:  time.Date(2023, 9, 16, 14, 15, 0, 0, time.UTC),
			FlightArrivalTime:    time.Date(2023, 9, 16, 15, 45, 0, 0, time.UTC),
			FlightAirline:        "Viva Air",
			FlightSeatClass:      "Economy",
			FlightEscalas:        []string{"No escalas"},
			FlightAvailableSeats: 150,
			FlightTicketPrice:    180000.0,
		},
		{
			AirportOriginName:    "Alfonso Bonilla Aragón International Airport (CLO)",
			AirportDestinoName:   "El Dorado International Airport (BOG)",
			FlightDepartureTime:  time.Date(2023, 9, 16, 16, 45, 0, 0, time.UTC),
			FlightArrivalTime:    time.Date(2023, 9, 16, 18, 15, 0, 0, time.UTC),
			FlightAirline:        "Avianca",
			FlightSeatClass:      "Economy",
			FlightEscalas:        []string{"No escalas"},
			FlightAvailableSeats: 90,
			FlightTicketPrice:    320000.0,
		},
		{
			AirportOriginName:    "El Dorado International Airport (BOG)",
			AirportDestinoName:   "José María Córdova International Airport (MDE)",
			FlightDepartureTime:  time.Date(2023, 9, 16, 20, 0, 0, 0, time.UTC),
			FlightArrivalTime:    time.Date(2023, 9, 16, 21, 30, 0, 0, time.UTC),
			FlightAirline:        "LATAM",
			FlightSeatClass:      "Economy",
			FlightEscalas:        []string{"No escalas"},
			FlightAvailableSeats: 60,
			FlightTicketPrice:    550000.0,
		},
	}

	graph := src.NewGraph(len(flightData))

	for _, flight := range flightData {
		graph.AddEdge(flight.AirportOriginName, flight.AirportDestinoName)
	}

	sorted := graph.TopologicalSort()
	fmt.Printf("Orden topológico: %v\n", sorted)

	routePath := src.FindRoute(sorted, route.Origin, route.Destiny)
	route.Ordering = fmt.Sprintf("[%s]", strings.Join(routePath, ", "))

	fmt.Printf("Ruta de %s a %s: %v\n", route.Origin, route.Destiny, route.Ordering)

	db.DB.Save(&route)

}

func PublishToRabbitMQ(message string) {
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

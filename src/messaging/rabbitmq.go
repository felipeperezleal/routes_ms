package messaging

import (
	"fmt"
	"log"
	"strings"

	"github.com/felipeperezleal/routes_ms/db"
	"github.com/felipeperezleal/routes_ms/models"
	"github.com/felipeperezleal/routes_ms/src"
	"github.com/streadway/amqp"
)

func ExecuteAlgorithm(route *models.Routes) {

	flightData, err := src.FetchFlights()
	if err != nil {
		log.Printf("Error al obtener los datos de vuelo desde el API: %v", err)
		return
	}

	PublishToRabbitMQ("Estamos calculando tu ruta...")

	graph := src.NewGraph(len(flightData))

	for _, flight := range flightData {
		graph.AddEdge(flight.AirportOrigin.AirportOriginName, flight.AirportDestination.AirportDestinoName)
	}

	sorted := graph.TopologicalSort()
	fmt.Printf("Orden topológico: %v\n", sorted)

	routePath := src.FindRoute(sorted, route.Origin, route.Destiny)
	route.Ordering = fmt.Sprintf("[%s]", strings.Join(routePath, ", "))
	route.NumNodes = len(routePath)

	fmt.Printf("Ruta de %s a %s: %v\n", route.Origin, route.Destiny, route.Ordering)

	db.DB.Save(&route)

}

func PublishToRabbitMQ(message string) {
	conn, err := amqp.Dial("amqp://guest:guest@host.docker.internal:5672/")

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

package models

import "time"

type FlightData struct {
	AirportOriginID      string    `json:"airport_origin_id"`
	AirportOriginName    string    `json:"airport_origin_name"`
	AirportDestinoID     string    `json:"airport_destino_id"`
	AirportDestinoName   string    `json:"airport_destino_name"`
	FlightDepartureTime  time.Time `json:"flight_departure_time"`
	FlightArrivalTime    time.Time `json:"flight_arrival_time"`
	FlightAirline        string    `json:"flight_airline"`
	FlightSeatClass      string    `json:"flight_seat_class"`
	FlightEscalas        []string  `json:"flight_escalas"`
	FlightAvailableSeats int       `json:"flight_available_seats"`
	FlightTicketPrice    float64   `json:"flight_ticket_price"`
}

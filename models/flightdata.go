package models	

type FlightAirportOrigin struct {
	AirportOriginName string `json:"airport_origin_name"`
}

type FlightAirportDestination struct {
	AirportDestinoName string `json:"airport_destino_name"`
}

type FlightData struct {
	AirportOrigin      FlightAirportOrigin      `json:"airport_origin"`
	AirportDestination FlightAirportDestination `json:"airport_destination"`
}

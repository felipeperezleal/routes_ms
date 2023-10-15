package src

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/felipeperezleal/routes_ms/models"
)

func FetchFlights() ([]models.FlightData, error) {
	apiGatewayURL := "http://localhost:5000/graphql"

	resp, err := http.Get(apiGatewayURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la solicitud al API Gateway: %s", resp.Status)
	}

	var responseData []models.FlightData
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, err
	}

	return responseData, nil
}

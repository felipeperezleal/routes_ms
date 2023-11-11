package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/felipeperezleal/routes_ms/models"
)

type graphqlResponse struct {
	Data struct {
		GetFlights []models.FlightData `json:"getFlights"`
	} `json:"data"`
}

func FetchFlights() ([]models.FlightData, error) {
	apiGatewayURL := "http://host.docker.internal:5000/graphql"

	query := `
		query {
			getFlights {
				airport_origin {
					airport_origin_name
				},
				airport_destination {
					airport_destino_name
				}
			}
		}`

	jsonQuery := map[string]string{
		"query": query,
	}
	jsonData, err := json.Marshal(jsonQuery)
	if err != nil {
		return nil, fmt.Errorf("error marshalling query: %w", err)
	}

	req, err := http.NewRequest("POST", apiGatewayURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var response graphqlResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return response.Data.GetFlights, nil
}

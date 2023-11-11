package src

import (
	"context"

	"github.com/felipeperezleal/routes_ms/models"
	"github.com/machinebox/graphql"
)

type graphqlResponse struct {
	Data struct {
		GetFlights []models.FlightData `json:"getFlights"`
	} `json:"data"`
}

func FetchFlights() ([]models.FlightData, error) {
	client := graphql.NewClient("http://host.docker.internal:5000/graphql")

	req := graphql.NewRequest(`
	query GetFLights{
		getFlights{
			airport_origin{
				airport_origin_name
			},
			airport_destination{
				airport_destino_name
			}
		}
	}`)

	ctx := context.Background()
	var respData graphqlResponse
	if err := client.Run(ctx, req, &respData); err != nil {
		return nil, err
	}

	return respData.Data.GetFlights, nil
}

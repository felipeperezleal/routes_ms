package src

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func FetchFlights() ([]byte, error) {
	apiGatewayURL := "http://localhost:5000/graphiql"

	query := `
	{
		flights {
            Origin
            Destination
            Duration
            Price
		}
	}
	`

	req, err := http.NewRequest("POST", apiGatewayURL, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la solicitud al API Gateway: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

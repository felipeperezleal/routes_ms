package src

import (
	"fmt"
	"io"
	"net/http"
)

func FetchFlights() ([]byte, error) {
	apiGatewayURL := "http://localhost:5000"

	resp, err := http.Get(apiGatewayURL)
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

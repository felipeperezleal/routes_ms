package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchFlights() ([]byte, error) {
	apiGatewayURL := "http://localhost:5000/graphql?query={flights{Origin,Destination,Duration,Price}}"

	resp, err := http.Get(apiGatewayURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la solicitud al API Gateway: %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

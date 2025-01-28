package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"manantaneja.com/go/crypto-masters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("Required 3 characters, received %v characters", len(currency))
	}
	currenyUpperCase := strings.ToUpper(currency)
	resp, err := http.Get(fmt.Sprintf(apiUrl, currenyUpperCase))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code returned: %v", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("unable to parse the response body")
	}

	var apiData ApiResponse
	err = json.Unmarshal(bodyBytes, &apiData)

	if err != nil {
		return nil, err
	}

	response := datatypes.Rate{Currency: currenyUpperCase, Price: apiData.Bid}

	return &response, nil
}

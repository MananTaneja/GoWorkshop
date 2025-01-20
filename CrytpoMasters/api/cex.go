package api

import (
	"fmt"
	"net/http"
	"strings"

	"manantaneja.com/go/crypto-masters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (datatypes.Rate, error) {
	currenyUpperCase := strings.ToUpper(currency)
	resp, err := http.Get(fmt.Sprintf(apiUrl, currenyUpperCase))

	if err != nil {
		return datatypes.Rate{}, err
	}

	// return datatypes.Rate{}
}

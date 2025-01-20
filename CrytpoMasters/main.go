package main

import (
	"fmt"

	"manantaneja.com/go/crypto-masters/api"
)

func main() {
	cyptoRate, err := api.GetRate("BTC")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Price of %s is: %.2f\n", cyptoRate.Currency, cyptoRate.Price)
}

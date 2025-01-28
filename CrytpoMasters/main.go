package main

import (
	"fmt"
	"sync"

	"manantaneja.com/go/crypto-masters/api"
)

func main() {
	var wg sync.WaitGroup // this is used to keep track of go routines

	currencies := []string{"BTC", "ETH", "BCH"}

	for _, currency := range currencies {
		wg.Add(1)

		// anonymous lambda function
		go func(currency string) {
			getCurrency(currency)
			wg.Done() // go support closure
		}(currency)
	}
	wg.Wait()
}

func getCurrency(currencyName string) {
	cyptoRate, err := api.GetRate(currencyName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Price of %s is: %.2f\n", cyptoRate.Currency, cyptoRate.Price)
}

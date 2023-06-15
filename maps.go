package main

import (
	"fmt"
)

func main() {

	stocks := map[string]float64{

		"IBM":  130,
		"AMZN": 2087.5,
		"GOOG": 2540.9,
	}

	fmt.Println("list of stocks", stocks)
	fmt.Println("number of stocks", len(stocks))
	fmt.Println("IBM stocks value", stocks["IBM"])
}

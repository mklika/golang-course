package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

func main() {
	b1 := Budget{"TravelBudget", 990.0, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1)
	fmt.Printf("%#v\n", b1)
	fmt.Println(b1.Expires.After(time.Date(2023, time.June, 23, 9, 32, 17, 23091344, time.Local)))

}

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
	fmt.Println(b1.TimeLeft())
	a := 20.0
	b1.Update(a)
	fmt.Printf("lets add", a)
	fmt.Println(b1.Balance)

}

func (b1 Budget) TimeLeft() time.Duration {
	return b1.Expires.Sub(time.Now().UTC())
}

func (b1 *Budget) Update(sum float64) {
	b1.Balance += sum
}

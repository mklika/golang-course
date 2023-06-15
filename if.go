package main

import (
	"fmt"
)

func main() {
	x := 3.0

	if x > 5 {
		fmt.Printf("x is %v and is bigger than 5", x)
	} else {
		fmt.Printf("x is %v and is smaller or equal to 5", x)
	}

	n := 1
	switch {
	case n < 2:
		fmt.Println("\nn is less than 2")
	case n >= 2:
		fmt.Println("\nn is two or more")

	}

}

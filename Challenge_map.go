package main

import (
	"fmt"
	"strings"
)

func main() {

	book := "The color of magic"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Println("book[0]")

	booklower := (strings.ToLower(book))
	fmt.Println(booklower)

	words := strings.Fields(booklower)

	fmt.Println(words)

	counts := map[string]int{}

	fmt.Println(counts)

}

package main

import (
	"fmt"
)

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Println("div is", div, "mod is", mod)

	values := []int{1, 2, 3, 4}
	fmt.Println(values)
	doubleAt(values, 2)
	fmt.Println(values)

	m := 3
	fmt.Println(m)
	doublePtr(&m)
	fmt.Println(m)
}

func add(a int, b int) int {

	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

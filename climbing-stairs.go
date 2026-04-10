package main

import "fmt"

func fibonacci(n int) int {

	a := 0
	b := 1

	for i := 0; i <= n; i++ {
		next := a + b
		a = b
		b = next
	}

	return a
}

func main() {
	n := 0

	fmt.Scan(&n)

	fmt.Println(fibonacci(n))
}
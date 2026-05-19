package main

import "fmt"

func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	prev, current := 0, 1
	for i := 2; i <= n; i++ {
		prev, current = current, prev+current
	}
	return current	
}

func main() {
	fmt.Println(fib(4))
}

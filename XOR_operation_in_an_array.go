package main

import "fmt"

func xorOperation(n int, start int) int {

	count, p := 0, 0

	for i := start; i <= 1000000; i += 2 {
		count ^= i
		p++
		if p == n {
			break
		}
	}
	return count
}

func main() {
	fmt.Println(xorOperation(4, 3))
}

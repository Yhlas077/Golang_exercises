package main

import "fmt"

func smallestEvenMultiple(n int) int {
	i, result := 2, 0
	for {
		if (i%2 == 0) && (i%n == 0) {
			result = i
			break
		}
		i++
		
	}
	return result
}

func main() {
	fmt.Println(smallestEvenMultiple(6))
}
package main

import "fmt"

func main() {
	fmt.Println(totalMoney(4))
}

func totalMoney(n int) int {

	sum, count, add, limit := 0, 0, 0, 7

	for i := 1; i <= n; i++ {

		if count == limit {
			add++
			count = add
			limit++

		}

		count++
		sum += count
	}
	return sum
}

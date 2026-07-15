package main

import "fmt"

func main() {
	fmt.Println(gcdOfOddEvenSums(4))
}

func gcdOfOddEvenSums(n int) int {
	sum_odd, sum_even, count, counter := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		counter++
		sum_odd += counter
		counter++
		sum_even += counter

		count++

		if sum_even % n == 0 && sum_odd % n == 0 {
			return count
		}
	}
	return count
}

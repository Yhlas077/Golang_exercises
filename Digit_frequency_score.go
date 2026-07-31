package main

import "fmt"

func main() {
	n := 224
	fmt.Println(digitFrequencyScore(n))
}

func digitFrequencyScore(n int) int {
	freq := [10]int{}

	temp := n
	for temp > 0 {
		digit := temp % 10
		freq[digit]++
		temp /= 10
	}

	score := 0
	for d := 0; d < 10; d++ {
		score += d * freq[d]
	}

	return score
}

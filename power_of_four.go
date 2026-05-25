package main

import "fmt"

func isPowerOfFour(n int) bool {

	var answer float64
	answer = float64(n)

	for answer > 4{
		answer /= 4
	}
	fmt.Println(answer)
	if answer == 4 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPowerOfFour(17))
}

package main

import "fmt"

func isPalindrome(x int) bool {
	b := false
	x1 := x
	x2 := 0
	for x1 > 0 {
		x2 *= 10
		x2 += x1 % 10
		x1 = x1 / 10
	}

	if x == x2 {
		b = true
	}
	return b
}

func main() {
	fmt.Println(isPalindrome(1211))
}

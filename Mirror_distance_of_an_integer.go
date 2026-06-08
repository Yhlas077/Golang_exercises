package main

import (
	"fmt"
	"math"
)

func mirrorDistance(n int) int {
	n_reverse := 0
	n1 := n
	for n > 0 {
		n_reverse += (n % 10)
		n_reverse *= 10
		n /= 10
	}
	return int(math.Abs(float64(n1 - (n_reverse / 10))))
}

func main() {
	fmt.Println(mirrorDistance(25))
}

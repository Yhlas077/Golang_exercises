package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(scoreOfString("zaz"))
}

func scoreOfString(s string) int {
	sum, result := 0.0, 0
	for i := 0; i < len(s)-1; i++ {
		sum = math.Abs(float64(s[i]) - float64(s[i+1]))
		result+=int(sum)
		}
	return result
}

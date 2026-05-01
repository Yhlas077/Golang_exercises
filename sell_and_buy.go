package main

import (
	"fmt"
	"math"
	"slices"
)

// import "slices"


func maxProfit(prices []int) int {
	min := math.MaxInt

	var b[]int

	for i:=0;i<=len(prices)-1; i++ {
		if min > prices[i] {
			min = prices[i]
		}
		b = append(b, prices[i] - min)
	}
	
	if len(b) == 0 {
		return 0
	}
	return slices.Max(b)
}

func main() {
	a := []int{2,4,1}
	fmt.Println(maxProfit(a))
}

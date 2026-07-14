package main

import "fmt"

func main() {
	a := []int {1,3,5,7}
	fmt.Println(alternatingSum(a))
}

func alternatingSum(nums []int) int {
	result, sign := 0, 1

	for i := 0; i < len(nums); i++ {
		sign *= -1
		result -= (nums[i] * sign)
	}
	return result
}

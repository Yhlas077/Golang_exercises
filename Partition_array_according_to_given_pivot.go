package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			result = append(result, nums[i])
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == pivot {
			result = append(result, nums[i])
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > pivot {
			result = append(result, nums[i])
		}
	}
	return result
}

func main() {
	a := []int{9, 12, 5, 10, 14, 3, 10}
	p := 10
	fmt.Println(pivotArray(a, p))
}

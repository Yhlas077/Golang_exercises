package main

import "fmt"

func singleNumber(nums []int) int {
	count := 0
	b := []int{}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		b = append(b, count)
		count = 0
	}

	for i := 0; i < len(nums); i++ {
		if b[i] == 1 {
			return nums[i]
		}
	}
	return 0
}

func main() {
	a := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(a))
}

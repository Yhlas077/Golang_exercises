package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var b []int
	for i := 0; i <= len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i]+nums[j] == target {
				b = append(b, i)
				b = append(b, j)
				return b

			}
		}
	}
	return nil
}

func main() {

	a := []int{3, 2, 4}

	fmt.Println(twoSum(a, 6))
}

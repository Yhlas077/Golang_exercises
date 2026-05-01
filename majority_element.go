package main

import (
	"fmt"
)

func majorityElement(nums []int) int {

	major := nums[0]
	value := 1

	for i := 1; i < len(nums); i++ {
		if value == 0 {
			value++
			major = nums[i]
		} else if major == nums[i] {
			value++
		} else {
			value--
		}
	}

	return major

}

func main() {
	nums := []int{3, 2, 3}
	a := majorityElement(nums)
	fmt.Println(a)
}

package main

import "fmt"

func findDegrees(matrix [][]int) []int {
	result := []int{}
	for i := 0; i < len(matrix); i++ {
		result = append(result, 0)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result[j] += matrix[i][j]
		}
	}

	return result
}

func main() {

	a := [][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}
	fmt.Println(findDegrees(a))

}

package main

import (
	"fmt"
)

// 	      1
//       1 1
//      1 2 1
//     1 3 3 1
// 	  1 4 6 4 1
// 	1 5 10 10 5 1
// 1 6 15 20 15 6 1
    var a [][]int

func generate(numRows int) [][]int {
	a = make([][]int, numRows)
    for i := 0; i < numRows; i++ {
		a[i] = make([]int, i+1)
		val := 1
		for j := 0; j <= i; j++ {
			a[i][j] = val
			val = val * (i - j) / (j + 1)
		}
	}
    return a
}

func main() {
	fmt.Println(generate(5))
}

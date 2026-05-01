package main

import (
	"fmt"
)

func romanToInt(s string) int {
	var result []int
	for i := 0; i <= len(s)-1; i++ {

		if s[i] == 'I' {
			result = append(result, 1)
		} else if s[i] == 'L' {
			result = append(result, 50)
		} else if s[i] == 'C' {
			result = append(result, 100)
		} else if s[i] == 'D' {
			result = append(result, 500)
		} else if s[i] == 'M' {
			result = append(result, 1000)
		} else if s[i] == 'X' {
			result = append(result, 10)
		} else if s[i] == 'V' {
			result = append(result, 5)
		}
	}
	get := 0
	for i := 0; i < len(result); i++ {
		if i < len(result)-1 && result[i] < result[i+1] {
			get -= result[i]
		} else {
			get += result[i]
		}
	}
	return get
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

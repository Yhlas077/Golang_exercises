package main

import "fmt"

func main() {
	s := "al&phaBET"
	fmt.Println(toLowerCase(s))
}

func toLowerCase(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			result += string(s[i] + 32)
		} else {
			result += string(s[i])
		}
	}
	return result
}

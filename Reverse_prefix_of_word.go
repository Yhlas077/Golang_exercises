package main

import "fmt"

func main() {
	s := "abcdefd"
	var ch byte = 'd'
	fmt.Println(reversePrefix(s, ch))
}

func reversePrefix(word string, ch byte) string {

	var char int
	result := ""

	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			char = i
			break
		}
	}

	for i := char; i >= 0; i-- {
		result += string(word[i])
	}

	for i := char+1; i < len(word); i++ {
		result+=string(word[i])
	}

	return result
}

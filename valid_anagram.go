package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	s1 := make(map[string]int)
	s2 := make(map[string]int)
	

	for _, v := range s {
		s1[string(v)] = 1 + s1[string(v)]
	}

	for _, v := range t {
		s2[string(v)] = 1 + s2[string(v)]
	}

	for i, _ := range s1 {
		if s1[i] != s2[i] {
			return false
		}

	}
	
	return true
}

func main() {
	var s, t string = "rat", "car"
	fmt.Println(isAnagram(s, t))
}

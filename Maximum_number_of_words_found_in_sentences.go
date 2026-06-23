package main

import "fmt"

func mostWordsFound(sentences []string) int {
	var s string
	var count = 1
	a := []int{}
	for i := 0; i < len(sentences); i++ {
		s = sentences[i]
		for j := 0; j < len(s); j++ {
			if s[j] == ' ' {
				count++
			}
		}
		a = append(a, count)
		count = 1

	}
	k := 0
	for i := 0; i < len(a); i++ {
		if k < a[i] {
			k = a[i]
		}
	}
	return k
}

func main() {
	s := []string{"please wait", "continue to fight", "continue to win"}
	fmt.Println(mostWordsFound(s))
}

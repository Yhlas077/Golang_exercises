package main

import "fmt"

func main() {
	s := "abc"
	fmt.Println(reverseDegree(s))
}

func reverseDegree(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "a" {
			result += (26 * (i + 1))
		} else if string(s[i]) == "b" {
			result += (25 * (i + 1))
		} else if string(s[i]) == "c" {
			result += (24 * (i + 1))
		} else if string(s[i]) == "d" {
			result += (23 * (i + 1))
		} else if string(s[i]) == "e" {
			result += (22 * (i + 1))
		} else if string(s[i]) == "f" {
			result += (21 * (i + 1))
		} else if string(s[i]) == "g" {
			result += (20 * (i + 1))
		} else if string(s[i]) == "h" {
			result += (19 * (i + 1))
		} else if string(s[i]) == "i" {
			result += (18 * (i + 1))
		} else if string(s[i]) == "j" {
			result += (17 * (i + 1))
		} else if string(s[i]) == "k" {
			result += (16 * (i + 1))
		} else if string(s[i]) == "l" {
			result += (15 * (i + 1))
		} else if string(s[i]) == "m" {
			result += (14 * (i + 1))
		} else if string(s[i]) == "n" {
			result += (13 * (i + 1))
		} else if string(s[i]) == "o" {
			result += (12 * (i + 1))
		} else if string(s[i]) == "p" {
			result += (11 * (i + 1))
		} else if string(s[i]) == "q" {
			result += (10 * (i + 1))
		} else if string(s[i]) == "r" {
			result += (9 * (i + 1))
		} else if string(s[i]) == "s" {
			result += (8 * (i + 1))
		} else if string(s[i]) == "t" {
			result += (7 * (i + 1))
		} else if string(s[i]) == "u" {
			result += (6 * (i + 1))
		} else if string(s[i]) == "v" {
			result += (5 * (i + 1))
		} else if string(s[i]) == "w" {
			result += (4 * (i + 1))
		} else if string(s[i]) == "x" {
			result += (3 * (i + 1))
		} else if string(s[i]) == "y" {
			result += (2 * (i + 1))
		} else if string(s[i]) == "z" {
			result += (1 * (i + 1))
		}
	}
	return result
}

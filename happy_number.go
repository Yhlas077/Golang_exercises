package main

import "fmt"

func isHappy(n int) bool {
	k, m, count := 0, 0, 0

	for {
		k = 0
		for n > 0 {
			m = n % 10
			k += (m * m)
			n = n / 10
		}
		n = k
		if k == 1 {
			return true
		}
		count++
		if count == 1000 {
			return false
		}
	}
}

func main() {
	fmt.Println(isHappy(2))
}

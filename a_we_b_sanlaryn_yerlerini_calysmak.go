package main

import "fmt"

var a,b int

func main() {
	fmt.Scan(&a, &b)
	a=b-a
	b=b-a
	a=a+b
	fmt.Printf("a: %d, b: %d", a, b)
}

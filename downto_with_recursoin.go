package main

import "fmt"

func Print_Down_to(n int) {
	if n==0 {return }
	fmt.Printf("%d\n",n)
	Print_Down_to(n-1)
}

func main() {
	n := 0
	fmt.Scan(&n)
	Print_Down_to(n)
}
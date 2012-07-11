package main

import "fmt"

func main() {
	fmt.Println(f(6))
}

func f(n int) int {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

package main

import "fmt"

func main() {
	i := 20
	fmt.Println("Fibonacci", i, fib(i))
}

func fib(n int) int {
	a, b, f := 0, 1, 0
	for i := 2; i <= n; i++ {
		f = a + b
		a = b
		b = f
	}
	return f
}

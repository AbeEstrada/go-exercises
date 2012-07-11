package main

import "fmt"

func fb(x int) {
	f, b := "Fizz", "Buzz"

	for i := 1; i <= x; i++ {
		if i%15 == 0 {
			fmt.Println(f + b)
		} else if i%5 == 0 {
			fmt.Println(b)
		} else if i%3 == 0 {
			fmt.Println(f)
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fb(100)
}

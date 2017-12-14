package main

import "fmt"

func f(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	return f(n-1) + f(n-2) + f(n-3)
}

func main() {
	fmt.Println(f(30))
}

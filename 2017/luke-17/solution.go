package main

import "fmt"

func main() {
	factor := 10
	for n := 1; ; n++ {
		A := 10*n + 6
		if factor < n {
			factor *= 10
		}
		B := 6*factor + n
		if 4*A == B {
			fmt.Println("Got it bro, 4 times", A, "equals", B)
			break
		}
	}
}

package main

import "fmt"

func main() {
	count := 0
	for b := 2; b < 136000/2; b++ {
		for a := 1; a < b; a++ {
			sum := b*(b+1)/2 - (a-1)*a/2
			if sum <= 130000 {
				count++
			}
		}
	}
	fmt.Println(count)
}

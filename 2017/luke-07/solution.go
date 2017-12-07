package main

import "fmt"

func main() {

	encrypt := make(map[rune]rune)
	for x := 'A'; x <= 'Z'; x++ {
		y := (x+x-'A'+1+x-'A')%26 + 'A'
		encrypt[x] = y
	}

	s := "OTUJNMQTYOQOVVNEOXQVAOXJEYA"

	for _, c := range s {
		fmt.Printf("%c", encrypt[encrypt[c]])
	}
}

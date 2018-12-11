package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	buf, _ := ioutil.ReadFile("input-crisscross.txt")
	isnum := true
	var X, Y int
	var num int
	for _, b := range buf {
		if isnum {
			num = int(b - '0')
		} else {
			switch b {
			case 'H':
				X += num
			case 'V':
				X -= num
			case 'F':
				Y += num
			case 'B':
				Y -= num
			}
		}
		isnum = !isnum
	}
	fmt.Printf("[%d,%d]\n", X, Y)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	N := 1337
	f := make([]int, N+1)
	for i, j := 1, 1; i <= N; i++ {
		if i%7 == 0 || strings.Contains(strconv.Itoa(i), "7") {
			f[i] = f[j]
			j++
		} else {
			f[i] = i
		}
	}
	fmt.Println(f[N])
}

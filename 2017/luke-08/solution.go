package main

import "fmt"

func Next(n int) int {
	sum := 0
	for {
		d := n % 10
		sum += d * d
		n = n / 10
		if n == 0 {
			break
		}
	}
	return sum
}

var cache map[int]bool

func SetCache(visited map[int]bool, valid bool) {
	for n, _ := range visited {
		cache[n] = valid
	}
}

func Valid(n int, visited map[int]bool) bool {
	if valid, ok := cache[n]; ok {
		return valid
	}

	visited[n] = true

	if n == 1 {
		SetCache(visited, true)
		return true
	}

	if n > 10000000 {
		SetCache(visited, false)
		return false
	}

	next := Next(n)
	if visited[next] {
		SetCache(visited, false)
		return false
	}

	return Valid(Next(n), visited)
}

func main() {
	cache = make(map[int]bool)
	sum := 0
	for i := 1; i <= 10000000; i++ {
		if i%100000 == 0 {
			fmt.Println("progress: ", 100*i/10000000, "%")
		}
		if Valid(i, map[int]bool{}) {
			sum += i
		}
	}
	fmt.Println(sum)
}

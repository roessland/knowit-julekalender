package main

import "fmt"
import "time"

var ps = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59,
	61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139,
	149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
	233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317,
	331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421,
	431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521}

func GenerateMonotonicallyDecreasingSequences(max int, size int, yield chan []int, shouldRecurse func([]int) bool) {
	var f func(int, int, []int, chan []int)

	f = func(max int, size int, prepend []int, yield chan []int) {
		if size == 0 {
			yield <- prepend
			return
		}
		if !shouldRecurse(prepend) {
			return
		}
		for first := 0; first <= max; first++ {
			pre := []int{}
			pre = append(pre, prepend...)
			pre = append(pre, first)
			f(first, size-1, pre, yield)
		}
	}

	f(max, size, []int{}, yield)
	close(yield)
}

func isBelowThreshold(indices []int) bool {
	numFactors := 9
	prod := 512
	for _, idx := range indices {
		prod *= ps[idx]
		numFactors++
	}
	for numFactors < 24 {
		prod *= 2
		numFactors++
	}
	return prod <= 4294967295
}

func main() {
	t0 := time.Now()
	seqsChan := make(chan []int)
	go GenerateMonotonicallyDecreasingSequences(len(ps)-1, 15, seqsChan, isBelowThreshold)
	count := 0
	for _ = range seqsChan {
		count++
	}
	fmt.Println("Found", count, "numbers")

	fmt.Println("Spent ", time.Since(t0))
}

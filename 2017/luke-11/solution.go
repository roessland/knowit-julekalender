package main

import "fmt"
import "github.com/roessland/gopkg/primegen"
import "github.com/roessland/gopkg/mathutil"

func Reverse(n int64) int64 {
	digits := mathutil.ToDigits(n, 10)
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	reversed := mathutil.FromDigits(digits, 10)
	return reversed
}

func main() {
	isPrime := primegen.Map(10001)
	count := 0
	for i := int64(1); i < int64(1000); i++ {
		if isPrime[i] && isPrime[Reverse(i)] && !mathutil.IsPalindrome(i) {
			fmt.Println(i)
			count++
		}
	}
	fmt.Println("=======")
	fmt.Println(count)
}

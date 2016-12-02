package main

// Fibonaccirekken er en tallrekke som genereres ved at man adderer de to
// foregående tallene i rekken. f.eks. om man starter med 1 og 2 blir de første
// 10 termene 1, 1, 2, 3, 5, 8, 13, 21, 34 og 55
//Finn summen av alle partall i denne rekken som er mindre enn 4.000.000.000

import "fmt"

func main() {
	sum := int64(0)
	a := int64(1)
	b := int64(1)
	for {
		if a >= 4000000000 {
			break
		}
		if a%2 == 0 {
			sum += a
		}

		a, b = b, a+b
	}
	fmt.Println(sum)
}

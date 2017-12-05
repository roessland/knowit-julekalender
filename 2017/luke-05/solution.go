package main

import "fmt"
import "os"

/*
Vi beskriver en tallrekke på følgende måte:

* Rekken inneholder positive heltall.

* Neste tall i rekken er alltid større enn eller likt det foregående tallet, det minker aldri.
* Verdien til tallet på n-te posisjon er antallet ganger n vil forekomme i rekken.
* Rekken starter med a(1) = 1.
* Sammen med de fem neste tallene i rekken får vi da: 1, 2, 2, 3, 3, 4 ...
* Finn summen av de 1000000 (1 million) første tallene i rekken.

1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21
1 2 2 3 3 4 4 4 5 5 5 4 4 4 4

*/

func main() {
	xs := make([]int, 1000000+1)
	xs[1] = 1
	curr := 2
	p := 2
	for i := 2; ; i++ {
		xs[p] = curr
		for j := 0; j < xs[curr]; j++ {
			xs[p] = curr
			p++
			if p >= len(xs) {
				sum := 0
				for _, x := range xs {
					sum += x
				}
				fmt.Println("sum: ", sum)
				os.Exit(0)
			}
		}
		curr++
	}
}

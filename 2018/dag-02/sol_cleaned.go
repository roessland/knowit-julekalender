package main

import "fmt"
import "log"
import "strings"
import "io/ioutil"
import "github.com/roessland/gopkg/mathutil"

func main() {

	buf, _ := ioutil.ReadFile("input.txt")
	counts := map[string]int{}

	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}

		var x0, y0, x1, y1 int
		_, err := fmt.Sscanf(line, "(%d,%d);(%d,%d)", &x0, &y0, &x1, &y1)
		if err != nil {
			log.Fatal(err)
		}

		dx := int64(x1 - x0)
		dy := int64(y1 - y0)

		gcd := mathutil.GCD(dx, dy)
		if dy < 0 {
			dx *= -1
			dy *= -1
		}
		slope := fmt.Sprintf("%d/%d", dx/gcd, dy/gcd)

		counts[slope]++
	}

	maxCount := 0
	maxCountSlope := ""
	for slope, cnt := range counts {
		if cnt > maxCount {
			maxCount = cnt
			maxCountSlope = slope
		}
	}

	fmt.Println(maxCount, "of the lines have a slope of", maxCountSlope)
}

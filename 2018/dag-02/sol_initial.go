package main

import "fmt"
import "strconv"
import "log"
import "bufio"
import "os"
import "strings"
import "io/ioutil"
import "encoding/csv"
import "github.com/roessland/gopkg/disjointset"
import "github.com/roessland/gopkg/mathutil"

type Vec struct {
	a, b int
}

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	counts := map[string]int{}
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, ";")
		parts[0] = parts[0][1 : len(parts[0])-1]
		parts[1] = parts[1][1 : len(parts[1])-1]
		fromStr := strings.Split(parts[0], ",")
		toStr := strings.Split(parts[1], ",")
		fromX := atoi(fromStr[0])
		fromY := atoi(fromStr[1])
		toX := atoi(toStr[0])
		toY := atoi(toStr[1])
		diffX := int64(toX - fromX)
		diffY := int64(toY - fromY)
		gcd := mathutil.GCD(diffX, diffY)
		slope := fmt.Sprintf("%d/%d", diffY/gcd, diffX/gcd)
		counts[slope]++
	}
	for _, cnt := range counts {
		if cnt > 300 {
			fmt.Println(cnt)
		}
	}
}

var _ = fmt.Println
var _ = strconv.Atoi
var _ = log.Fatal
var _ = bufio.NewScanner // (os.Stdin) -> scanner.Scan(), scanner.Text()
var _ = os.Stdin
var _ = strings.Split    // "str str" -> []string{"str", "str"}
var _ = ioutil.ReadFile  // ("input.txt") -> (buf, err)
var _ = csv.NewReader    // (os.Stdin)
var _ = disjointset.Make // (10) -> ds. ds.Union(a,b), ds.Connected(a,b), ds.Count

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

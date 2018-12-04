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

func main() {
	nums := []int{}
	buf, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		num := atoi(line)
		nums = append(nums, num)
	}

	fmt.Println(VekkSort([]int{5, 4, 3, 6, 7, 5, 2, 7, 5, 1, 1, 10}))
	fmt.Println(VekkSort(nums))
}

func VekkSort(nums []int) int {
	prevmax := -999
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < prevmax {
			continue
		}
		prevmax = nums[i]
		fmt.Println(nums[i])
		sum += nums[i]
	}
	return sum
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

package main

import "fmt"
import "strconv"
import "log"
import "math"
import "strings"
import "io/ioutil"

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

	fmt.Println(VekkSort(nums))
}

func VekkSort(nums []int) int {
	prevmax := math.MinInt32
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < prevmax {
			continue
		}
		prevmax = nums[i]
		sum += nums[i]
	}
	return sum
}

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

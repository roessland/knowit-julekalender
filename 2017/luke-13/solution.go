package main

import "fmt"
import "sort"
import "bufio"
import "strings"
import "strconv"
import "os"

func Complement(vals []bool) []bool {
	complement := make([]bool, len(vals))
	for i, val := range vals {
		complement[i] = !val
	}
	return complement
}

func ThueMorse(length int) []bool {
	s := []bool{false}
	for len(s) < length {
		s = append(s, Complement(s)...)
	}
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	values := []int{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		value, _ := strconv.Atoi(strings.Split(scanner.Text(), ", ")[1])
		values = append(values, value)
	}
	bobSum := 0
	bobsTurn := ThueMorse(len(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	for i := 0; i < len(values); i++ {
		if bobsTurn[i] {
			bobSum += values[i]
		}
	}
	fmt.Println(bobSum)
}

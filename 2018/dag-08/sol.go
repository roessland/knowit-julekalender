package main

import "fmt"
import "sort"
import "io/ioutil"
import "strings"
import "strconv"

type Doll struct {
	Size  int
	Color string
	m     int
}

func M(dolls []Doll, i int) int {
	if dolls[i].m != 0 {
		return dolls[i].m
	}
	maxM := 0
	for inner := 0; inner < i; inner++ {
		d := dolls[inner]
		if d.Size < dolls[i].Size && d.Color != dolls[i].Color {
			maxM = max(maxM, M(dolls, inner))
		}
	}
	dolls[i].m = 1 + maxM
	return 1 + maxM
}

func main() {
	buf, _ := ioutil.ReadFile("input-dolls.txt")
	dolls := []Doll{}
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		fields := strings.Split(line, ",")
		color := fields[0]
		size, _ := strconv.Atoi(fields[1])
		dolls = append(dolls, Doll{size, color, 0})
	}

	sort.Slice(dolls, func(i, j int) bool {
		return dolls[i].Size < dolls[j].Size
	})

	fmt.Println(M(dolls, len(dolls)-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Pop(s []int) ([]int, int) {
	if len(s) == 0 {
		log.Fatal("stack underflow during pop")
	}
	ret := s[len(s)-1]
	s = s[0 : len(s)-1]
	return s, ret
}

func Push(s []int, n int) []int {
	s = append(s, n)
	fmt.Println("  push", n)
	return s
}

func main() {
	buf, _ := ioutil.ReadFile("input.spp")
	s := []int{}
	var A, B, C int
	commands := map[rune]func(){
		':': func() {
			sum := 0
			for len(s) > 0 {
				s, A = Pop(s)
				sum += A
			}
			s = Push(s, sum)
		},
		'|': func() {
			s = Push(s, 3)
		},
		'\'': func() {
			s, A = Pop(s)
			s, B = Pop(s)
			s = Push(s, A+B)
		},
		'.': func() {
			s, A = Pop(s)
			s, B = Pop(s)
			s = Push(s, A-B)
			s = Push(s, B-A)
		},
		'_': func() {
			s, A = Pop(s)
			s, B = Pop(s)
			s = Push(s, A*B)
			s = Push(s, A)
		},
		'/': func() {
			s, _ = Pop(s)
		},
		'i': func() {
			s, A = Pop(s)
			s = Push(s, A)
			s = Push(s, A)
		},
		'\\': func() {
			s, A = Pop(s)
			s = Push(s, A+1)
		},
		'*': func() {
			s, A = Pop(s)
			s, B = Pop(s)
			s = Push(s, A/B)
		},
		']': func() {
			s, A = Pop(s)
			if A%2 == 0 {
				s = Push(s, 1)
			}
		},
		'[': func() {
			s, A = Pop(s)
			if A%2 == 1 {
				s = Push(s, A)
			}
		},
		'~': func() {
			s, A = Pop(s)
			s, B = Pop(s)
			s, C = Pop(s)
			s = Push(s, Max3(A, B, C))
		},
		' ': func() {
			s = Push(s, 31)
		},
	}
	for _, c := range string(buf) {
		r := rune(c)
		fmt.Println("Stack is", s)
		fmt.Printf("%c\n", r)
		if commands[r] != nil {
			commands[r]()
		} else {

		}
	}
	fmt.Println(s)
}

func Max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Max3(a, b, c int) int {
	return Max2(Max2(a, b), c)
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Person struct {
	Name    string
	Friends map[string]bool
	Hatees  map[string]bool
}

func NewPerson(name string) Person {
	return Person{name, make(map[string]bool), make(map[string]bool)}
}

type Graph struct {
	m map[string]Person
}

func (g Graph) CreatePerson(name string) Person {
	if _, ok := g.m[name]; !ok {
		g.m[name] = NewPerson(name)
	}
	return g.m[name]
}

func (g Graph) MakeFriends(name1, name2 string) {
	g.CreatePerson(name1)
	g.CreatePerson(name2)
	g.m[name1].Friends[name2] = true
	g.m[name2].Friends[name1] = true
}

func (g Graph) AddHatee(name1, name2 string) {
	g.CreatePerson(name1)
	g.CreatePerson(name2)
	g.m[name1].Hatees[name2] = true
}

func main() {
	// Create social graph
	g := Graph{make(map[string]Person)}
	buf, _ := ioutil.ReadFile("friendlist.txt")
	for _, line := range strings.Split(strings.TrimSpace(string(buf)), "\n") {
		fs := strings.Fields(line)
		f0, f1, f2 := fs[0], fs[1], fs[2]
		if f0 == "friends" {
			g.MakeFriends(f1, f2)
		} else if f1 == "hates" {
			g.AddHatee(f0, f2)
		}
	}

	maxCount := -1
	maxName := ""
	for name, person := range g.m {
		count := 0
		// Find friends she hates who don't hate her back
		for hatee, _ := range person.Hatees {
			if g.m[name].Friends[hatee] && !g.m[hatee].Hatees[name] {
				count += 1
			}
		}
		if count > maxCount {
			maxCount = count
			maxName = name
		}
	}
	fmt.Println("It's", maxName)
	fmt.Println("Friends:", len(g.m[maxName].Friends))
	fmt.Println("Hatees:", len(g.m[maxName].Hatees))
	fmt.Println("Chameleon relationships:", maxCount)
}

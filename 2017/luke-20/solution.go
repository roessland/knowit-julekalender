package main

import "fmt"
import "log"
import "encoding/csv"
import "os"

const (
	undetermined int = iota
	neutral
	friend
	enemy
)

type Person struct {
	Type    int
	Friends []string
	Enemies []string
}

func ReadInput() map[string]*Person {
	g := map[string]*Person{}
	reader := csv.NewReader(os.Stdin)
	reader.Comma = ' '
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("error readall", err)
	}
	for _, record := range records {
		nA, nB := record[1], record[2]
		if g[nA] == nil {
			g[nA] = &Person{undetermined, make([]string, 0), make([]string, 0)}
		}
		if g[nB] == nil {
			g[nB] = &Person{undetermined, make([]string, 0), make([]string, 0)}
		}
		switch record[0] {
		case "fiendskap":
			g[nA].Enemies = append(g[nA].Enemies, nB)
			g[nB].Enemies = append(g[nB].Enemies, nA)
		case "vennskap":
			g[nA].Friends = append(g[nA].Friends, nB)
			g[nB].Friends = append(g[nB].Friends, nA)
		default:
			log.Fatal("Unknown:" + record[0])
		}
	}
	return g
}

func Dfs(g map[string]*Person, start string, isFriend bool) {
	if g[start].Type != undetermined {
		return
	}
	if isFriend {
		g[start].Type = friend
	} else {
		g[start].Type = enemy
	}
	for _, enemyName := range g[start].Enemies {
		Dfs(g, enemyName, !isFriend)
	}
	for _, friendName := range g[start].Friends {
		Dfs(g, friendName, isFriend)
	}
}

func main() {
	g := ReadInput()
	Dfs(g, "Asgeir", true)
	Dfs(g, "Beate", false)
	numFriend := 0
	numEnemy := 0
	numNeutral := 0
	for _, p := range g {
		switch p.Type {
		case friend:
			numFriend++
		case enemy:
			numEnemy++
		case undetermined:
			numNeutral++
		}
	}
	fmt.Println(numFriend, numEnemy, numNeutral)
}

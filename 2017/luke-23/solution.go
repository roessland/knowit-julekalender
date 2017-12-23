package main

import "fmt"
import "os"
import "log"
import "io/ioutil"

type Game [10]rune

func (g Game) GetWinner() rune {
	if g[1] != 0 && g[1] == g[2] && g[2] == g[3] {
		return g[1]
	}
	if g[4] != 0 && g[4] == g[5] && g[5] == g[6] {
		return g[4]
	}
	if g[7] != 0 && g[7] == g[8] && g[8] == g[9] {
		return g[7]
	}
	if g[1] != 0 && g[1] == g[4] && g[4] == g[7] {
		return g[1]
	}
	if g[2] != 0 && g[2] == g[5] && g[5] == g[8] {
		return g[2]
	}
	if g[3] != 0 && g[3] == g[6] && g[6] == g[9] {
		return g[3]
	}
	if g[1] != 0 && g[1] == g[5] && g[5] == g[9] {
		return g[1]
	}
	if g[3] != 0 && g[3] == g[5] && g[5] == g[7] {
		return g[3]
	}
	return 0
}

func main() {
	wins := [2]int{0, 0}
	currX, currO := 0, 1
	locs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	game := Game{}
	numDraws := 0
	totalMoves := 0
	numMoves := 0
	curr, next := 'X', 'O'
	for _, loc := range locs {
		pos := loc - '0'
		if pos < 1 || pos > 9 {
			fmt.Println("WTF", pos)
			break
		}
		game[pos] = curr
		numMoves++
		totalMoves++
		w := game.GetWinner()
		if w == 'X' {
			fmt.Println("X wins, player", currX)
			wins[currX]++
			numDraws = 0
			game = Game{}
			numMoves = 0
			currX, currO = currO, currX
			curr, next = 'X', 'O'
		} else if w == 'O' {
			fmt.Println("O wins, player", currO)
			wins[currO]++
			numDraws = 0
			game = Game{}
			numMoves = 0
			curr, next = 'X', 'O'
		} else if numMoves == 9 {
			fmt.Println("draw")
			game = Game{}
			numMoves = 0
			numDraws++
			curr, next = 'X', 'O'
			if numDraws == 3 {
				fmt.Println("three draws in a row")
				currX, currO = currO, currX
				numDraws = 0
			}
		} else {
			curr, next = next, curr
		}
		if totalMoves == 10000 {
			break
		}
	}
	fmt.Println(wins)
}

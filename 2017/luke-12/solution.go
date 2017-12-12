package main

import "fmt"

type Pos struct {
	X, Y int
}

func (p Pos) IsLegal() bool {
	return 0 <= p.X && p.X < 10 && 0 <= p.Y && p.Y < 10
}

func (p Pos) Add(v Pos) Pos {
	return Pos{p.X + v.X, p.Y + v.Y}
}

type Color int

const black Color = Color(1)
const white Color = Color(0)

func (c Color) Invert() Color {
	if c == black {
		return white
	}
	return black
}

func main() {
	board := map[Pos]Color{}
	knightMoves := []Pos{
		{-2, -1},
		{-2, 1},
		{-1, -2},
		{-1, 2},
		{1, -2},
		{1, 2},
		{2, -1},
		{2, 1},
	}
	currPos := Pos{0, 0}
	for i := 0; i < 200; i++ {
		currColor := board[currPos]
		lastLegalPos := Pos{99, 99}
		for _, dp := range knightMoves {
			nextPos := currPos.Add(dp)
			if !nextPos.IsLegal() {
				continue
			}
			lastLegalPos = nextPos
			if board[lastLegalPos] == currColor {
				break
			}
		}
		board[currPos] = board[currPos].Invert()
		currPos = lastLegalPos
	}
	numBlack := 0
	for _, c := range board {
		if c == black {
			numBlack++
		}
	}
	fmt.Println(numBlack)

}

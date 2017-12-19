package main

import tl "github.com/JoelOtter/termloop"
import "encoding/csv"
import "strconv"
import "os"
import "time"
import "strings"
import "log"

type Pos struct {
	X, Y int
}

type Command struct {
	Times int
	Dir   Pos
}

func InputChannel() <-chan (Command) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	cmdChan := make(chan (Command))
	go func() {
		for _, record := range records {
			times, err := strconv.Atoi(record[0])
			if err != nil {
				log.Fatal(err)
			}
			var dir Pos
			switch strings.Trim(record[1], " ") {
			case "east":
				dir = Pos{1, 0}
			case "west":
				dir = Pos{-1, 0}
			case "south":
				dir = Pos{0, 1}
			case "north":
				dir = Pos{0, -1}
			default:
				log.Fatal("Unknown direction")
			}
			cmdChan <- Command{times, dir}
		}
	}()
	return cmdChan
}

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorWhite,
		Fg: tl.ColorBlue,
		Ch: '.',
	})
	go func() {
		width, height := game.Screen().Size()
		curr := Pos{width/4 - 5, height/2 + 10} // chosen by fair dice roll
		level.AddEntity(tl.NewRectangle(2*curr.X, curr.Y, 2, 1, tl.ColorBlue))
		for cmd := range InputChannel() {
			for i := 0; i < cmd.Times; i++ {
				curr.X += cmd.Dir.X
				curr.Y += cmd.Dir.Y
				level.AddEntity(tl.NewRectangle(2*curr.X, curr.Y, 2, 1, tl.ColorBlue))
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()
	game.Screen().SetLevel(level)
	game.Start()
}

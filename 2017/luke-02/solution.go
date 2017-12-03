package main

import "fmt"
import "math/bits"

type Pos struct {
	x, y int
}

type Cell struct {
	visited bool
	isWall  bool
}

// returned for out-of-bounds indices
var DummyCell Cell = Cell{visited: true, isWall: true}

type Map struct {
	width  int
	height int
	cells  []Cell
}

func (m *Map) GetCell(x, y int) *Cell {
	if x < 0 || m.width <= x {
		return &DummyCell
	}
	if y < 0 || m.height <= y {
		return &DummyCell
	}
	idx := y*m.width + x

	return &m.cells[idx]
}

func (m *Map) SetCell(x, y int, cell Cell) {
	m.cells[y*m.width+x] = cell
}

func (m *Map) Print() {
	fmt.Println()
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			c := m.GetCell(x, y)
			if c.isWall {
				fmt.Printf("#")
			} else {
				if c.visited {
					fmt.Printf(".")
				} else {
					fmt.Printf("_")
				}
			}
		}
		fmt.Println()
	}
}

func NewMap(width, height int) Map {
	m := Map{width, height, nil}
	m.cells = make([]Cell, width*height)
	for y := 1; y <= height; y++ {
		for x := 1; x <= width; x++ {
			c := Cell{}
			if bits.OnesCount(uint(x*x*x+12*x*y+5*x*y*y))%2 == 1 {
				c.isWall = true
			}
			m.SetCell(x-1, y-1, c)
		}
	}
	return m
}

func (m *Map) DFS(x, y int) {
	c := m.GetCell(x, y)
	if !c.visited && !c.isWall {
		c.visited = true
		m.DFS(x-1, y)
		m.DFS(x+1, y)
		m.DFS(x, y-1)
		m.DFS(x, y+1)
	}
}

func (m *Map) CountUnvisited() map[string]int {
	ret := make(map[string]int)
	for _, c := range m.cells {
		if !c.isWall && !c.visited {
			ret["unvisitedempty"]++
		}
		if c.isWall {
			ret["wall"]++
		}
		if !c.isWall {
			ret["empty"]++
		}
		if c.visited {
			ret["visited"]++
		}
	}
	return ret
}

func main() {
	N := 1000
	m := NewMap(N, N)
	m.DFS(0, 0)
	fmt.Printf("Unvisited cells count: %v\n", m.CountUnvisited()["unvisitedempty"])
}

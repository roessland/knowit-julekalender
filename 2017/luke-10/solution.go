package main

import "fmt"

type Node struct {
	Number int
	Next   *Node
}

func main() {
	N := 1500
	nodes := make([]Node, N)
	for i := 0; i < N; i++ {
		nodes[i].Number = i + 1
		nodes[i].Next = &nodes[(i+1)%N]
	}
	server := &nodes[0]
	for server.Next != server {
		fmt.Println(server.Number, "serves", server.Next.Number, "who plummets to the ground")
		server.Next = server.Next.Next
		server = server.Next
	}
	fmt.Println("Last man standing: ", server.Number)
}

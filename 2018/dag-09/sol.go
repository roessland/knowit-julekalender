package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
)

type Node struct {
	Ch   string `json:"ch"`
	Hash string `json:"hash"`
}

func main() {
	f, _ := os.Open("input-hashchain.json")
	nodes := []Node{}
	decoder := json.NewDecoder(f)
	decoder.Decode(&nodes)
	hash := Md5("julekalender")
	found := true
	for found == true {
		found = false
		for _, node := range nodes {
			if Md5(hash+node.Ch) == node.Hash {
				found = true
				hash = node.Hash
				fmt.Printf("%s", node.Ch)
				break
			}
		}
	}
}

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

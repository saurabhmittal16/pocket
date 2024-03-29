package main

import (
	"fmt"

	"github.com/saurabhmittal16/pocket/core"
)

func main() {
	node1 := core.CreateNode(1)
	node1.Put("1", "A")
	node1.Put("2", "B")

	fmt.Println(node1.Get("1"))
	fmt.Println(node1.Get("2"))
	fmt.Println(node1.Get("3"))

	node2 := core.CreateNode(2)
	fmt.Println(node2.Get("1"))
}

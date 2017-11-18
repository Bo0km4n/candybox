package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Next *Node
	Data int
}

func main() {
	nodes := make([]Node, 0, 10)
	var result Node
	head := mockNode()
	debugPrint(head)
	RecurrentFindNode(head, 4, 0, nodes, result)
}

func debugPrint(node *Node) {
	fmt.Println(node.Next, node.Data)
	if node.Next != nil {
		debugPrint(node.Next)
	}
}

func RecurrentFindNode(node *Node, k int, i int, buffer []Node, result Node) {
	buffer = append(buffer, *node)

	if node.Next == nil {
		if k > i {
			return
		}
		fmt.Println(i, k)
		fmt.Println(buffer[i-k])
		result = buffer[i-k]
		return
	}

	RecurrentFindNode(node.Next, k, i+1, buffer, result)
}

func mockNode() *Node {
	head := Node{
		Next: nil,
		Data: rand.Intn(100),
	}

	genNode(&head)
	buf := head.Next
	for i := 0; i < 10; i++ {
		genNode(buf)
		buf = buf.Next
	}

	return &head
}

func genNode(prev *Node) {
	node := Node{
		Next: nil,
		Data: rand.Intn(100),
	}
	prev.Next = &node
}

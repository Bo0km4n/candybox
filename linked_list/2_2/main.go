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
	head := mockNode()
	debugPrint(head)
	result := RecurrentFindNode(head, 4)
	fmt.Println(result)
}

func debugPrint(node *Node) {
	fmt.Println(node.Next, node.Data)
	if node.Next != nil {
		debugPrint(node.Next)
	}
}

func RecurrentFindNode(head *Node, k int) *Node {
	if k < 0 {
		return nil
	}

	p1 := head
	p2 := head

	for i := 0; i < k-1; i++ {
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}
	if p2 == nil {
		return nil
	}

	for {
		if p2.Next == nil {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
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

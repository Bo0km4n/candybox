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

	fmt.Println("=============================")

	result := Partition(head, 50)
	debugPrint(result)
}

func debugPrint(node *Node) {
	for {
		fmt.Println(*node)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}

func Partition(node *Node, x int) *Node {
	var beforeStart *Node
	var beforeEnd *Node
	var afterStart *Node
	var afterEnd *Node

	for {
		if node == nil {
			break
		}
		next := node.Next
		node.Next = nil

		if node.Data < x {
			if beforeStart == nil {
				beforeStart = node
				beforeEnd = beforeStart
			} else {
				beforeEnd.Next = node
				beforeEnd = node
			}
		} else {
			if afterStart == nil {
				afterStart = node
				afterEnd = afterStart
			} else {
				afterEnd.Next = node
				afterEnd = node
			}
		}

		node = next
	}

	if beforeStart == nil {
		return afterStart
	}

	beforeEnd.Next = afterStart
	return beforeStart
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

package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Prev *Node
	Next *Node
	Data int
}

func main() {
	fmt.Println("Hello, playground")
	head := mockNode()
	debugPrint(head)
	DescNode(head)
	debugPrint(head)
}

func debugPrint(node *Node) {
	fmt.Println(node.Next, node.Data)
	if node.Next != nil {
		debugPrint(node.Next)
	}
}

func mockNode() *Node {
	head := Node{
		Prev: nil,
		Next: nil,
		Data: rand.Intn(6),
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
		Prev: prev,
		Next: nil,
		Data: rand.Intn(6),
	}
	prev.Next = &node
}

func compareData(n1, n2 *Node) bool {
	return n1.Data == n2.Data
}

func deleteNode(node *Node) {
	if node == nil {
		return
	}
	node.Prev.Next = node.Next
	if node.Next == nil {
		fmt.Println("Deleted %v", node)
	} else {
		node.Next.Prev = node.Prev
	}
	fmt.Println("Deleted %v", node)
}

func Compare(node, nextNode *Node) {
	if compareData(node, nextNode) {
		deleteNode(nextNode)
		if nextNode.Next == nil {
			return
		}
		Compare(node, nextNode.Next)
	} else {
		if nextNode.Next == nil {
			return
		}
		Compare(node, nextNode.Next)
	}
	return
}

func DescNode(head *Node) {
	for {
		if head == nil || head.Next == nil {
			return
		}
		Compare(head, head.Next)
		DescNode(head.Next)
		return
	}
	return
}

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
	DeleteMidNode(head)
	debugPrint(head)
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

func DeleteMidNode(head *Node) {
	p1 := head
	p2 := head
	var prev *Node
	var next *Node
	var delete *Node
	k := 0

	if p2.Next == nil || p1.Next == nil {
		return
	}
	for {
		if p2.Next == nil {
			break
		}
		p2 = p2.Next
		k = k + 1
	}
	k = k / 2
	for i := 0; i <= k+1; i++ {
		switch {
		case i == k-1:
			prev = p1
		case i == k:
			delete = p1
		case i == k+1:
			next = p1
		}

		p1 = p1.Next
	}

	prev.Next = next
	delete.Next = nil
	return
}

// 中央のノードのみにアクセス可能
// リストが2, 1個の場合削除不可能
func AnswerDelete(n *Node) {
	if n == nil || n.Next == nil {
		return
	}
	next := n.Next
	n.Data = next.Data
	n.Next = next.Next
	return
}

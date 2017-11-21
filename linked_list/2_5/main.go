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
	head1 := mockNode()
	head2 := mockNode()

	fmt.Println("=====head_1=====")
	debugPrint(head1)
	fmt.Println("=====head_2=====")
	debugPrint(head2)

	result := ListSum(head1, head2)
	fmt.Println("=====result=====")
	debugPrint(result)
}

func ListSum(head1 *Node, head2 *Node) *Node {
	var result *Node
	resultHead := result
	node1 := head1
	node2 := head2
	var carry int
	var value int
	for {
		if node1 == nil && node2 == nil && carry == 0 {
			break
		}

		// 加算処理
		value = carry
		if node1 != nil {
			value += node1.Data
		}
		if node2 != nil {
			value += node2.Data
		}
		if value >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		buf := Node{
			Next: nil,
			Data: value % 10,
		}

		// リストに結果を格納していく
		if result == nil {
			result = &buf
			resultHead = result
		} else {
			result.Next = &buf
			result = result.Next
		}

		// インクリメント処理
		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}

	}
	return resultHead
}

func debugPrint(node *Node) {
	for {
		fmt.Print(node.Data)
		if node.Next == nil {
			fmt.Print("\n")
			break
		}
		fmt.Print(" -> ")
		node = node.Next
	}
}

func mockNode() *Node {
	head := Node{
		Next: nil,
		Data: rand.Intn(9),
	}

	genNode(&head)
	buf := head.Next
	for i := 0; i < rand.Intn(5); i++ {
		genNode(buf)
		buf = buf.Next
	}

	return &head

}

func genNode(prev *Node) {
	node := Node{
		Next: nil,
		Data: rand.Intn(9),
	}
	prev.Next = &node
}

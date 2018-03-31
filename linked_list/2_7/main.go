package main

import (
	"fmt"
)

type Node struct {
	Next *Node
	Name string
}

func main() {
	a := Node{Name: "A"}
	b := Node{Name: "B"}
	c := Node{Name: "C"}
	d := Node{Name: "C"}
	e := Node{Name: "B"}
	f := Node{Name: "A"}

	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = &f
	f.Next = nil

	if answerCheckPalindrome(&a) {
		fmt.Println("回文です")
		return
	}
	fmt.Println("回文ではない")
}

func answerCheckPalindrome(head *Node) bool {
	fast := head
	slow := head

	stack := []string{}

	// リストの前半部分をスタックに積む
	// ファストランナーが終点に到達した時, slowは丁度半分
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		stack = append(stack, slow.Name)
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 要素数が奇数の場合には真ん中の要素を飛ばす
	if fast != nil {
		slow = slow.Next
	}

	// 前半部分と後半部分を比較
	for {
		if slow == nil {
			break
		}
		popVal := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if popVal != slow.Name {
			return false
		}
		slow = slow.Next
	}
	return true
}

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
	d := Node{Name: "D"}
	e := Node{Name: "E"}

	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = nil

	r := getCycleHeadNode(&a)
	fmt.Printf("%v\n", r)
	answer := answerGetCycleHeadNode(&a)
	fmt.Printf("%v\n", answer)
}

func getCycleHeadNode(head *Node) *Node {
	cache := map[*Node]string{}
	for {
		if head.Next == nil {
			break
		}
		if _, ok := cache[head]; ok {
			return head
		}
		cache[head] = head.Name
		head = head.Next
	}
	return nil
}

func answerGetCycleHeadNode(head *Node) *Node {
	slow := head
	fast := head

	// ループ内の衝突点探し
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// fastがnilならループ無し
	if fast == nil || fast.Next == nil {
		return nil
	}

	// slowをリストの先頭に移動する
	// fastは衝突点に残す
	// k個同じ速さで進むとループ開始地点で出会う
	slow = head
	for {
		if slow == fast {
			break
		}
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

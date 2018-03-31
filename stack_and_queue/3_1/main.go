package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	stack := genStack()
	stack.insertToStack(1, "h")
	stack.insertToStack(2, "h")
	stack.insertToStack(3, "h")
	stack.insertToStack(1, "p")
	stack.insertToStack(3, "p")
	stack.insertToStack(3, "p")
	stack.insertToStack(3, "p")
	pp.Println(stack.popFromStack(3))
	pp.Println(stack.popFromStack(3))
	pp.Println(stack.popFromStack(3))
	pp.Println(stack.popFromStack(3))
}

type stack struct {
	stack []string
	head  []int
	tail  []int
	init  []int
}

func genStack() *stack {
	s := make([]string, 10)
	h := []int{-1, -1, -1}
	t := []int{0, 0, 0}
	sLen := len(s)
	hLen := len(h)
	unit := sLen / hLen
	for i := 0; i < hLen; i++ {
		if i == hLen-1 {
			t[i] = sLen - 1
			h[i] = sLen - 2 - unit
		}
		t[i] = (i+1)*unit - 1
		h[i] = i*unit - 1
	}
	return &stack{stack: s, head: h, tail: t, init: h}
}

func (s *stack) insertToStack(stackNum int, data string) {
	n := stackNum - 1
	if s.head[n] == s.tail[n] {
		fmt.Printf("stack [%d] is max\n", stackNum)
		return
	}
	s.head[n]++
	s.stack[s.head[n]] = data
}

func (s *stack) popFromStack(stackNum int) string {
	n := stackNum - 1
	if s.init[n] < s.head[n]-1 {
		return ""
	}
	s.head[n]--
	return s.stack[s.head[n]+1]
}

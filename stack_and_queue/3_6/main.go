package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {
	s := newStack()
	s.push(5)
	s.push(3)
	s.push(4)
	s.push(1)
	s.push(2)
	fmt.Println(s)
	pp.Println(s.sort())
}

type stack struct {
	data []int
}

func (s *stack) sort() *stack {
	buf := newStack()

	for {
		if s.isEmpty() {
			break
		}
		tmp := s.pop()
		for {
			if !buf.isEmpty() && buf.peek() < tmp {
				s.push(buf.pop())
			} else {
				break
			}
		}
		buf.push(tmp)
	}
	return buf
}

func newStack() *stack {
	return &stack{data: make([]int, 0, 0)}
}

func (s *stack) push(v int) {
	s.data = append(s.data, v)
}

func (s *stack) pop() int {
	if len(s.data) == 0 {
		return 0
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *stack) peek() int {
	if len(s.data) == 0 {
		return 0
	}
	v := s.data[len(s.data)-1]
	return v
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

package main

import "github.com/k0kubun/pp"

type stack struct {
	stack []int
	min   []int
}

type node struct {
	next *node
	data int
}

func main() {
	s := stack{}
	s.init()
	s.push(5)
	pp.Println(s)
	s.push(1)
	pp.Println(s)
	s.push(3)
	pp.Println(s)
	s.push(2)
	pp.Println(s)

}

func (s *stack) init() {
	s.stack = make([]int, 0, 0)
	s.min = make([]int, 0, 0)
}

func (s *stack) push(v int) {
	if v < s.getMin() {
		s.pushMin(v)
	}
	s.stack = append(s.stack, v)
}

func (s *stack) pop() int {
	v := s.stack[len(s.stack)-1]
	if v == s.getMin() {
		s.min = s.min[:len(s.min)-1]
	}
	s.stack = s.stack[:len(s.stack)-1]
	return v
}

func (s *stack) getMin() int {
	if len(s.min) == 0 {
		return 999999999
	}
	m := s.min[len(s.min)-1]
	return m
}

func (s *stack) pushMin(v int) {
	s.min = append(s.min, v)
}

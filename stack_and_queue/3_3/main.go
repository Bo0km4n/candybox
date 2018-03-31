package main

import "github.com/k0kubun/pp"

var stackSize = 5

type setOfStacks struct {
	stacks [][]int
	cursor int
}

func main() {
	s := setOfStacks{}
	s.init()
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(3)
	s.push(3)
	s.push(3)
	s.push(3)
	s.push(3)
	pp.Println(s)
	pp.Println(s.pop())
	pp.Println(s.pop())
	pp.Println(s.pop())
	pp.Println(s.pop())
	pp.Println(s)
}

func (s *setOfStacks) init() {
	s.stacks = make([][]int, 1)
	s.cursor = 0
}

func (s *setOfStacks) push(v int) {
	stack := s.stacks[s.cursor]
	if len(stack) >= stackSize {
		s.addStack()
		s.cursor++
	} else {
		s.stacks[s.cursor] = append(s.stacks[s.cursor], v)
	}
}

func (s *setOfStacks) addStack() {
	s.stacks = append(s.stacks, make([]int, 0, 0))
}

func (s *setOfStacks) pop() int {
	stack := s.stacks[s.cursor]
	v := stack[len(stack)-1]
	s.stacks[s.cursor] = stack[:len(stack)-1]

	// stackの長さチェック
	// 0だったらstacksから削除
	if len(s.stacks[s.cursor]) == 0 {
		s.stacks = s.stacks[:s.cursor]
		s.cursor--
	}
	return v
}

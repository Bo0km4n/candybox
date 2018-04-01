package main

import "fmt"

func main() {
	q := newQueue()
	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)
	fmt.Println(q.pop())
	fmt.Println(q.pop())
}

type myQueue struct {
	stackNewest []int
	stackOldest []int
}

func newQueue() *myQueue {
	sn := make([]int, 0, 0)
	so := make([]int, 0, 0)
	return &myQueue{stackNewest: sn, stackOldest: so}
}

func (m *myQueue) push(v int) {
	m.stackNewest = append(m.stackNewest, v)
}

func (m *myQueue) shiftStack() {
	for {
		if len(m.stackNewest) == 0 {
			break
		}
		v := m.stackNewest[len(m.stackNewest)-1]
		m.stackOldest = append(m.stackOldest, v)
		m.stackNewest = m.stackNewest[:len(m.stackNewest)-1]
	}
	return
}

func (m *myQueue) pop() int {
	m.shiftStack()
	v := m.stackOldest[len(m.stackOldest)-1]
	m.stackOldest = m.stackOldest[:len(m.stackOldest)-1]
	return v
}

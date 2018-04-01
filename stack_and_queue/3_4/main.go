package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type tower struct {
	disks []int
	index int
}

func main() {
	towers := newTowers()
	towers[0].push(10)
	towers[0].push(9)
	towers[0].push(8)
	towers[0].push(7)
	towers[0].push(6)
	towers[0].moveDisks(2, &towers[2], &towers[1])
	pp.Println(towers)
}

func newTowers() []tower {
	t1 := tower{disks: make([]int, 0, 0), index: 1}
	t2 := tower{disks: make([]int, 0, 0), index: 2}
	t3 := tower{disks: make([]int, 0, 0), index: 3}

	return []tower{t1, t2, t3}
}

func (t *tower) moveDisks(n int, dest *tower, buf *tower) {
	if n > 0 {
		t.moveDisks(n-1, buf, dest)
		t.moveTopTo(dest)
		buf.moveDisks(n-1, dest, t)
	}
}

func (t *tower) moveTopTo(dest *tower) {
	if len(t.disks) == 0 {
		return
	}
	top := t.pop()
	dest.push(top)
	fmt.Printf("Move disk %d from %d to %d\n", top, t.index, dest.index)
}

func (t *tower) push(v int) {
	if len(t.disks) != 0 && t.peek() <= v {
		fmt.Printf("Error placing disk %d\n", v)
	} else {
		t.disks = append(t.disks, v)
	}
}

func (t *tower) pop() int {
	if len(t.disks) == 0 {
		return 0
	}
	v := t.disks[len(t.disks)-1]
	t.disks = t.disks[:len(t.disks)-1]
	return v
}

func (t *tower) peek() int {
	v := t.disks[len(t.disks)-1]
	return v
}

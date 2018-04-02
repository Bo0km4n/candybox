package main

import (
	"math"
)

func main() {

}

type node struct {
	right *node
	left  *node
	data  int
}

func (n *node) checkHeight() int {
	var leftHeight int
	var rightHeight int
	if n == nil {
		return 0
	}

	if n.left == nil {
		leftHeight = 0
	} else {
		leftHeight = n.left.checkHeight()
		if leftHeight == -1 {
			return -1
		}
	}

	if n.right == nil {
		rightHeight = 0
	} else {
		rightHeight = n.right.checkHeight()
		if rightHeight == -1 {
			return -1
		}
	}

	heightDiff := leftHeight - rightHeight
	if math.Abs(float64(heightDiff)) > 1 {
		return -1
	}
	return int(math.Max(float64(leftHeight), float64(rightHeight)) + 1)
}

func (n *node) isBalanced() bool {
	if n.checkHeight() == -1 {
		return false
	}
	return true
}

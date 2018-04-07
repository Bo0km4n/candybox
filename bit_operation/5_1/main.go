package main

import (
	"fmt"
	"strconv"
)

func main() {
	updateBits(10000000000, 10011, 2, 6)
}

func updateBits(n, m, i, j int) uint32 {
	n2, _ := strconv.ParseInt(fmt.Sprintf("%d", n), 2, 0)
	m2, _ := strconv.ParseInt(fmt.Sprintf("%d", m), 2, 0)

	allOnes := ^uint32(0)
	left := allOnes << (uint32(j) + 1)
	right := (uint32(1) << uint32(i)) - 1
	mask := uint32(left | right)

	n_cleared := uint32(n2) & mask
	m_shifted := uint32(m2) << uint32(i)

	return n_cleared | m_shifted
}

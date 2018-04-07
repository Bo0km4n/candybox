package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(printBinary(0.3))
}

func printBinary(n float64) string {
	binary := make([]string, 0, 32)
	binary = append(binary, "0.")
	for {
		if n <= 0 {
			break
		}
		if len(binary) >= 32 {
			return "ERROR"
		}
		r := n * 2
		if r >= 1 {
			binary = append(binary, "1")
			n = r - 1
		} else {
			binary = append(binary, "0")
			n = r
		}
	}
	return strings.Join(binary, "")
}

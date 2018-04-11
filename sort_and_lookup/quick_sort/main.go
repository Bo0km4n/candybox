package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	fmt.Printf(">> ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	t := sc.Text()

	array := strings.Split(t, " ")
	numArray := make([]int, 0)
	for i := range array {
		index, err := strconv.Atoi(array[i])
		if err != nil {
			continue
		}
		numArray = append(numArray, index)
	}

	pp.Println(numArray)

	quickSort(&numArray, 0, len(numArray)-1)

	pp.Println(numArray)
}

func quickSort(a *[]int, left int, right int) {
	var pivot int

	if right > left {
		pivot = partition(a, left, right)
		quickSort(a, left, pivot-1)
		quickSort(a, pivot+1, right)
	}
}

func partition(a *[]int, left int, right int) int {
	tmp_left := left
	pivotItem := (*a)[left]
	for left < right {
		for (*a)[left] <= pivotItem {
			left++
		}
		for (*a)[right] > pivotItem {
			right--
		}
		if left < right {
			(*a)[right], (*a)[left] = (*a)[left], (*a)[right]
		}
	}
	(*a)[tmp_left] = (*a)[right]
	(*a)[right] = pivotItem
	return right
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{10, 20, 45, 55, 67, 85, 48, 95, 0, 12, 54}
	index := sort.SearchInts(a, 67)
	fmt.Println(index)
}

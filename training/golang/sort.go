package main

import (
	"fmt"
	"sort"
)

func testSortSlice() {
	var data = []int{4, 8, 1, 3, 0}
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	fmt.Println(data)
}

func main() {
	testSortSlice()
}

package main

import (
	"fmt"
)

// 1. index on array ptr is ok
// 2.index on slice ptr is not ok
func testArrayPtr() {
	var d = [2]int{1, 2}
	var p *[2]int = &d
	fmt.Println(p[0])
	var s = d[:]
	fmt.Println(s[0])
}

func main() {
	testArrayPtr()
}

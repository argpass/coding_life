package main

import (
	"fmt"
)

func testOffset() {
	var i int32 = 2
	// 0
	fmt.Println(i >> 2)
	// 4
	fmt.Println(i << 1)
	// -2
	var j uint32 = 0
	fmt.Println(^j >> 1)
}
func main() {
	testOffset()
}

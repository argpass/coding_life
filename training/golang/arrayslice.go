package main

import (
	"fmt"
)

// append 会生成新的slice(因为原来的slice空间可能不足，所生成新的并拷贝值过去)
// 所以才有append操作的函数都应该返回新slice
func testSliceArgs() {
	var raw = []int{1, 2, 3}
	func(x []int) {
		// i can change #0 to 88 means that x is reference to raw
		x[0] = 88
		// now x is point to the `append` result (new slice)
		x = append(x, 99)
		// to change on new slice does not change raw slice
		x[1] = 776
	}(raw)
	fmt.Println(raw)
}

func testBasic() {
	var s []int
	// s is nil
	fmt.Println(s == nil)

	// space is assigned
	s = []int{}
	fmt.Println(s != nil)

	// s is assigned space when used
	s = append(s, 99)
	fmt.Println(s)
}

func main() {

	testSliceArgs()
}

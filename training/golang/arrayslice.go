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

// 1.slice[length:] => [], cap 0
// 2.slice[i:i] => [], len 0, cap is cap of [i:]
// 3.slice[i:i-1] is wrong
func testSlice() {
	var one = []int{1}
	s := one[1:]
	fmt.Println("len-1, slice[1:]", s)
	fmt.Println("len-1, slice[1:].cap is 0? ", cap(s) == 0)
	two := []int{1, 2}
	s2 := two[1:1]
	fmt.Println("slice[0:0] is ", s2)
	fmt.Println("slice[0:0] cap is ", cap(s2))
	fmt.Println("slice[0:0] len is ", len(s2))
	var d = make([]int, 4)
	fmt.Println("len of d:", len(d))
}

func foo(p []int) {
	p = append(p, 99)
	p[1] = 77
}

// arry is value type, so would never be nil
func tArray() {
	var brr [2]int
	fmt.Println("brr is nil ?", brr == nil)
}

func main() {
	tArray()
	d := make([]int, 2, 20)
	d[0] = 1
	d[1] = 2
	foo(d)
	fmt.Println(d)
	testSlice()
}

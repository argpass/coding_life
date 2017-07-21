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
// 4.数组指针可以下标索引不用解除引用，切片指针则不行
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
// unsupported -1 index
func tArray() {
	//var arr = [2]int{88, 99}
	//fmt.Println("last is ", arr[-1])
}

// [:0] is the same with the origin [] in lower-layer memory,
// thus cap() is the same
// [1:] cap减少1, cap只看所指向内存区域左端
func testClear() {
	var s = make([]int, 100)
	s[0] = 1997
	fmt.Println("s is :", s)
	var b = s[:0]
	b = append(b, 1998)
	fmt.Println("b is:", b)
	fmt.Println("b's cap is:", cap(b))
	fmt.Println("s is :", s)
}


func main() {
	var arr1 [6]int
	var slice1 []int = arr1[1:2]
	fmt.Println(len(slice1), cap(slice1))
}

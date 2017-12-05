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

// 迭代中改变slice会生效
func testModifyOnItering() {
	var data []int = []int{1, 2, 3, 4, 5, 6}
	for i, d := range data {
		fmt.Printf("%d->%d\n", i, d)
		fmt.Println(data)
		data[len(data)-i-1] = 99
	}
}

func testCap() {
	var arr1 [6]int
	var slice1 []int = arr1[1:2]
	fmt.Println(len(slice1), cap(slice1))
}

func testAppend() {
	arr := make([]int, 3)
	fmt.Println(arr)
	arr = append(arr, 99)
	fmt.Println(arr)
}

func testSliceCap() {
	s := make([]int, 0, 10)
	s1 := s[:0:5]
	s2 := s[5:5:10]
	fmt.Printf("len %d, cap %d\n", len(s1), cap(s1))
	fmt.Printf("len %d, cap %d\n", len(s2), cap(s2))
	fmt.Printf("s:%v\n", s)
	s2 = append(s2, 99)
	s1 = append(s1, 11)
	fmt.Printf("len %d, cap %d\n", len(s1), cap(s1))
	fmt.Printf("len %d, cap %d\n", len(s2), cap(s2))
	s3 := s[0:10]
	fmt.Printf("s3:%+v\n", s3)
	// 这很神奇，copy是使用len计算的，如果任意一个len为0 则copy 0
	d := []int{1, 2}
	// 但你可以动态切片来修改, 就像这样
	n := copy(s1[len(s1):cap(s1)], d)
	s1 = s1[0 : len(s1)+n]
	fmt.Printf("copied n:%d\n", n)
	fmt.Printf("len %d, cap %d, data:%v\n", len(s1), cap(s1), s1)
}

type MyData struct {
	Age int
}

func (m *MyData) ChangeAge(v int) {
	m.Age = v
}

func testModifyElem() {
	s := make([]MyData, 5, 10)
	s[0].ChangeAge(99)
	s[0].Age++
	s[0].Age += 88
	v := &(s[0])
	v.ChangeAge(55)
	fmt.Printf("s:%+v\n", s)
}


func main() {
	testModifyElem()
}

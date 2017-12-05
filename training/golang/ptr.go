package main

import (
	"fmt"
)

func testStructPtrField()  {
	type MyData struct {
		Value *int
	}
	data := MyData{}
	fmt.Printf("value is nil ? `%v`\n", data.Value == nil)
}

func testScopeAddr()  {
	for _, v := range []int{0,1,2} {
		// 每一次都是一个新的a
		// 每一次都是同一个v
		a := 99
		fmt.Printf("addr of a is :%p, addr of v is :%p\n", &a, &v)
	}
}

// 不能直接取函数调用结果的地址(编译器检查)
//func testTakeAddr()  {
//	fn := func() (v int){
//		return 99
//	}
//	fmt.Printf("return.addr:%p\n", &fn())
//}

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
}

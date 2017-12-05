package main

import "fmt"

func testMyTypeType() {
	type statusType int

	const (
		a statusType = 1
		b
	)
	fmt.Printf("a type:%T, b type:%T\n", a, b)
	// 右值是常量，自动做了类型转换
	var c statusType = b
	fmt.Println(c)
}

func main() {
	testMyTypeType()
}

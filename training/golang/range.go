package main

import (
	"fmt"
)

type MyData struct {
	Age int
}

func (m *MyData) ChangeAge(v int) {
	m.Age = v
}

func testChangeOnItering() {
	ds := make([]MyData, 5)
	// 实际上d只是一个拷贝，所以不能修改ds
	for _, d := range ds {
		d.ChangeAge(99)
	}
	fmt.Printf("after itering, ds:%+v\n", ds)
}

func main() {
	testChangeOnItering()
}

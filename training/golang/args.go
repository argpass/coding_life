package main

import (
	"fmt"
)

// nil 切片做变参是ok的
func testNilSliceArgs()  {
	fn := func(args ...string){
		fmt.Printf("got args:%+v\n", args)
	}
	var args []string
	fn(args...)
}

type MyData struct {
}

func(m MyData) Foo(data MyData) *MyData {
	return &data
}

func testArgRef()  {
	d := MyData{}
	fmt.Printf("%p\n", d.Foo(d))
	fmt.Printf("%p\n", d.Foo(MyData{}))
}

func main()  {
	testArgRef()
}

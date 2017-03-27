package main

import "fmt"

// nil 切片做变参是ok的
func testNilSliceArgs()  {
	fn := func(args ...string){
		fmt.Printf("got args:%+v\n", args)
	}
	var args []string
	fn(args...)
}

func main()  {
	testNilSliceArgs()
}

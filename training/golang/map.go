package main

import (
	"fmt"
)

// 1.函数不能作为map key，函数不能hash
// 2.指针可以作key
func testCanMapKey() {
	var i = 99
	var a = map[*int]int{&i: i}
	fmt.Println(a)
	var c = map[interface{}]int{nil: i}
	fmt.Println(c)
}

func testMapBasic() {
	var d = map[string]string{}
	d["name"] = "akun"
	name := d["name"]
	fmt.Println("got name ", name)
}

func main() {
	testMapBasic()
	testCanMapKey()
}

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

// v := map[key], key不存在的时候v为该类型的默认值而不是nil，也不panic
func testMapBasic() {
	var d = map[string]string{}
	d["name"] = "akun"
	name := d["name"]
	fmt.Println("got name ", name)
	age := d["age"]
	fmt.Println("[not exist string key]== ?", age == "")
	var id = map[string]int{}
	fmt.Println("[not exist int key]== ?", id["age"] == 0)
}

func testOp() {
	var d = map[string]int{}
	d["a"] = 9
	d["e"] = 8
	d["b"] = 2
	d["yu"] = 2
	d["kk"] = 2
	d["jj"] = 2
	d["ef"] = 2
	fmt.Printf("len of map:%d\n", len(d))
	for k, v := range d {
		fmt.Println(k, v)
	}
	fmt.Println("-----")
	for k, v := range d {
		fmt.Println(k, v)
	}
	fmt.Println("-----")
	for k, v := range d {
		fmt.Println(k, v)
	}
	fmt.Println("-----")
	delete(d, "abc")
}

func main() {
	testOp()
}

package main

import "fmt"

type Meta struct {
	Data map[string]interface{}
	Name string
	age  int
}

func testStructEqual() {
	var m interface{} = Meta{Data: map[string]interface{}{"name": "akun", "age": 99}}
	var n interface{} = Meta{Data: map[string]interface{}{"name": "akun", "age": 99}}
	//q := m
	fmt.Printf("m==n?:%v\n", m == n)
}

func testMyTypeEqual() {
	type MyInt int
	var a MyInt = 99
	var b int = 99
	var c MyInt = 99
	result := func(a interface{}, b interface{}) bool {
		return a == b
	}(a, b)
	fmt.Println(result)

	result = func(a interface{}, b interface{}) bool {
		return a == b
	}(a, c)
	fmt.Println(result)
}

func main() {
	testMyTypeEqual()
}

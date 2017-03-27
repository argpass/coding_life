package main

import "fmt"

type Meta struct {
	Data map[string]interface{}
	Name string
	age  int
}

func testStructEqual() {
	var m interface{} = Meta{Data:map[string]interface{}{"name": "akun", "age": 99}}
	var n interface{} = Meta{Data:map[string]interface{}{"name": "akun", "age": 99}}
	//q := m
	fmt.Printf("m==n?:%v\n", m == n)
}

func main()  {
	testStructEqual()
}

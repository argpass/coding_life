package main

import (
	"fmt"
)

type Parent struct{}

type Child struct {
	*Parent
}

type VChild struct {
	Parent
}

type dict map[string]interface{}

type MyDict struct {
	dict
}

func (m MyDict) Add(key string, value interface{}) {
	m.dict[key] = value
}

func (m MyDict) Show() {
	fmt.Printf("\nMydict:%v\n", m.dict)
}

func testMyDict() {
	var d = MyDict{dict{}}
	d.Add("name", "akun")
	d.Add("age", 99)
	d.Show()
}

func foo(pp *Parent) {
	fmt.Println(pp)
}

func main() {
	testMyDict()
}

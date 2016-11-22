package main

import (
	"fmt"
)

type Model interface {
	isModel() bool
}

type ModelBase struct {
}

func (p *ModelBase) isModel() bool {
	return true
}

type Person struct {
	ModelBase
	Name string
}

type Writer interface {
	Write() string
}

func testNil() {
	var w Writer = nil
	fmt.Println("w.Write():", w.Write())
}

func main() {
	testNil()
	var m Model = &Person{}
	fmt.Println(m.isModel())
	var a Model
	// yes
	fmt.Println("a is nil?", a == nil)
	var b Model = nil
	// yes
	fmt.Println("b is nil?", b == nil)
	var c *Person = nil
	// 相当于被装箱， d is Model<nil> but not nil
	var d Model = (Model)(c)
	fmt.Println("(Model)(*Person)  d is nil? ", d == nil)
}

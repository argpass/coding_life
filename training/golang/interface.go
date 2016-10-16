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

func main() {
	var m Model = &Person{}
	fmt.Println(m.isModel())
	var a Model
	fmt.Println("a is nil?", a == nil)
	var b Model = nil
	fmt.Println("b is nil?", b == nil)
	var c *Person = nil
	// 相当于被装箱， d is Model<nil> but not nil
	var d Model = (Model)(c)
	fmt.Println("(Model)(*Person)  d is nil? ", d == nil)
}

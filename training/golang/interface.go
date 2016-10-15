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
}

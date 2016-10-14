package main

import (
	"fmt"
	"reflect"
)

type Model interface {
	Fields()
}

type Base struct{}

func (p *Base) Fields() {
	fmt.Println("Fields")
	fmt.Println(p)
	// nil pointer will panic
	i := reflect.ValueOf(*p).NumField()
	fmt.Println(i)
}

type MyModel struct {
	*Base
	ModelName string
}

func main() {
	var b Model = &MyModel{}

	// 0, only got fields num of *Base
	b.Fields()

}

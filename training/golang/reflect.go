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
}

type MyModel struct {
	Base
	ModelName string
}

func show(m Model) {
	tp := reflect.Indirect(reflect.ValueOf(m))
	i := tp.NumField()
	name := tp.String()
	fmt.Println(i, name)
}

type Field interface {
	Incr()
}

type FInt struct {
	value int
}

func (p *FInt) Incr() {}

type Student struct {
	Age *FInt
}

func testNilInterface() {
	var s = &Student{}
	tp := reflect.Indirect(reflect.ValueOf(s))
	fi := tp.Field(0)
	// fi IsNil but fi.Interface()!= nil
	fmt.Printf("\nfi is nil?:%v", fi.IsNil())
	if m, ok := fi.Interface().(Field); ok {
		fmt.Printf("\nm == nil ?:%v", m == nil)
		fmt.Printf("\nfi.Interface():%v, ==nil?:%v", fi.Interface(), fi.Interface() == nil)
		s.Age = &FInt{99}
		// m is interface wrapper and value copy, so m is still <nil> object(not nil)
		fmt.Printf("\nm now is:%v", m)
		// fi.Interface() can get the newest value of the field
		fmt.Printf("\nfi.Interface() is :%v", fi.Interface())
	}
}

func main() {
	var b Model = &MyModel{}

	// 0, only got fields num of *Base
	b.Fields()
	show(b)
	testNilInterface()
}

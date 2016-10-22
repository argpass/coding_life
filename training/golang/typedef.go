// 1.函数能作为receiver,可以实现接口
// 2.interface 不能作为receiver,即使重新定义了一个type 自interface
// 3. type NewType OldType 定义新Type后，不能再用OldType的方法，还可以用旧Type的域;
//    如果要使用旧Type的方法，应该组合该OldType

package main

import (
	"fmt"
	"reflect"
)

type Runner interface {
	Run()
}

type Fn func()

func (p Fn) Run() {
	fmt.Println("func as receiver and p is ", p)
}

func testFnAsInterface() {
	var f interface{} = Fn(func() {})
	f.(Runner).Run()
}

type Base struct {
	Data int
}

func (p *Base) SayP() {}
func (p Base) SayV()  {}

type MyValue Base

type MyPointer *Base

// 不能再用原来type的方法，但是fields还可以用
func testNewTypeMethods() {
	var mv = MyValue{99}
	fmt.Println("MyValue.Data:", mv.Data)
	fmt.Println("MyValue methods cnt:", reflect.ValueOf(mv).NumMethod())
	var mp = &MyValue{}
	fmt.Println("MyValue methods cnt:", reflect.ValueOf(mp).NumMethod())
}

func main() {
	testNewTypeMethods()
}

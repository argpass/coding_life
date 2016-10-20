// 1.函数能作为receiver
// 2.interface 不能作为receiver,即使重新定义了一个type 自interface

package main

import (
	"fmt"
)

type Runner interface {
	Run()
}

type Fn func()

func (p Fn) Run() {
	fmt.Println("func as receiver and p is ", p)
}

func main() {
	var f interface{} = Fn(func() {})
	f.(Runner).Run()
}

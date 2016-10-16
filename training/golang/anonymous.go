// 使用匿名结构
// 匿名结构动态转换
package main

import (
	"fmt"
)

type Base struct {
}

func (p Base) Say() {
	fmt.Println("hello base")
}

type Model struct{}

func (p Model) Run() {
	fmt.Println("model run")
}

func main() {
	var d = map[string]struct{}{}
	d["key"] = Model{}
	if c, ok := d["key"]; ok {
		// 结构定义一样所以可以进行类型cast
		var s = (Base)(c)
		s.Say()
	}
}

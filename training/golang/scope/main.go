package main

import "github.com/demos/training/scope/inner"

// 通过public方法暴露私有类型对象是可以的
// 非导出的字段 在外部匿名赋值是不行的
func main()  {
	v := inner.MyType{9, 10}
	v.Say()
}

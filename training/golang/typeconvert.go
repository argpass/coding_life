// 类型标记：
// 1.声明、设置变量时如左边未指定类型则编译器根据右值推断类型，如左边指定则检查左右类型；最后都会给变量设置类型标识
// 2.强制类型转换如果成功则改变类型标识
// 3.类型断言其实就是判定是否是某个类型标识
// 4.不能把[]interface{} 转 []int 尽管该interface{}是可以转为int,因为类型标识不兼容
package main

import (
	"fmt"
)

type ParamMap map[string]interface{}

var E_NONE = fmt.Errorf("no data")
var E_INVALID_INT = fmt.Errorf("Invalid %s", "int")

func (m ParamMap) GetInt(key string) (value int, err error) {
	var ok bool
	var v interface{}
	value, err = 0, nil
	v, ok = m[key]
	if !ok {
		err = E_NONE
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		err = E_INVALID_INT
	}
	return value, err
}

func convertCustomType() {
	var err error
	var v int
	var d = map[string]interface{}{"int": 99}
	// convert to ParamMap type
	var c = ParamMap(d)
	_, err = c.GetInt("age")
	fmt.Println(err == E_NONE)

	v, err = c.GetInt("int")
	fmt.Println(v, err)
}

func convertBasicType() {
	var valueFloat float64 = 99
	var valueInt int = 33

	// convert float to int
	valueInt = int(valueFloat)
	fmt.Println(valueInt)
}

// 不能把[]interface{} 转 []int 尽管该interface{}是可以转为int
func convertSlice() {
	type VSlice []interface{}
	type ISlice []int
	//var v = VSlice{2, 3}
	//var i = ISlice(v)
	//fmt.Println("convert v to ISlice:", i)
}

func convertFunc() {
	type MapFunc func(x interface{})

	// 在定义变量的时候相当于编译器需要给其打上类型标记，标记设为`func(x interface{})`
	// 所以断言失败
	var f interface{} = func(x interface{}) {}
	if _, ok := f.(MapFunc); ok {
		fmt.Println("assert f func(x interface{}){} is MapFunc")
	} else {
		fmt.Println("assert f func(x interface{}){} isn't MapFunc")
	}

	// 在定义变量的时候相当于编译器需要给其打上类型标记，此处强制类型转换把标记设为`MapFun`
	// 所以断言成功
	var f2 interface{} = MapFunc(func(x interface{}) {})
	if _, ok := f2.(MapFunc); ok {
		fmt.Println("assert f2 MapFunc(func(x interface{}){}) is MapFunc")
	} else {
		fmt.Println("assert f2 MapFunc(func(x interface{}){}) isn't MapFunc")
	}
}

type Sayer interface {
	Say()
}
type Person struct {
	Name string
}

func (p *Person) Say() {}

type PT Person

func main() {
	//convertSlice()
	convertFunc()
	var p = &PT{"akun"}
	p.Say()
	fmt.Println(p)
	//p.Say()
	//(*Person)(p).Say()
}

package main

import (
	"errors"
	"fmt"
)

func mockValue() (value int, err error) {
	return 99, nil
}

// 测试":"赋值
func testColonAssignment() {
	var value int
	fmt.Printf("value1 %p\n", &value)
	defer func() {
		fmt.Printf("value in defer %p\n", &value)
	}()
	// 还是旧的value；同时创建新的err
	value, err := mockValue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("value2 %p\n", &value)

	// 新作用于创建了新value,类型随意
	for value, _ := range []int64{1} {
		fmt.Printf("value3 %p\n", &value)
	}

	// 还是旧的value
	fmt.Printf("value4 %p\n", &value)
}

func testNamedReturnInner() (err error) {
	a := 99
	err = errors.New("err1")
	if a == 99 {
		// 当前直接作用域下无err，故err 是新建立的变量
		r, err := func() (int, error) {
			return 88, nil
		}()
		if err != nil {
			return err
		}
		a = r
	}
	// err1
	fmt.Printf("err1:%v\n", err)
	// 不会新建err，使用当前作用域下的err(即返回值err)
	r, err := func() (int, error) {
		return 88, nil
	}()
	if err != nil {
		return err
	}
	a = r
	return
}

// 默认err返回值为nil
func testNamedReturn2() (err error) {
	return
}

func testNamedReturn() {
	err := testNamedReturnInner()
	fmt.Printf("err2:%v\n", err)
}

func main() {
	testNamedReturn()
}

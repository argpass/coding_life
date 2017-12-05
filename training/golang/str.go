package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

//1.len(s) 是按byte来计算string的长度的
//2.[]rune 来存储s后，len是按rune来计算
//3.rune(byte) 与 utf8.RuneSelf 比较来看byte是不是一个rune的开始
//4.range 按照rune来便利string
func testRune() {
	var s string = "我的天"
	fmt.Println("s is :", s)
	fmt.Println("len(s) is ", len(s))
	fmt.Println("utf8.RuneSelf: ", utf8.RuneSelf)
	fmt.Println("rune(s[0]) is rune start? ", rune(s[0]) >= utf8.RuneSelf)

	// ru is rune (int32 type)
	for _, ru := range s {
		fmt.Println(ru)
		fmt.Println(reflect.ValueOf(ru).Type())
	}

	// ru is rune (int32 type) also
	for _, ru := range "abc" {
		fmt.Println(ru)
		fmt.Println(reflect.ValueOf(ru).Type())
	}

	var r = ([]rune)(s)
	fmt.Println("r is :", string(r))
	fmt.Println("len(r) is ", len(r))
}

func testStrings() {
	fmt.Println(`len "" is `, len(""))
	fmt.Println(strings.Join([]string{"name"}, " "))
}

func testChar()  {
	s := "50d"
	var unit byte = s[len(s) - 1]
	fmt.Printf("get unit:%s\n", (unit))
}

func main() {
	//testRune()
	//testStrings()
	//testChar()
}

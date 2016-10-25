package main

import "fmt"

type Printer interface {
	Print(string)
}

type Base struct{}

type SimplePrinter struct {
	*Base
}

func (p *SimplePrinter) Print(s string) {
	fmt.Println(s)
}

type MapFunc func(x interface{}) (v interface{})

func testTypeSwitch() {
	// switch type
	var i interface{} = 99
	switch tp := i.(type) {
	case int:
		fmt.Println("tp case int")
	default:
		fmt.Println("tp is ", tp)
	}

	// switch type
	fmt.Println("test i64")
	var i64 int64 = 990
	var s string = ""
	i = i64
	switch tp := i.(type) {
	case int64:
		fmt.Println("tp case int64")
		// now i can set tp to a int64 variable
		i64 = tp
	case int, bool:
		// now tp is interface{} value not int or bool type
		// typeassert tp before to use it
		fmt.Println("tp is int or bool")
	case string:
		s = tp
		fmt.Println(s)
	default:
		fmt.Println("tp is ", tp)
	}

	// test rune
	var rus = []rune{'我', '的', '天'}
	var r interface{} = rus[0]
	switch tp := r.(type) {
	case rune:
		fmt.Println("got rune ", string(tp))
	}
}

func main() {
	var p Printer = &SimplePrinter{}
	var f MapFunc
	fmt.Println("f is nil ? ", f == nil)
	f2 := func(x interface{}) (v interface{}) { return x }
	fmt.Println("convert func:", MapFunc(f2))

	// type assert
	p.(*SimplePrinter).Print("assert successfully")

	if sp, ok := p.(*SimplePrinter); ok {
		fmt.Println("it's SimplePrinter")
		sp.Print("done")
	}

	var b interface{} = &SimplePrinter{}
	if _, ok := b.(*Base); ok {
		fmt.Println("it's *Base")
	} else {
		fmt.Println("it's not *Base")
	}

	// can't cast sub struct to base
	// (*Base)(&SimplePrinter{})
	testTypeSwitch()
}

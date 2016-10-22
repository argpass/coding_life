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

	// switch type
	var i interface{} = 99
	switch tp := i.(type) {
	case int:
		fmt.Println("tp case int")
	default:
		fmt.Println("tp is ", tp)
	}
}

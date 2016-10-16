package main

import "fmt"

type Printer interface {
	Print(string)
}

type Base struct{}

type SimplePrinter struct {
	Base
}

func (p *SimplePrinter) Print(s string) {
	fmt.Println(s)
}

func main() {
	var p Printer = &SimplePrinter{}
	if sp, ok := p.(*SimplePrinter); ok {
		fmt.Println("it's SimplePrinter")
		sp.Print("done")
	}

	var b interface{} = &SimplePrinter{}
	if _, ok := b.(*Base); ok {
		fmt.Println("it's *Base")
	}

	// can't cast sub struct to base
	(*Base)(&SimplePrinter{})
}

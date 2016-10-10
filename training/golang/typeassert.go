package main

import "fmt"

type Printer interface {
	Print(string)
}

type SimplePrinter struct{}

func (p *SimplePrinter) Print(s string) {
	fmt.Println(s)
}

func main() {
	var p Printer = &SimplePrinter{}
	if sp, ok := p.(*SimplePrinter); ok {
		fmt.Println("it's SimplePrinter")
		sp.Print("done")
	}
}

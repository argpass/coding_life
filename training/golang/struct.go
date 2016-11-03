package main

import (
	"fmt"
)

type Parent struct{}

type Child struct {
	*Parent
}

type VChild struct {
	Parent
}

func foo(pp *Parent) {
	fmt.Println(pp)
}

func main() {
	var pc = &Child{}
	// bad
	foo(pc)
}

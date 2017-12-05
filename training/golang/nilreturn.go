package main

import "fmt"

type any interface {
}

type fooType struct {
}

type fooInterface interface {
	Foo()
}

func testFoo() any {
	var f *fooType
	return f
}

func testFooInterface() any {
	var f fooInterface
	return f
}

func testFoo2() *fooType {
	var f *fooType
	return f
}

func testNamedFooReturn() (v any) {
	var f *fooType
	v = f
	return
}

func main()  {
	// false
	fmt.Println(testNamedFooReturn() == nil)
	// true
	fmt.Println(testFoo2() == nil)
	// false
	fmt.Println(testFoo() == nil)
	// true
	fmt.Println(testFooInterface() == nil)
}

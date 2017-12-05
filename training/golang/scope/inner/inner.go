package inner

import "fmt"

type MyType struct {
	a int
	b int
}

func (m MyType) Say()  {
	fmt.Printf("hello from inner, a %d, b %d\n", m.a, m.b)
}

func GetMyType() MyType {
	return MyType{}
}


package main

import (
	"fmt"
	"math"
)

func foo(i int) int {
	if i > (math.MaxInt64 -10 ){
		return 99
	}
	return foo(i+1)
}

func main()  {
	foo(0)
	fmt.Println("done")
}

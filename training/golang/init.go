// multi init function in a single package
// the order is filename order
// multi init function in a single file is ok
package main

import (
	"fmt"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	fmt.Println("main")
}

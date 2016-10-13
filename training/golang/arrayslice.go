package main

import (
	"fmt"
)

func main() {
	var s []int
	// s is nil
	fmt.Println(s == nil)

	// space is assigned
	s = []int{}
	fmt.Println(s != nil)

	// s is assigned space when used
	s = append(s, 99)
	fmt.Println(s)
	fmt.Println(error("ok"))
}

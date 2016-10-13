package main

import (
	"fmt"
)

func main() {
	var arr = []int{3, 4, 5}

	// only got index of the slice
	fmt.Println("\n-------------only index --------")
	for d := range arr {
		fmt.Print(d, ",")
	}

	fmt.Println("\n------------data--------")
	for _, d := range arr {
		fmt.Print(d, ",")
	}
}

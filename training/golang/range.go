package main

import (
	"fmt"
	"time"
)

func main() {
	var arr = []int{3, 4, 5}
	fmt.Printf("now:%v\n", time.Now().Format("060102 15:04:05"))

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

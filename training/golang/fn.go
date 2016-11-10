package main

import (
	"fmt"
)

func namedReturn() (ret int) {
	ret = 99
	return ret
}

func main() {
	fmt.Println(namedReturn())
}

package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a uint32 = 9
	fmt.Println(^uint32(0))
	fmt.Println(^int32(0))
	atomic.AddUint32(&a, ^uint32(0))
	fmt.Println(a)
}

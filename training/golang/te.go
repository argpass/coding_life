package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	c1 := make(chan int, 1)
	for i := 0; i < 5; i++ {
		go func(i int) { c1 <- i }(i)

	}
	for i := 0; i < 5; i++ {
		println(<-c1)
	}

}

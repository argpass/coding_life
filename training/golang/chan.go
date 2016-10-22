// 1.chan 可以设置只读，只写，最小权限模式

package main

import (
	"fmt"
)

// use select to test whether channel is full or empty
// select only can use io operate
func testSelect() {
	ch := make(chan []byte, 2)
	// len method can test whether channel is empty
	fmt.Println("chan len is:", len(ch))
	// cap and len  can test whether channel is full
	fmt.Println("chan cap is full ?:", cap(ch) == len(ch))

	select {
	case <-ch:
	default:
		fmt.Println("chan is empty")
	}

	ch <- make([]byte, 1)
	ch <- make([]byte, 1)

	select {
	case ch <- make([]byte, 1):
	default:
		fmt.Println("chan is full")
	}
}

func readOnlyChan() {
	foo := func(ch <-chan int) {
		i := <-ch
		// can not write chan
		// ch <- 4
		fmt.Println("got ", i)
	}
	ch := make(chan int, 1)
	ch <- 99
	foo(ch)

}

func main() {
	readOnlyChan()
	testSelect()
}

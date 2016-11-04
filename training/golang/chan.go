// 1.chan 可以设置只读，只写，最小权限模式
// 1.close chan , 所有的'<-'都能收到信号

package main

import (
	"fmt"
	"sync"
	"time"
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

// close chan , 所有的<-都能收到信号
func closeChan() {
	var w sync.WaitGroup
	var exitChan = make(chan int, 1)
	for i := 0; i < 3; i++ {
		w.Add(1)
		go func(index int) {
			defer w.Done()
			for {
				select {
				case <-exitChan:
					goto exit
				default:
				}
				fmt.Println("working #-", index)
				time.Sleep(1000000 * 1)
			}
		exit:
			fmt.Println("exit #", index)
		}(i)
	}
	time.Sleep(1000000 * 2)
	close(exitChan)
	w.Wait()
	fmt.Println("closeChan done")
}

func echoClosedChan() {
	var done = make(chan int, 1)
	fmt.Println("chan is nil ? ", nil == done)
	close(done)
	fmt.Println("closed chan is nil ? ", nil == done)
}

func main() {
	echoClosedChan()
	closeChan()
	readOnlyChan()
	testSelect()
}

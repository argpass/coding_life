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

// close chan , 所有的<-都能收到信号， 判断chan关闭可以使用第二参数bool值
func closeChan() {
	var w sync.WaitGroup
	var exitChan = make(chan int, 1)
	for i := 0; i < 3; i++ {
		w.Add(1)
		go func(index int) {
			defer w.Done()
			for {
				select {
				case got, ok := <-exitChan:
					fmt.Println(got, ok)
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
	select {
	case <-done:
		fmt.Println("select closed chan done")
	default:
		fmt.Println("select closed chan fail")
	}
	select {
	case <-done:
		fmt.Println("select closed chan done again")
	default:
		fmt.Println("select closed chan fail")
	}
}

// 1.可以select收发nil chan，但都是失败
// 2.发送nil到非nil chan是可以的 (nil可以转为chan接受的类型时)
func selectNilChan() {
	var d int = 99
	var sendChan chan<- *int
	var doneChan <-chan *int
	select {
	case sendChan <- &d:
		fmt.Println("send nil to nil sendChan")
	default:
		fmt.Println("send nil fail to nil chan")
	}
	select {
	case <-doneChan:
		fmt.Println("select success from nil chan")
	default:
		fmt.Println("select fail from nil chan")
	}
	sendChan = make(chan *int, 1)
	select {
	case sendChan <- nil:
		fmt.Println("send nil to non-nil sendChan successfully")
	default:
		fmt.Println("send nil fail to none-nil chan")
	}
}

// 测试select 多个项有效时，是顺序还是随机选择读写,
// 是随机选择
func multiSelect() {
	var ch_a = make(chan int, 1)
	var ch_b = make(chan int, 1)
	var ch_c = make(chan int, 1)
	ch_a <- 1
	ch_b <- 2
	ch_c <- 3
	select {
	case <-ch_a:
		fmt.Println("a")
	case <-ch_b:
		fmt.Println("b")
	case <-ch_c:
		fmt.Println("c")
	}
}

func main() {
	//multiSelect()
	//selectNilChan()
	//echoClosedChan()
	closeChan()
	//readOnlyChan()
	//testSelect()
}

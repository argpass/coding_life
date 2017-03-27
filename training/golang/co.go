// 1.并发读写贡献map时需要加锁否则十分容易触发panic
// 2.并发只读不需要加锁
// 3.并发只写需要加锁

package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

// 实验并发读map不用加锁
func concurrencyReadMap() {
	var pool = map[string]int{}
	num := 2000
	ch := make(chan int, num)
	for i := num; i > 0; i-- {
		go func(index int, po map[string]int, c chan int) {
			if index%2 == 0 {
			} else {
				for j := 100; j > 0; j-- {
					if m, ok := po[fmt.Sprintf("#%d-%d", index, j)]; ok {
						fmt.Println(m)
					}
				}
			}
			c <- index
		}(i, pool, ch)
	}
	for t := num; t > 0; t-- {
		<-ch
	}
	fmt.Println("concurrency read map test bye")
}

// 实验并发写map,需要加锁，否则panic
func concurrencyWriteMap() {
	var pool = map[string]int{}
	num := 20000
	var lock = &sync.RWMutex{}
	ch := make(chan int, num)
	for i := num; i > 0; i-- {
		go func(index int, po map[string]int, c chan int) {
			if index%2 == 0 {
				for j := 100; j > 0; j-- {
					lock.Lock()
					po[fmt.Sprintf("#%d-%d", index, j)] = j
					lock.Unlock()
				}
			}
			c <- index
		}(i, pool, ch)
	}
	for t := num; t > 0; t-- {
		<-ch
	}
	fmt.Println("pool len:", len(pool))
	fmt.Println("concurrencyWriteMap test bye")
}

// 并发写map的时候，读map也要加锁才能读
func concurrencyWRMap() {
	var pool = map[string]int{}
	num := 20000
	var lock = &sync.RWMutex{}
	ch := make(chan int, num)
	for i := num; i > 0; i-- {
		go func(index int, po map[string]int, c chan int) {
			if index%2 == 0 {
				for j := 100; j > 0; j-- {
					lock.Lock()
					po[fmt.Sprintf("#%d-%d", index, j)] = j
					lock.Unlock()
				}
			} else {
				for j := 100; j > 0; j-- {
					lock.RLock()
					m, ok := po[fmt.Sprintf("#%d-%d", index, j)]
					lock.RUnlock()
					if ok {
						fmt.Println("got m:", m)
					}
				}
			}
			c <- index
		}(i, pool, ch)
	}
	for t := num; t > 0; t-- {
		<-ch
	}
	fmt.Println("pool len:", len(pool))
	fmt.Println("concurrencyWriteMap test bye")
}

func concurrencyChan() {
	ch := make(chan int, 2)
	for i := 10; i > 0; i-- {
		go func() {
			for j := 0; j < 10000; j++ {
				if len(ch) != 0 {
					fmt.Println("len chan:%d", len(ch))
				}
			}
		}()
	}
	go func() {
		ch <- 1
		ch <- 2
	}()
	time.Sleep(1000 * 500)
	<-ch
	<-ch
}

// 并发读写变量:并不会发生并发异常，可以不用锁,
func concurrencyVar()  {
	s := 88
	var v = &s
	for i := 0; i < 30000; i++ {
		go func(j int){
			count := 0
			for {
				if count > 10 {
					break
				}
				if j % 2 == 0 {
					if count % 2 != 0 {
						*v = 99
					}else{
						*v = 77
					}
				}else{
					if count % 2 == 0 {
						fmt.Printf("read *v:%d\n", *v)
					}
				}
				runtime.Gosched()
				count++
			}
		}(i)
	}
}

// 实验并发写map,即使是map的引用拷贝也需要加锁，否则panic
func concurrencyWriteMap2() {
	var pool = map[string]int{}
	num := 20000
	//var lock = &sync.RWMutex{}
	ch := make(chan int, num)
	for i := num; i > 0; i-- {
		p := pool
		go func(index int, po map[string]int, c chan int) {
			if index%2 == 0 {
				for j := 100; j > 0; j-- {
					//lock.Lock()
					po[fmt.Sprintf("#%d-%d", index, j)] = j
					//lock.Unlock()
				}
			}
			c <- index
		}(i, p, ch)
	}
	for t := num; t > 0; t-- {
		<-ch
	}
	fmt.Println("pool len:", len(pool))
	fmt.Println("concurrencyWriteMap test bye")
}

func main() {
	//concurrencyChan()
	//concurrencyWRMap()
	//concurrencyReadMap()
	//concurrencyWriteMap()
	//concurrencyVar()
	concurrencyWriteMap2()
}

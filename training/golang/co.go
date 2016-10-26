// 1.并发读写贡献map时需要加锁否则十分容易触发panic
// 2.并发只读不需要加锁
// 3.并发只写需要加锁

package main

import (
	"fmt"
	"sync"
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

func main() {
	concurrencyWRMap()
	//concurrencyReadMap()
	//concurrencyWriteMap()
}

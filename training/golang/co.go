// 1.并发读写贡献map时需要加锁否则十分容易触发panic
// 2.并发只读不需要加锁
// 3.并发只写需要加锁

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
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
func concurrencyVar() {
	s := 88
	var v = &s
	for i := 0; i < 30000; i++ {
		go func(j int) {
			count := 0
			for {
				if count > 10 {
					break
				}
				if j%2 == 0 {
					if count%2 != 0 {
						*v = 99
					} else {
						*v = 77
					}
				} else {
					if count%2 == 0 {
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

func testGoDeadLoop() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go func(n int) {
			for {
				time.Sleep(1 * time.Microsecond)
				fmt.Println("go ", n)
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
	for {
		runtime.Gosched()
		// 其他go程只能调度一会儿就没法得到调度了
	}
}

func testGoSchedule() {
	for i := 0; i < 241767266; i++ {
		if i%1234 == 1 {
			go func() {
				fmt.Printf("inner go i:%d\n", i)
			}()
		}
	}
}

// 读锁排写，但不拒绝读 go1 go2 可以造成死锁(都能获得读锁)
func testRLock() {
	mu := &sync.RWMutex{}
	go func() {
		mu.RLock()
		fmt.Println("go 1")
		//mu.RUnlock()
	}()
	go func() {
		mu.RLock()
		fmt.Println("go 2")
		//mu.RUnlock()
	}()
	time.Sleep(1 * time.Second)
}

// slice是并发安全的
func testSliceConcurrency() {
	data := []int{1, 2, 3}
	go func() {
		for {
			data[1] = int(time.Now().Unix())
			runtime.Gosched()
		}
	}()
	go func() {
		for {
			for _, v := range data {
				fmt.Printf("read v:%d\n", v)
				runtime.Gosched()
			}
		}
	}()
	go func() {
		for {
			for i := 0; i < 1000; i++ {
				data = append(data, i*2)
				runtime.Gosched()
			}
			data = data[:3:3]
		}
	}()
	select {}
}

func main() {
	//concurrencyChan()
	//concurrencyWRMap()
	//concurrencyReadMap()
	//concurrencyWriteMap()
	//concurrencyVar()
	//concurrencyWriteMap2()
	//testGoDeadLoop()
	//testGoSchedule()
	//testRLock()
	testSliceConcurrency()
}

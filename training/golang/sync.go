package main

import (
	"fmt"
	"sync"
	"time"
)

// 需要调用w.Wait()才能阻塞主线程
func waitGroup() {
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer w.Done()
		time.Sleep(10)
		fmt.Println("inner done")
	}()
	fmt.Println("noWait done")
	w.Wait()
}

func main() {
	waitGroup()
	fmt.Println("main done")
}

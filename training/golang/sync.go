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

type Locker struct {
	lock sync.RWMutex
	doneC chan struct{}
}

func (l *Locker) HoldLock()  {
	l.lock.Lock()
	//defer l.lock.Unlock()
	time.Sleep(5 * time.Second)
	fmt.Println("hold done")
}

func (l *Locker) Acquire()  {
	l.lock.Lock()
	defer l.lock.Unlock()
	fmt.Println("got lock and do work")
	close(l.doneC)
}

// 测试锁等待
func testLock()  {
	locker := &Locker{doneC:make(chan struct{})}
	startC := make(chan struct{})
	go func(){
		close(startC)
		locker.HoldLock()
	}()
	<-startC
	go func(){
		locker.Acquire()
	}()
	<-locker.doneC
}

func main() {
	//waitGroup()
	testLock()
	fmt.Println("main done")
}

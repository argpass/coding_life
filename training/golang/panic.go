package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"time"
)

// 1.panic 前的defer会被执行，之后的当然不会被执行
// 2.os.Exit 不会执行defer
func testDefer()  {
	
	defer func(){
		fmt.Printf("panic defer")
	}()
	os.Exit(12)
	err := errors.New("test err")
	panic(err)
	
	fmt.Println("fn done")
}

// panic会导致进程退出
func testPanicInCoroutine()  {
	go func() {
		func(){
			fmt.Println("coroutine is running")
			panic(errors.New("404"))
		}()
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main coroutine bye")
}

// recover只能捕获所在协程中的panic
func testPanicRecover()  {
	go func() {
		defer func(){
			panicError := recover()
			fmt.Printf("panic err:%v\n", panicError)
		}()
		func(){
			fmt.Println("coroutine is running")
			panic(errors.New("404"))
		}()
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main coroutine bye")
}

func main()  {
	//testDefer()
	testPanicInCoroutine()
}

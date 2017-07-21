// 1.panic 后能触发defer
// 2.os.Exit 立即退出进程，不会出发defer
// 3.defer 顺序是倒序(包括defer 中的内嵌defer)
// 4.有名返回值,返回值可以在defer中得到修改; 使用defer可以在defer中观察函数的返回值，甚至修改返回值(针对有名返回值的情况)
package main

import (
	"fmt"
	"os"
)

// order is #2 #1 inner
func testDeferOrder() {
	defer func() {
		fmt.Println("defer #1")
		defer func() {
			fmt.Println("defer inner")
		}()
	}()

	defer func() {
		fmt.Println("defer #2")
	}()
}

// exit后defer不能得到执行
func testDeferOnExit() {
	defer func() {
		fmt.Println("defer on exit")
	}()

	os.Exit(1)
}

// panic 后能触发defer
func testDeferOnPanic() {
	defer func() {
		fmt.Println("defer on panic")
	}()

	panic("panic exit")
}

func deferReturn() {
	// 普通返回
	fn := func() int {
		defer func() int {
			// defer 的返回不会影响函数的返回
			return 9
		}()
		return 8
	}
	i := fn()
	fmt.Println("got return ", i)

	// 有名返回值,返回值可以在defer中得到修改
	fn = func() (r int) {
		defer func() int {
			r++
			return 9
		}()
		return r
	}
	i = fn()
	fmt.Println("#2 got return ", i)
}

func testDeferClosure()  {
	var code int
	defer func(){
		if code == 99 {
			fmt.Println("code is 99")
		}
	}()
	code = 99
}

func main() {
	//deferReturn()
	//testDeferOrder()
	//testDeferOnPanic()
	//testDeferOnExit()
	testDeferClosure()
}

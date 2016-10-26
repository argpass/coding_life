// 1.panic 后能触发defer
// 2.os.Exit 立即退出进程，不会出发defer
// 3.defer 顺序是倒序(包括defer 中的内嵌defer)
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

func main() {
	testDeferOrder()
	testDeferOnPanic()
	testDeferOnExit()
}

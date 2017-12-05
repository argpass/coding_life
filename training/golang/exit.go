package main

import (
	"fmt"
	"time"
)

func foo() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("foo go")
	}
}

func main() {
	go func() {
		defer func() {
			recover()
		}()
		go foo()
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
			panic("timeout")
		}
	}()
	<-make(chan struct{})
}

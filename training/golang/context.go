package main

import (
	"context"
	"sync"
	"fmt"
)

func testStopContext(wg *sync.WaitGroup)  {
	ctx, stop := context.WithCancel(context.Background())
	ctx, stop2 := context.WithCancel(ctx)
	wg.Add(1)
	go func(){
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("done")
		}
	}()
	stop()
	stop2()
	fmt.Println("innder done")
}

func main()  {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(){
		defer wg.Done()
		testStopContext(wg)
	}()
	wg.Wait()
	fmt.Println("main bye")
}

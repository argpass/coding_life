package main

import (
	"golang.org/x/sync/errgroup"
	"context"
	"fmt"
	"time"
	"github.com/pkg/errors"
)

// 返回non-nil error 会取消ctx（触发与之关联的go程退出)
func testerrgroup()  {
	ctx, stop := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		for{
			select {
			case <-ctx.Done():
				err := ctx.Err()
				fmt.Printf("got err:%v\n", err)
				return err
			}
		}
		return nil
	})
	// 返回非nil error 会取消ctx
	g.Go(func() error {
		time.Sleep(20 * time.Second)
		return errors.New("i'm donw")
	})
	stop()
	g.Wait()
	fmt.Println("main bye")
	
}

func main()  {
	testerrgroup()
}

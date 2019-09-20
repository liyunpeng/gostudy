package context1

import (
	"context"
	"fmt"
	"time"
)

func context1() {
	ctx, cancel := context.WithCancel(context.Background())

	key := "key1"
	valueCtx := context.WithValue(ctx, key, "add value")

	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	cancel()

	//time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	key := "key1"
	for {
		select {
		case <-ctx.Done():
			//get value
			fmt.Println(ctx.Value(key), "is cancel")

			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")

			time.Sleep(2 * time.Second)
		}
	}
}

func Context()  {
	context1()

	context2()
}

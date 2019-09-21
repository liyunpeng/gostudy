package context1

import (
	"context"
	"fmt"
	"time"
)

func context1() {
	/*
		创建一个可以随时取消的上下文
		在上下文里可以创建键值对

	*/
	ctx, _ := context.WithCancel(context.Background())

	key := "key1"
	valueCtx := context.WithValue(ctx, key, "add value")

	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	/*
		cancel和ctx.done都是结束上下文
	*/
	ctx.Done()
	//cancel()

	//time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	key := "key1"
	for {
		select {
		/*
			调用cancel即是向done管道里些个值， 这样在done的鼓捣读的阻塞动作借解除了
		*/
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "is cancel")

			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")

			time.Sleep(2 * time.Second)
		}
	}
}

func Context() {
	context1()

	context2()
}

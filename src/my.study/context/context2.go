package context1

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func work(ctx context.Context) error {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		select {
		/*
			time.After是系统的定时器，指定多少时间就会在指定的时间发送这个解除读阻塞的动作， 基本所有阻塞都是管道的阻塞
		*/
		case <-time.After(2 * time.Second):
			fmt.Println("Doing some work ", i)

			// we received the signal of cancelation in this channel
			/*
				到指定时间就会发送写done管道， done管道读的阻塞动作就会解除， 表示context结束了
				这个只表示开始执行退出动作了
			*/
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			/*
				在return前wg.done, 表示退出动作执行完毕
			*/
			return ctx.Err()
		}
	}
	return nil
}

func context2() {

	/*
		创建指定存活时间的上下文
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	fmt.Println("Hey, I'm going to do some work")

	wg.Add(1)
	go work(ctx)
	wg.Wait()

	fmt.Println("Finished. I'm going home")
}

package basic

import (
	"fmt"
	"sync"
)

//   chan<- //只写
func producer(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		fmt.Println("produce: ", i)
		out <- i //如果对方不读 会阻塞
	}
}

//   <-chan //只读
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("consume: ", num)
	}
}

func producerConsumer() {

	c := make(chan int) //   chan   //读写

	go producer(c) //生产者

	consumer(c) //消费者

	fmt.Println("done")
}

func sync1() {
	wg := sync.WaitGroup{}

	const synNum int = 5
	wg.Add(synNum)
	for i := 0; i < synNum; i++ {
		go func(i int) {
			fmt.Println(i, " sync1 done-1")
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(" sync1 end")
}

func closeChan() {
	done := make(chan struct{})
	go func() {
		/*
		这是个bug, 原因：主程序已经把done通道close了， 但此处读通道的动作还是阻塞在这里
		从<- done routine exit没有打印出来，可以断定此routine在主程序结束前没有退出， 后通过实验
		得知， close(done)是能解除所有读done通道的阻塞操作， 稍微会晚一点时间，1秒左右的时间， 但是主程序退出了
		导致1秒后routine的读done通道阻塞解除时， 主程序结束，对命令函的输出控制也就结束了，导致routine不能向控制台输出
		所以close通道， 能解除routine里读通道的阻塞，只是时间稍微延后
		为了放心，每个routine结尾都要打印一个结束语句，这个不是调试语句，应该是info语句，而且要打印到文件里， 因为控制台已经结束了
		 */
		<-done
		fmt.Println("<- done closeChan routine exit")
	}()
	close(done)
	fmt.Println("closeChan  programe  exit")
}

func closeChan1() {
	done := make(chan struct{})
	go func() {
		for {
			select {

			case <-done:
				fmt.Println("<- done closeChan1 routine exit")
				return
			}
		}

	}()
	close(done)
	fmt.Println("closeChan1  programe  exit")
}

func closeChan2() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	wq := make(chan interface{})
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i, wq, done, &wg)
	}
	for i := 0; i < workerCount; i++ {
		wq <- i
	}
	close(done)
	wg.Wait()
	fmt.Println("all done!")
}
func doit(workerId int, wq <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()
	for {
		select {
		case m := <-wq:
			fmt.Printf("[%v] m => %v\n", workerId, m)
		case <-done:
			fmt.Printf("[%v] is done\n", workerId)
			return
		}
	}
}

func Chan() {
	fmt.Println("<------------------------- Chan begin -------------------->")
	producerConsumer()

	sync1()

	closeChan1()

	closeChan2()

	closeChan()

	fmt.Println("<------------------------- Chan end -------------------->")
}

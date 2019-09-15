package basic

import (
	"fmt"
	"sync"
)

//   chan<- //只写
func producer(out chan<- int) {
	defer close(out)  // 在最后一个写通道动作后，close通道
	for i := 0; i < 5; i++ {
		fmt.Println("produce: ", i)
		out <- i //如果对方不读 会阻塞
	}
}

//   <-chan //只读
func consumer(in <-chan int) {

	for num := range in {   // range无缓存通道, 要求必须在最后一个写通道的后面，close通道
		fmt.Println("consume: ", num)
	}
}

func producerConsumer() {

	c := make(chan int) //   chan   //读写

	go producer(c) //生产者

	consumer(c) //消费者

	fmt.Println("done")
}

func writeTwoChan() {
	ch := make(chan string)
	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
			//global.Logger.Println("processed:", m)
		}

	}()

	/*
		这是bug2
		一般在主程序里都是读通道的动作， 现在把两个写通道放在主程序里，是为了验证个结论，
		我们故意把这个函数放在main函数最后执行， 最后程序只输出了
		processed: cmd.1
		没有输出 processed: cmd.2
		原因是在主程序在写了管道之后， routine的读管道的阻塞的解除是需要一点时间的，
		而主程序在写完管道后，输出一条语句后，就直接退除了，也就失去了对控制台的输出权利
		这时routine再向控制台输出时，控制台是接收不到的。 想看到routine的完整输出打印，有两个办法
		一是在最后一个写管道后等待两秒，即添加
		time.Sleep(2*time.Second)
		另一种办法是把输出的log打印到文件里，即在routine里用下面语句打印：
		global.Logger.Println("processed:", m)
		bug1和bug2的原因相同，都是主程序在routine还没结束就退出了
	 */
	ch <- "cmd.1"
	ch <- "cmd.2" //won't be processed
	//time.Sleep(2*time.Second)
	close(ch)
	fmt.Println("writeTwoChan end")
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
			这是bug1,
			原因：从<- done routine exit没有打印出来，可以断定此routine在主程序退出前没有结束，
			通过实验得知， close(done)能够解除所有读done通道的阻塞操作，但是稍微会晚一点时间，1秒左右的时间， 但是主程序退出了
			导致1秒后routine的读done通道阻塞解除时， 主程序已经结束了，对控制台的输出就结束了，导致routine不能向控制台输出。
			所以close通道， 能解除routine里读通道的阻塞，只是时间稍微延后
			为了能从log判断出routine是否正常退出，每个routine的结尾都要打印一个结束log info，不是只有在debug时才打开的log.
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
	/*
		命名习惯: 用于信号通知的通道命名为done, 用于生产者消费者模式的通道命名为wq
	 */
	done := make(chan struct{})
	wq := make(chan interface{})
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i, wq, done, &wg)
	}

	/*
		通道被用在两种情况：
	   	一，用于生产者消费者模式，
	    判断标准： 有连续写的动作就是被用于生产者消费者模式， 这里连续写两次，这里的wq即是， wq是work queue缩写
		二，用于信号通知，
	 	判断标准： 只有一次写，或没有写，只有一个close动作， done通道即是
	*/
	for i := 0; i < workerCount; i++ {

		wq <- i
	}

	/*
	 解除所有的读通道， 不管是几次读，
	 这里读通道阻塞的routine被启动了两次， 就有两个读通道阻塞，close将这两个阻塞全部解除
	*/
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

	writeTwoChan()
	fmt.Println("<------------------------- Chan end -------------------->")
}
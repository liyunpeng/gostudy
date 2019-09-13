package basic

import "fmt"

//   chan<- //只写
func producer(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {

		out <- i //如果对方不读 会阻塞
	}
}

//   <-chan //只读
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func Chan() {

	c := make(chan int) //   chan   //读写

	go producer(c) //生产者

	consumer(c) //消费者

	fmt.Println("done")
}

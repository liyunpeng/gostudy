package gopoll

import (
	"fmt"
	"runtime"
	"time"
)

type Score struct {
	Num int
}

func (s *Score) Do() {
	fmt.Println("num:", s.Num)
	//这里延迟是模拟处理数据的耗时
	time.Sleep(1 * 1 * time.Second)
}

func GopollMain() {
	//num := 100 * 100 * 20
	num := 2
	// debug.SetMaxThreads(num + 1000) //设置最大线程数
	// 注册工作池，传入任务
	// 参数1 worker并发个数
	p := NewWorkerPool(num)
	p.Run()

	//datanum := 100 * 100 * 100 * 100
	datanum := 5
	go func() {
		for i := 1; i <= datanum; i++ {
			sc := &Score{Num: i}
			p.JobQueue <- sc
			p.JobQueue <- sc
			p.JobQueue <- sc
			p.JobQueue <- sc
			p.JobQueue <- sc
		}
	}()

	for {
		fmt.Println("启动的routine个数统计： runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}

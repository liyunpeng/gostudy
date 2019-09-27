package basic

import "fmt"

func defer_call() {
	/*
		defer 的执行顺序是后进先出。
		当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，
		最后才会执行panic。
			输出结果
			打印后
			打印中
			打印前
			panic: 触发异常
	*/
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

func Defer() {

	defer_call()
}

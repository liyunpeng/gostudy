package basic

import (
	"fmt"
	"runtime"
)

func run1() {

	/*
		GOMAXPROCS表示了CPU的数量，Go将使用这个数量来运行goroutine。
	    而runtime.GOMAXPROCS()函数的文档让人更加的迷茫
	*/
	fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 1
	fmt.Println(runtime.NumCPU())       //prints: 1 (on play.golang.org)
	runtime.GOMAXPROCS(20)
	fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 20
	runtime.GOMAXPROCS(300)
	fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 256
}

func Runtime()  {
	fmt.Println("<-------------------------------- Runtime ----------------------->")
	run1()
	fmt.Println("<-------------------------------- Runtime ----------------------->")
}
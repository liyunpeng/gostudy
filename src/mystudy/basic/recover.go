package basic

import "fmt"

func recover1() {
	fmt.Println("c")
	/*
	 必须要先声明defer，否则不能捕获到panic异常
	*/
	defer func() {
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		fmt.Println("e")
	}()
	f1()
	fmt.Println("f") //这里开始下面代码不会再执行
}

func f1() {
	fmt.Println("a")
	panic("异常信息")
	fmt.Println("b") //这里开始下面代码不会再执行
}

/*
-------output-------
c
a
d
异常信息
e
*/

func Recover() {
	recover1()
}

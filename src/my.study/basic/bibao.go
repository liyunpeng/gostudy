package basic

import "fmt"

/*
	go的闭包（Closure）定义: 引用了外部变量的匿名函数
 */

/*
	函数胡返回值类型可以是函数
*/
func func1() func() int {
	/*
		因为没有引用外部变量, 所以返回的这个匿名函数还联不是闭包
	 */
	return func() int {
		return 10
	}
}

func closer1() {
	str := "hello world"

	/*
		func()定义了一个匿名函数
	*/
	foo := func() {
		/*
			这个匿名函数引用了外部变量, 所以这个匿名函数就成为胃闭包
		 */
		str = "hello dude"
	}

	foo()
}

func Accumulate(value int) func() int {
	/*
		返回了一个闭包
	 */
	return func() int {
		value++
		return value
	}
}

func closer2() {
	accumulator := Accumulate(1)
	fmt.Println("Accumulate()=", accumulator())
	fmt.Println("Accumulate()=", accumulator())
	/*
		打印函数地址
	 */
	fmt.Printf("Accumulate=%p\n", accumulator)

	/*
		返回值为函数的作用：
		有点函数指针的感觉， 在运行时， 动态的指向不同的函数，并执行， 体现了运行时概念
	 */
	accumulator = func1()
	fmt.Println("func1()=", accumulator())

	fmt.Printf("func1=%p\n", accumulator)

	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", accumulator2)

	/*
		对比输出的日志发现 accumulator 与 accumulator2 输出的函数地址不同，
		因此它们是两个不同的闭包实例。
	*/
}

func Closer1() {
	fmt.Println("<--------------------------- Closer begin---------------------> ")
	f1 := func1()

	closer1()

	closer2()
	fmt.Println(f1())
	fmt.Println("<--------------------------- Closer end---------------------> ")
}

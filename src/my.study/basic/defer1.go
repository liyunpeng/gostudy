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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic在此得到处理， 不会使程序崩溃")
		}
		fmt.Println("打印后")
	}()
	panic("触发异常")
	//defer_call()
}

func closerReference() (r int) {
	/*
		闭包引用， 修改了外部变量r，这个变量是返回值变量，
		所以返回值在defer里被修改
	*/
	defer func() {
		r++
	}()
	return 0
}

func closerReferenceBackup() (r int) {
	t := 5
	/*
		r是t的拷贝， 而且这个拷贝的时间发生在defer之前，
		这又是一种defer前做的事情。 defer前除了return, 其他事情全做了， 返回值是变量的拷贝也在这个范围内
		所以虽然是闭包引用， 但这个引用的不是返回值，而是返回值的副本，并不能改到返回值本身
		所以返回值在defer里没被修改
	*/
	defer func() {
		t = t + 5
	}()
	return t
}

func pamaeter() (r int) {
	/*
		defer修改的是参数r, 并不是外部返回值变量r,
		所以返回值在defer里没被修改
	*/
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

/*
	defer里修改外部参数
*/
func testCloserReferenceDefer() {
	/*
		通过闭包引用修改外部参数
	*/
	closerReference()
	closerReferenceBackup()

	/*
		入参修改
	*/
	pamaeter()
}

func f(n int) (r int) {
	defer func() {
		fmt.Println("c")
		r += n
		/*
			因为第一个defer抛出了一个panic, recover是专门接收panic的
		*/
		recover()
	}()

	var f func()

	/*
		局部变量f在defer后面定义， 所以不被执行
		因为没有定义， 所以这里会出一个panic
		发生了panic， defer也是是按顺序执行
	*/
	defer f()

	f = func() {
		r += 2
		fmt.Println("b")
	} // f 只是定义，没有小括号， 所以没被执行

	fmt.Println("a")
	/*
		这里先返回了， 这时r是4，但返回这个4，虽然返回，但不是真返回， 因为f还没真正结束，
		所以这个返回值也不是调用者接收的，
		本函数的所有的defer执行完了才算真正的返回， 所以是7
	*/
	return n + 1
}

type Person struct {
	age int
}

func deferParameter() {
	person := &Person{28}

	/*
		不构成闭包引用
		person.age 此时是将 28 当做 defer 函数的参数，
		会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；
	*/
	defer fmt.Println(person.age)

	/*
		defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，
		所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29；
	*/
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	/*
		构成闭包引用，输出 29；
	*/
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29

	/*
		构成闭包引用， 虽然person指向了另一个结构体， 因为闭包引用， 还椒得到修改后的值
		defer func() {
			fmt.Println(person.age)
		}()

		person = &Person{29}
	*/
}

func Defer() {
	fmt.Println(f(3))

	defer_call()

	testCloserReferenceDefer()

	deferParameter()
}

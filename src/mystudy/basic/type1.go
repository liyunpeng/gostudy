package basic

import "fmt"

type MyInt1 int

/*
这种带等号的定义出的新类型， 只是定义了类型的别名，
int 和MyInt2 都是相同的类型， 而上面的
type MyInt1 int
定义出的MyInt1 和int是不同的类型
*/
type MyInt2 = int

/*
Go 是强类型语言的测试
*/
func type1() {
	var i int = 0
	/*
		MyInt1 和int是不同的类型， 所以
		var i1 MyInt1 = i
		编译不过，，编译不通过
	*/

	/*
		MyInt1 和int的别名， 是相同的类型，所以编译通过
	*/
	var i2 MyInt2 = i
	fmt.Println( i2)
}

func type2(){
	type t1 [10]int
	var s t1
	fmt.Println(s)
}

/*
func main() {
	a := 5
	b := 8.1
	a 的类型是 int，b 的类型是 float，两个不同类型的数值不能相加，编译报错。
	fmt.Println(a + b)
}
*/

type N int
func (n N) test(){
	fmt.Println("n=", n)
}



func type3(){
	var n N = 20

	/*
		调用类型的方法, 必须是调用类型实例的方法
	*/

	/*
		1. 以实例中成员的形式调用方法, 可以是实例, 也可以是指向实例的指针
	 */
	n.test()
	(&n).test()

	/*
		2. 以类型中成员的形式调用方法, 实例作为参数传入, 实例指针不能作为参数传入
		如N.test(&n) 不合法
	 */
	N.test(n)

	/*
		3.用函数变量方式调用, 实例以函数参数形式调用
	 */
	f := N.test
	f(n)
}

func zhengfu() {
	i := -5
	j := +5
	/*
		%d表示输出十进制数字，+表示输出数值的符号。这里不表示取反。
	*/
	fmt.Printf("%+d %+d", i, j)
}

func typeConvert(){
	a := 5.0
	b := int(a)
	/*
		b := a.(int)
		这样写a.(int)， 是非法的， a 只有是指针的时候才可以这样写
	*/

	fmt.Println(b)

	var a1 interface{}

	a1 = 1
	c := a1.(int)
	/*
		接口不能用如下方式类型转换， 语法错误
		c := int(a1)
	*/
	fmt.Println(c)

	/*
		对接口做类型转换时， 如果类型不是当初对接口变量赋值的类型， 运行时会有panic错误：
		d := a1.(string) 运行时， 报如下panic：
		panic : interface conversion: interface {} is int, not string
		goroutine 16 [running]:
	*/
	d := a1.(int)

	fmt.Println(d)
}

func Type1() {
	fmt.Println("<----------------------------- Type begin ---------------------------->")
	type1()
	type2()
	zhengfu()
	type3()
	typeConvert()
	fmt.Println("<----------------------------- Type end ---------------------------->")
}

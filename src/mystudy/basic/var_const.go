package basic

import "fmt"

/*
x := 100
这是变量声明的简短模式，但这种声明方式有三个限制：
1. 必须使用显示初始化；
2. 不能提供数据类型，编译器会自动推导；
3. 只能在函数内部使用简短模式； 不能再全局使用这种剪短方式
*/
var x int = 100

/*
全局变量可以只声明，不使用，不像局部变量，不使用就包编译错误
*/
var (
	ag int
	bg int
)

/*
iota是go的常量赋值自增量， 从0开始增
下面代码将a设置为0，b设置为1
*/
const (
	a = iota
	b
)

const (
	x1 = iota
	_   // 不影响iota的递增
	y
	z = "zz"
	/*
		字符串会被赋值到下一个， 但是iota会自自增，只是隐式的递增，
		在下一个出现的iota，这个递增的值就显露出来, 即在这里的P就显露出来
	*/
	k
	p = iota
)

/*
itoa1输出0 2 zz zz 5
*/
func itoa1() {
	fmt.Println(x1, y, z, k, p)
}

/*
nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。
error 类型，它是一种内置接口类型，

type error interface {
	   Error() string
}
所以nil也可以赋值给error变量， 如下两句都是合法的
var x interface{} = nil
var x error = nil
var x map[stirng]int = nil
var x []int = nil
vav x chan = nil
但字符串变量不能nil
即
vav x string = nil
是编译错误的，网上简单查了 没有为什么，这个不需要知道为什么
*/

func add(num1 int, num2 int) (result int) {
	// 变量声明
	var (
		a int
		b int
	)
	a = 1
	b = 2
	return a + b
}

func increase() {
	/*
		data := []int{1,2,3}
		i := 0
		++i //error

		fmt.Println(data[i++]) //error
	*/

	data := []int{1, 2, 3}
	i := 0
	i++ // ++ 不能放在变量前，否则编译错误
	fmt.Println(data[i])
}

func TestMultiPara(i int, m ...int) {
	fmt.Println(i)

	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
func main() {
	var i float64 = 3.13
	var j int64 = 5
	p := &i
	p = (*float64)(&j)
	fmt.Println(p)

	编译结果：
	.\main.go:11:16: cannot convert &j (type *int64) to type *float64
	错误: 进程退出代码 2.
	指针类型和uintptr的值。
	注意：指针之间不能直接转换，但是数值之间可以

}
*/

var block = "package"
func sameVar() {
	/*
		函数里面可重复定义
	 */
	block := "function"
	{
		/*
			代码块里面可重复定义
		 */
		block := 1
		fmt.Printf("The block is %d.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)

	var a interface{}
	a = 1
	if a != nil {
		/*
			代码快里面胡可以遮挡外面胡同名变量
		 */
		a := "2"
		fmt.Println(a)
	}
	fmt.Println(a)

	/*
	运行结果:
	The block is 1.
	The block is function.
	2
	1
	 */
}

func VarConst()  {
	add(3, 5)

	itoa1()

	increase()

	m := []int{7, 8, 9}
	TestMultiPara(5, m...)

	/*
		整形转为字符串的方法
	 */
	i := 65
	fmt.Println(string(i))

	sameVar()
}

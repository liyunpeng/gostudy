package basic

import (
	"fmt"
)

/*
x := 100
这是变量声明的简短模式，但这种声明方式有三个限制：
1. 必须使用显示初始化；
2. 不能提供数据类型，编译器会自动推导；
3. 只能在函数内部使用简短模式； 不能再全局使用这种剪短方式
*/
var x int = 100

/*
全局变量可以只声明，不使用，不想局部变量，不使用就包编译错误
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

func AddA(num1 int, num2 int) (result int) {
	// 变量声明
	var (
		a int
		b int
	)
	a = 1
	b = 2
	return a + b
}

func BitOperation() {
	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%08b [A]\n", a)
	fmt.Printf("%08b [B]\n", b)
	fmt.Printf("%08b (NOT B)\n", ^b)                                 // ^ 即取反
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n", b, 0xff, b^0xff) // 异或运算， 两个不一样才为1
	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n", a, b, a^b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n", a, b, a&b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n", a, b, a&^b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n", a, b, a&(^b))
}

func BitPriority() {
	fmt.Printf("0x2 & 0x2 + 0x4 -> %#x\n", 0x2&0x2+0x4)
	//prints: 0x2 & 0x2 + 0x4 -> 0x6
	//Go:    (0x2 & 0x2) + 0x4    go 位运算 优先于 加减运算
	//C++:    0x2 & (0x2 + 0x4) -> 0x2
	fmt.Printf("0x2 + 0x2 << 0x1 -> %#x\n", 0x2+0x2<<0x1)
	//prints: 0x2 + 0x2 << 0x1 -> 0x6
	//Go:     0x2 + (0x2 << 0x1)  go 位运算 优先于 加减运算
	//C++:   (0x2 + 0x2) << 0x1 -> 0x8
	fmt.Printf("0xf | 0x2 ^ 0x2 -> %#x\n", 0xf|0x2^0x2)
	//prints: 0xf | 0x2 ^ 0x2 -> 0xd
	//Go:    (0xf | 0x2) ^ 0x2   go 从左往右运算
	//C++:    0xf | (0x2 ^ 0x2) -> 0xf
}
func Increase1() {
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

func Switch11() {
	isSpace := func(ch byte) bool {
		switch (ch) {
		case ' ': //error
		case '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t')) //prints true (ok)
	fmt.Println(isSpace(' '))  //prints false (not ok)  原因是不能顺序执行

	isSpace1 := func(ch byte) bool {
		switch (ch) {
		case ' ', '\t': // 要想顺序执行， 就得放在一起
			return true
		}
		return false
	}
	fmt.Println(isSpace1('\t')) //prints true (ok)
	fmt.Println(isSpace1(' '))  //prints true (ok)
}

func TestMultiPara(i int, m ...int) {
	fmt.Println(i)

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func assert1() {
	/*
	  var data interface{} = "great"
	    if data, ok := data.(int); ok {  // data类型转换失败，被赋值为0
	        fmt.Println("[is an int] value =>",data)
	    } else {
	        fmt.Println("[not an int] value =>",data)
	        //prints: [not an int] value => 0 (not "great")
	    }

	*/
	var data interface{} = "great"
	if res, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", res)
	} else {
		fmt.Println("[not an int] value =>", data)
		//prints: [not an int] value => great (as expected)
	}
}

func Base() {
	fmt.Println("<--------------------- Base begin ----------------------->")

	BitOperation()

	BitPriority()

	Increase1()

	Switch11()

	m := []int{7, 8, 9}
	TestMultiPara(5, m...)

	assert1()

	itoa1()

	fmt.Println("<--------------------- Base end ----------------------->")

}

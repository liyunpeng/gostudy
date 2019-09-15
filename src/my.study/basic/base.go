package basic

import (
	"fmt"
	"my.study/global"
)
// 变量声明
func AddA(num1 int, num2 int) (result int){
	var (
		a int
		b int
	)
	a = 1
	b = 2

	return a+b
}

func BitOperation() {
	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%08b [A]\n",a)
	fmt.Printf("%08b [B]\n",b)
	fmt.Printf("%08b (NOT B)\n",^b)  // ^ 即取反
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n",b,0xff,b ^ 0xff) // 异或运算， 两个不一样才为1
	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n",a,b,a ^ b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n",a,b,a & b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n",a,b,a &^ b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n",a,b,a & (^b))
}

func BitPriority()  {
	fmt.Printf("0x2 & 0x2 + 0x4 -> %#x\n",0x2 & 0x2 + 0x4)
	//prints: 0x2 & 0x2 + 0x4 -> 0x6
	//Go:    (0x2 & 0x2) + 0x4    go 位运算 优先于 加减运算
	//C++:    0x2 & (0x2 + 0x4) -> 0x2
	fmt.Printf("0x2 + 0x2 << 0x1 -> %#x\n",0x2 + 0x2 << 0x1)
	//prints: 0x2 + 0x2 << 0x1 -> 0x6
	//Go:     0x2 + (0x2 << 0x1)  go 位运算 优先于 加减运算
	//C++:   (0x2 + 0x2) << 0x1 -> 0x8
	fmt.Printf("0xf | 0x2 ^ 0x2 -> %#x\n",0xf | 0x2 ^ 0x2)
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

	data := []int{1,2,3}
	i := 0
	i++   // ++ 不能放在变量前，否则编译错误
	fmt.Println(data[i])
}

func Switch11()  {
	isSpace := func(ch byte) bool {
		switch(ch) {
		case ' ': //error
		case '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t')) //prints true (ok)
	fmt.Println(isSpace(' '))  //prints false (not ok)  原因是不能顺序执行

	isSpace1 := func(ch byte) bool {
		switch(ch) {
		case ' ', '\t':  // 要想顺序执行， 就得放在一起
			return true
		}
		return false
	}
	fmt.Println(isSpace1('\t')) //prints true (ok)
	fmt.Println(isSpace1(' '))  //prints true (ok)
}


func TestMultiPara(i int, m...int){
	fmt.Println(i)

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func Base(){
	global.Logger.Println("base")

	fmt.Println("<--------------------- Base begin ----------------------->")

	BitOperation()

	BitPriority()

	Increase1()

	Switch11()

	m := []int{7, 8, 9}
	TestMultiPara(5, m...)

	fmt.Println("<--------------------- Base end ----------------------->")

}
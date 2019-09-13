package basic

import (
	"fmt"
)
// 变量声明
func AddA(num1 int, num2 int) (result int){
	//fmt.Println("addA")
	//
	var a, b, c int
	a = num1 + num2
	b = num1
	c = a+b
	fmt.Println("AddA:%d", c)

	return a+b
}

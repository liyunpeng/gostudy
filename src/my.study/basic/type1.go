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
func Type1() {
	fmt.Println("<----------------------------- Type begin ---------------------------->")
	type1()
	type2()
	fmt.Println("<----------------------------- Type end ---------------------------->")
}

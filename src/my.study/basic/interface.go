package basic

import "fmt"

type interface1 interface {
	call()
}

type implement1 struct {
}

func (im implement1 ) call(){
	fmt.Println("the implementaion method of interface is called")
}

func newCall()  {
	var p interface1
	p = new(implement1)
	p.call()
}

type dataI struct {
	name string
}
func (p *dataI) print() {
	fmt.Println("name:",p.name)
}
type printer interface {
	print()
}
func interfaceStruct() {
	d1 := dataI{"one"}
	d1.print() //ok
	//var in printer = dataI{"two"} //error
	var in printer = new(dataI)
	in.print()

	/*
		interface变量的类型转换
		int1.(string) 表示转换为字符串
	 */
	var int1 interface{}
	int1 = 1
	/*
	 interface变量.() 表示取interface变量的值，不是类型转换，
	 括号里的类型名字必须是interface变量里存放的实际值得类型
	*/
	fmt.Println( int1.(int))

	//m := map[string]dataI {"x":dataI{"three"}}
	//m["x"].print() //error  value为结构体， 用索引方式是娶不到结构体的
}

func Interface()  {
	fmt.Println("<--------------------------- Interface begin ------------------->")
	newCall()
	interfaceStruct()
	fmt.Println("<--------------------------- Interface end ------------------->")
}
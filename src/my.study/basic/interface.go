package basic

import "fmt"

type interface1 interface {
	call()
}

type implement1 struct {
}

func (im implement1) call() {
	fmt.Println("the implementaion method of interface is called")
}

func newCall() {
	var p interface1
	p = new(implement1)
	p.call()
}

type dataI struct {
	name string
}

func (p *dataI) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func interfaceStruct() {
	d1 := dataI{"one"}
	d1.print() //ok
	//var in printer = dataI{"two"} //error

	/*
		new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。

		new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），
		即类型为 *T 的值。换句话说就是，
		返回一个指针，该指针指向新分配的、类型为 T 的零值。
		适用于值类型，如数组、结构体等。

		make(T,args) 返回初始化之后的 T 类型的值，
		这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make()
		只适用于 slice、map 和 channel.
	*/
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
	fmt.Println(int1.(int))

	//m := map[string]dataI {"x":dataI{"three"}}
	//m["x"].print() //error  value为结构体， 用索引方式是娶不到结构体的
}

func do(i interface{}) {
	/*
		只有接口类型的变量，才能使用type类型选择，
		其他任何类型变量都不能有这种变量名.(type)的写法
	*/
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
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
	// data.(int) 是个类型断言
	if res, ok := data.(int); ok {  // 要用res返回， 不然panic出现
		fmt.Println("[is an int] value =>", res)
	} else {
		fmt.Println("[not an int] value =>", data)
		//prints: [not an int] value => great (as expected)
	}
}

func nil2() {
	var i interface{}
	/*
		当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil。
	*/
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}


func Interface() {
	fmt.Println("<--------------------------- Interface begin ------------------->")
	newCall()
	interfaceStruct()

	do(21)
	do("multiPara")
	do(true)

	assert1()

	nil2()
	fmt.Println("<--------------------------- Interface end ------------------->")
}

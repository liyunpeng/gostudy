package basic

import "fmt"

type data struct {
	name string
}

func map2()  {
	m := map[string]*data {"x":{"one"}}
	m["x"].name = "two" //ok
	//   m["z"].name = "what"  报运行时错误： runtime error: invalid memory address or nil pointer dereference
	fmt.Println(m["x"]) //prints: &{two}
}

func map1() {
	m := map[string]data{"x": {"one"}}
	/*
		如果map的value是结构体， 想要改变结构体里的成员的值，就要借用临时变量。
		因为map不能取地址
		如果频繁改变结构里的值，可以将value从结构体改为结构体指针 ， 示例程序：
	*/
	r := m["x"]
	r.name = "two"
	m["x"] = r
	fmt.Printf("%v", m)
}

func Map()  {
	fmt.Println("<-------------------------- Map begin ---------------------->")

	map1()

	map2()
	fmt.Println("<-------------------------- Map begin ---------------------->")
}
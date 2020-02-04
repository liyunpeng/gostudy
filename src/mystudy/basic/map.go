package basic

import "fmt"

type data struct {
	name string
}

func map2() {
	m := map[string]*data{"x": {"one"}}
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

func map3() {

	m1 := map[string]string{"key1": "value1", "key2": "value2"}

	m2 := map[string]string{"key1": "value1", "key2": "value2"}

	/*
		map  不能比较
		if m1 == m2 {
			fmt.Println("m1 == m2")
		}
	*/

	fmt.Println("map  不能比较", m1, m2)
	/*
		这种对map的赋值没有问题
	*/
	s1 := struct {
		age int
		m   map[string]string
	}{age: 1, m: map[string]string{"key1": "value1"}}

	s2 := struct {
		age int
		m   map[string]string
	}{age: 1, m: map[string]string{"key1": "value1"}}

	/*
	   结构体里右map, 这个接头体也不能比较
	   	if( s1 == s2){
	   		fmt.Println("s1 == s2")
	   	}
	*/
	fmt.Println("结构体力的map也不能比较", s1, s2)
}

type person struct {
	name string
}

func map4() {
	var m map[person]int
	p := person{"mike"}
	/*
		打印一个 map 中不存在的值时，返回元素类型的零值。
		这个例子中，m 的类型是 map[person]int，
		因为 m 中不存在 p，所以打印 int 类型的零值，即 0。
	*/
	fmt.Println(m[p])
}

func mapIf()  {
	m := make(map[string]int)
	m["a"] = 1
	/*
		v,k := m["b"] 当 key 为 b 的元素不存在的时候，
		v 会返回值类型对应的零值，k 返回 false。
		以下是判断键是否存在胡标准写法
	*/
	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}
}
func Map1() {
	fmt.Println("<-------------------------- Map begin ---------------------->")

	map1()

	map2()

	map3()

	map4()

	fmt.Println("<-------------------------- Map begin ---------------------->")
}

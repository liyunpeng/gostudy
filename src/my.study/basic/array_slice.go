package basic

import "fmt"

/*
new([]int) 之后的 list 是一个 *[]int 类型的指针，
不能对指针执行 append 操作。可以使用 make() 初始化之后再用。
同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。
所以下面蛮熟不能编译通过
func list1(){
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}
*/

func array() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	/*
		从数组里提取， 就变切片了
		创建切片有三种方式：
		基于数组创建切片和直接创建切片的方式外,
		还存在第三种创建切片的方式,也是使用比较多的方式,那就是 make 函数.

	*/
	s1 := arr[2:6]
	fmt.Println(s1)

	/*
		updateSlice(s []int) 入参是slice,  如果写为数组， 就是非法的
		即 updateSlice(arr) 是编译错误的
		出入的参数必须是数组， 而且实际穿进去的时速组length和入参length必须相同
	*/

	updateArray(arr)

	/*
		fmt.Println(arr) 和
		fmt.Printf("%v", arr)
		是一样的输出结果， 都是：[0 1 2 3 4 5 6 7 8 9]
	*/

	fmt.Printf("%v", arr)

	updateArraypointer(&arr)

	fmt.Println(arr)

	updateSlice(s1)

	fmt.Println(s1, "len=", len(s1), "cap=", cap(s1))

	/*
		对切片扩容时， 只有append， 符合go的一件事只有一个方法
	*/
	s1 = append(s1, 100)

	fmt.Println(s1, "len=", len(s1), "cap=", cap(s1))
}

/*
	不带指针的数组， 是不能改变传入数组的值, 必须明确要求传入的是数组指针
*/
func updateArray(s [10]int) {
	s[0] = 888
}

/*
	只有传入的明确指定是指针，才可以改变数组里的值
*/
func updateArraypointer(s *[10]int) {
	s[0] = 888
}

/*
	调用的时候，入参只能是slice
*/
func updateSlice(s []int) {
	s[0] = 666
}

func rangeBak() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	/*	TODO 理解range 每个元素副本
		for range 循环的时候会创建每个元素的副本，
		而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，
		所以最后 map 中的所有元素的值都是变量 val 的地址，
		因为最后 val 被赋值为3，所有输出都是3.
	*/
	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

/*
输出：
[0 0 0 0 0 1 2 3]
*/

func slice1() {
	s := make([]int, 5)
	/*
		append只能对slice操作，
	*/
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

/*
输出：
[1 2 3 4]
*/
func slice2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}

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
func ArrraySlice() {
	fmt.Println("<-------------------------ArrraySlice begin -------------------> ")
	array()
	rangeBak()
	fmt.Println("<-------------------------ArrraySlice end -------------------> ")
}

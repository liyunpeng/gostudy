package basic

import "fmt"

/*
new([]int) 之后的 list 是一个 *[]int 类型的指针，
不能对指针执行 append 操作。可以使用 make() 初始化之后再用。
同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。
所以下面代码不能编译通过
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
		对切片扩容时， 只有append， 符合go的一件事只有一个方法的风格
	*/
	s1 = append(s1, 100)

	fmt.Println(s1, "len=", len(s1), "cap=", cap(s1))
}

/*
	不带指针的数组， 不能改变传入数组的值, 必须明确要求传入的是数组指针
*/
func updateArray(s [10]int) {
	s[0] = 888
}

/*
	只有传入的明确指定是数组指针，才可以改变数组里的值
*/
func updateArraypointer(s *[10]int) {
	s[0] = 888
}

/*
	这样声明的函数，调用的时候，入参只能是slice，不能是数组
	在Go语言中，函数参数是按值传递的。当使用切片(slice)作为函数参数时，
	意味着函数将获得切片的副本：指向基础数组的起始地址的指针，以及切片的长度和容量
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

func slice1() {
	s := make([]int, 5)
	/*
		append只能对slice操作，
	*/
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	/*
		输出：
		[0 0 0 0 0 1 2 3]
	*/
}

func slice2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
	/*
		输出：
		[1 2 3 4]
	*/
}

func multiPara(num ...int) {
	num[0] = 18
}

func testmultiPara() {
	i := []int{5, 6, 7}
	multiPara(i...)
	fmt.Println(i[0])

	/*
		切片可以作为多参数传入， 但数组不行， 如下程序编译报错
		j := [5]int{1, 2, 3}
		multiPara(j...)
	*/
}

func addValue(s *[]int) {
	/*
		apppend函数返回的已经不是传入的slice的指针了， 是一个新指针，
		为了能返回这个新指针， 就把入参定义为切片指针 *[]int， 即指向切片的指针
		*s = 这样的赋值，是把指针的地址改了，即是一个新指针了
	 */
	*s = append(*s, 3)
	fmt.Printf("In addValue: s is %v\n", s)
}

func addValuetest() {
	s := []int{1, 2}
	fmt.Printf("In main, before addValue: s is %v\n", s)
	addValue(&s)
	fmt.Printf("In main, after addValue: s is %v\n", s)
	/*
		结果：
		In main, before addValue: s is [1 2]
		In addValue: s is &[1 2 3]
		In main, after addValue: s is [1 2 3]
	 */
}

func ArrraySlice() {
	fmt.Println("<-------------------------ArrraySlice begin -------------------> ")
	array()
	rangeBak()
	slice1()
	slice2()
	testmultiPara()
	addValuetest()
	fmt.Println("<-------------------------ArrraySlice end -------------------> ")
}

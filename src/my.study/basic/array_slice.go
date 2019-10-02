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

/*
	可变参数函数
 */
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


func sliceExtract() {
	a := [5]int{1, 2, 3, 4, 5}
	/*
		操作符 [i,j]。基于数组（切片）可以使用操作符 [i,j] 创建新的切片，
		从索引 i，到索引 j 结束，截取已有数组（切片）的任意部分，返回新的切片，
		新切片的值包含原数组（切片）的 i 索引的值，但是不包含 j 索引的值。i、j 都是可选的，
		i 如果省略，默认是 0，j 如果省略，默认是原数组（切片）的长度。i、j 都不能超过这个长度值。
		假如底层数组的大小为 k，截取之后获得的切片的长度和容量的计算方法：长度：j-i，容量：k-i。
		截取操作符还可以有第三个参数，形如 [i,j,k]，第三个参数 k 用来限制新切片的容量，
		但不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：j-i、k-i。
		所以例子中，切片 t 为 [4]，长度和容量都是 1。
	*/
	t := a[3:4:4]
	fmt.Println(t[0])
}

/*
func arrayLenth() {
	a := [2]int{5, 6}
	b := [3]int{5, 6}

		数组的长度也是数组类型的组成部分，
		所以 a 和 b 是不同的类型，是不能比较的，所以编译错
		a == b 这个会编译报错
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
*/

/*
cap() 函数不适用 map。
 */

func makeNon() {
	s := make(map[string]int)
	/*
	删除 map 不存在的键值对时，不会报错，相当于没有任何作用；获取不存在的减值对时，
	返回值类型对应的零值，所以返回 0。
	 */
	delete(s, "h")
	fmt.Println(s["h"])
}

func ArrraySlice() {
	fmt.Println("<-------------------------ArrraySlice begin -------------------> ")
	array()
	rangeBak()
	slice1()
	slice2()
	testmultiPara()
	addValuetest()
	makeNon()
	sliceExtract()
	fmt.Println("<-------------------------ArrraySlice end -------------------> ")
}

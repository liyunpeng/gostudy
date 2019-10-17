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
/*
nil 切片和空切片。nil 切片和 nil 相等，
一般用来表示一个不存在的切片；
空切片和 nil 不相等，表示一个空的集合。
*/
func nil1() {
	//var s1 []int
	var s2 = []int{}
	if s2 == nil {
		fmt.Println("yes nil")
	}else{
		fmt.Println("no nil")
	}
}

/*
	对于make slice而言，有两个概念需要搞清楚：长度跟容量。
	容量表示底层数组的大小，长度是你可以使用的大小。
	容量的用处在哪？在与当你用 append扩展长度时，
	如果新的长度小于容量，不会更换底层数组，否则，go 会新申请一个底层数组，
	拷贝这边的值过去，把原来的数组丢掉。也就是说，
	容量的用途是：在数据拷贝和内存申请的消耗与内存占用之间提供一个权衡。
	而长度的用途，则是为了帮助你限制切片可用成员的数量，提供边界查询的。
	所以用 make 申请好空间后，需要注意不要越界【越 len 】
 */

/*
	数组或切片的截取操作。截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，
	从数组里取出胡切片和原数组共用,
	原数组的容量决定了切片的容量不会大于原数组的容量, 除非切片apped扩容了, 有了的校报的地址
	容量是总的容量, 不是剩余的容量
	容量用在, append时, 是否要申请新的连续内存

	假设截取对象的基础数组长度为 l。
	对于操作符 [i:j]: 截取后的元素, 包括i, 包括j-1, 不包括j
	如果 i 省略，i 就默认 0，
	如果 j 省略，j 就默认为底层数组的长度，
	截取得到的切片长度和容量计算方法是 j-i、l-i。

	对于操作符 [i:j:k]，
	k用来限制切片的容量，但是不能大于数组的长度 l，
	截取得到的切片长度和容量计算方法是 j-i、k-i。
 */
func sliceLenCap() {
	s := [3]int{5, 6, 7}
	a := s[:0]
	fmt.Println("a=", a,  ", len(a)=", len(a), ", cap(a)=", cap(a))
	b := s[:2]  // 从0开始索引, 取到的是第0个, 和第1个, 即5, 6
	fmt.Println("b=", b, ", len(b)=", len(b), ", cap(b)=", cap(b))
	c := s[1:2:cap(s)]
	fmt.Println("c=", c, ", len(c)=", len(c), ", cap(c)=", cap(c))
	b[0] = 10
	fmt.Println("s=", s)
/*
结果:
   a= [] , len(a)= 0 , cap(a)= 3
   b= [5 6] , len(b)= 2 , cap(b)= 3
   c= [6] , len(c)= 1 , cap(c)= 2
   s= [10 6 7]
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
	makeNon()
	sliceExtract()
	nil1()
	sliceLenCap()
	fmt.Println("<-------------------------ArrraySlice end -------------------> ")
}

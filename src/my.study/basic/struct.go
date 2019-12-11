package basic

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

type S struct {
	m string
}

/*
	new一个类型对象的标准写法, 类型编写的标准写法， 在类型定义后面， 跟着就是一个New方法，
	不要说结构体，要说类型，go里面struct代表的是类型
*/
func NewS() *S {
	return &S{"foo"}
}
func (S) SVal() {
	fmt.Println("SVal call")
}
func (*S) SPtr() {
	fmt.Println("SPtr call")
}

type Ta struct {
	S
}

func (t Ta) TVal() {
	fmt.Println("TVal call")
}
func (t *Ta) TPtr() {
	fmt.Println("TPtr call ")
}
func (t *Ta) tPtr() {
	fmt.Println("tPtr call ")
}

func printStructMethod(a interface{}) {
	typeAll := reflect.TypeOf(a)

	fmt.Println("类型=", typeAll, "方法个数=", typeAll.NumMethod())

	for i, n := 0, typeAll.NumMethod(); i < n; i++ {
		m := typeAll.Method(i)
		fmt.Println("方法名=", m.Name, "方法类型=", m.Type)
	}
}

func receiverMethodTest() {
	// 这样的声明已经为类型分配了空间， 可不是Nil， 但是通道这样声明就是nil
	var t Ta
	t.m = "m string"
	fmt.Println("自定义类型对象t=", t)
	/*
		自定义类型变量不能与nil比较，
		if t == nil {
			fmt.Println("只声明的变量，其值是nil")
		}
	*/
	fmt.Println("自定义类型对象t的方法集：")
	printStructMethod(t)
	fmt.Println("自定义类型指针&t的方法集：")
	printStructMethod(&t)
	t.TPtr()
	t.tPtr()
}

type dataStruct struct {
	S // 声明一个匿名类型， S是上面定义的类型
	num   int
	key   *string
	bool1 bool
	items map[string]bool
	in    interface{}
}

func (d *dataStruct) pointerMethod() {
	d.num = 7
	d.m = "pointerMethod from S "
}

func (d dataStruct) valueMethod() {
	d.num = 8
	//声明的不是指针， 又想改变成员的值， 就在前面加*
	*d.key = "v.key"
	d.items["valueMethod"] = true
}

func (d *dataStruct) pointerMethod1() {
	fmt.Println("pointerMethod1 ")
}

func (d dataStruct) valueMethod1() {
	fmt.Println("valueMethod1 ")
}

func struct1() {
	var d0 dataStruct
	fmt.Println("d0 长度=", unsafe.Sizeof(d0), "字节")
	d0.pointerMethod1()
	d0.valueMethod1()

	var d1 *dataStruct
	fmt.Println("d1 长度=", unsafe.Sizeof(d0), "字节")
	d1.pointerMethod1()
	//d1.valueMethod1()
	/*
		导致：
		panic: runtime error: invalid memory address or nil pointer dereference
		[signal 0xc0000005 code=0x0 addr=0x0 pc=0x7298f0]

		goroutine 1 [running]:
		my.study/basic.struct1()
		        F:/GoWorkSpace/gostudy/src/my.study/basic/struct.go:90 +0x2b0
		my.study/basic.Struct()
		        F:/GoWorkSpace/gostudy/src/my.study/basic/struct.go:165 +0x8a
		main.base()
		        F:/GoWorkSpace/gostudy/main.go:231 +0x6d
		main.main()
		        F:/GoWorkSpace/gostudy/main.go:72 +0x30a
		exit status 2
	*/

	key := "key.1"
	/*
		自定义类型变量初始化时， 要写上字段名, 下面是类型对象初始化标准姿势
	*/
	d := dataStruct{
		num:   1,
		key:   &key,
		items: make(map[string]bool),
	}
	d.m = " from s"
	fmt.Println("d=", d)

	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=1 key=key.1 items=map[]
	d.pointerMethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=7 key=key.1 items=map[]
	d.valueMethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=7 key=v.key items=map[valueMethod:true]
}

func compare() {
	type dataStruct struct {
		num     int
		fp      float32
		complex complex64
		str     string
		char    rune
		yes     bool
		events  <-chan string
		handler interface{}
		ref     *byte
		raw     [10]byte
	}

	v1 := dataStruct{}
	v2 := dataStruct{}
	fmt.Println("v1 == v2:", v1 == v2) //prints: v1 == v2: true
}

type dataStruct1 struct {
	num    int               //ok
	checks [10]func() bool   //not comparable
	doit   func() bool       //not comparable
	m      map[string]string //not comparable
	bytes  []byte            //not comparable
}

func compare1() {
	v1 := dataStruct{}
	v2 := dataStruct{}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2)) //prints: v1 == v2: true
	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2)) //prints: m1 == m2: true
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2)) //prints: s1 == s2: true
}

func compare3() {
	var b1 []byte = nil
	b2 := []byte{}

	fmt.Println("reflect.DeepEqual(b1, b2) b1 == b2:", reflect.DeepEqual(b1, b2)) //prints: b1 == b2: false

	fmt.Println("bytes.Equal(b1, b2) b1 == b2:", bytes.Equal(b1, b2)) //prints: b1 == b2: true
}

func Struct() {
	fmt.Println("<------------------ Struct begin ------------------> ")
	receiverMethodTest()
	struct1()
	compare()
	compare1()
	compare3()

	p := *(NewS())
	fmt.Println(p.m)
	fmt.Println("<--------------------- Struct end ----------------- >")
}

/*
运行结果：
<------------------ Struct begin ------------------>
t= basic.Ta t.nummethod= 0
t= basic.Ta t.nummethod= 0
---------------- printStructMethod -------------------
t= *basic.Ta t.nummethod= 0
d0 长度= 64 字节
pointerMethod1
valueMethod1
d1 长度= 64 字节
pointerMethod1


*/

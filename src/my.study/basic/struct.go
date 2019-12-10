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
func (S) sVal(){}
func (*S) sPtr(){}

type Ta struct{
	S
}
func (Ta) tVal(){
	fmt.Println("aaaaaaaaaaaaa ")
}
func (*Ta) tPtr(){}

func method1 (a interface{}){
	t := reflect.TypeOf(a)
	fmt.Println("t=", t, "t.nummethod=", t.NumMethod())

	for i, n :=0,  t.NumMethod(); i< n; i++ {
		m := t.Method(i)
		fmt.Println("方法名=", m.Name, "方法类型=", m.Type)
	}
}

func  test1()  {
	var t Ta
	t1 := reflect.TypeOf(t)
	fmt.Println("t=", t1, "t.nummethod=", t1.NumMethod())

	method1(t)

	fmt.Println("---------------- method1 -------------------")
	method1(&t)
}
/*
	new一个结构体对象的标准写法
*/
func NewS() *S {
	return &S{"foo"}
}

type dataStruct struct {
	S     // 声明一个匿名结构体， S是上面定义的类型
	num   int
	key   *string
	bool1 bool
	items map[string]bool
	in interface{}
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
	d1.valueMethod1()

	key := "key.1"
	/*
		结构体初始化时， 要写上字段名, 下面是结构体对象初始化标准姿势
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
	test1()
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
---------------- method1 -------------------
t= *basic.Ta t.nummethod= 0
d0 长度= 64 字节
pointerMethod1
valueMethod1
d1 长度= 64 字节
pointerMethod1
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
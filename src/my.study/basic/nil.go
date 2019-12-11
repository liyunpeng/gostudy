package basic

import (
	"fmt"
	"reflect"
	"unsafe"
)

func niladdress() {
	var m map[int]string
	var ptr *int
	var sl []int
	fmt.Printf("%p\n", m)   //0x0
	fmt.Printf("%p\n", ptr) //0x0
	fmt.Printf("%p\n", sl)  //0x0
}

func nilNotKey() {
	nil := 123
	fmt.Println(nil) // 123
	/*
		cannot use nil (type int) as type map[string]int in assignment
		var _ map[string]int = nil
	*/
}

/*
一般将nil值表示为异常。
nil值的大小始终与其类型与nil值相同的non-nil值大小相同。
因此, 表示不同零值的nil标识符可能具有不同的大小。
*/
func nilSize() {
	var int1 int
	fmt.Println("int变量size=", unsafe.Sizeof(int1), "字节")

	var bool1 bool
	fmt.Println("bool变量size=", unsafe.Sizeof(bool1), "字节")

	var byte1 [1]byte
	fmt.Println("字节数组size=", unsafe.Sizeof(byte1), "字节")

	var rune1 [1]rune
	fmt.Println("rune数组size=", unsafe.Sizeof(rune1), "字节")

	var byte5 [5]byte
	fmt.Println("5个字节元素的数组size=", unsafe.Sizeof(byte5), "字节")

	var arrInt [1]int
	fmt.Println("只有一个int的数组size=", unsafe.Sizeof(arrInt), "字节")

	var p *struct{} = nil
	fmt.Println("指针size=", unsafe.Sizeof(p), "字节")

	var s []int = nil
	fmt.Println("nil切片size=", unsafe.Sizeof(s), "字节")

	var m map[int]bool = nil
	fmt.Println("nil mapsize=", unsafe.Sizeof(m), "字节")

	var c chan string = nil
	fmt.Println("nil 通道size=", unsafe.Sizeof(c), "字节")

	var f func() = nil
	fmt.Println("nil 函数变量size=", unsafe.Sizeof(f), "字节")

	var i interface{} = nil
	fmt.Println("nil 接口size=", unsafe.Sizeof(i), "字节") // 8
	/*
	运行结果：
	int变量size= 8 字节
	bool变量size= 1 字节
	只有一个int的数组size= 8 字节
	指针size= 8 字节
	nil切片size= 24 字节
	nil mapsize= 8 字节
	nil 通道size= 8 字节
	nil 函数变量size= 8 字节
	nil 接口size= 16 字节
	 */

}

/*
不同类型的nil是不能比较的

func nilCompare() {
	var m map[int]string
	var ptr *int
	fmt.Printf(m == ptr) //invalid operation: m == ptr (mismatched types map[int]string and *int)
}
*/

/*
只能在一个变量类型可以隐式转换为另一个变量的类型的情况下才可进行比较。
*/
func typeCompare() {
	type IntPtr *int
	/*
		一个变量的类型是另一个变量的的基础类型可以比较。
		下面的nil在实际代码中, 可为变量名
	*/
	fmt.Println(IntPtr(nil) == (*int)(nil))        //true
	fmt.Println((interface{})(nil) == (*int)(nil)) //false
}

/*
同一类型的两个nil值也可能无法比较
因为golang中存在map、slice和函数类型是不可比较类型，
所以比较它们的nil亦是非法的。
func sliceMapFuncCannotcompare() {
	var v1 []int = nil
	var v2 []int = nil

	fmt.Println(v1 == v2)
	fmt.Println((map[string]int)(nil) == (map[string]int)(nil))
	fmt.Println((func())(nil) == (func())(nil))
	invalid operation: v1 == v2 (slice can only be compared to nil)
	invalid operation: (map[string]int)(nil) == (map[string]int)(nil) (map can only be compared to nil)
	invalid operation: (func())(nil) == (func())(nil) (func can only be compared to nil)
}
*/

/*
如果两个比较的nil值之一是一个接口值,
而另一个不是, 假设它们是可比较的, 则比较结果总是 false。

原因是在进行比较之前, 接口值将转换为接口值的类型。
转换后的接口值具有具体的动态类型, 但其他接口值没有。这就是为什么比较结果总是错误的。
*/
func interfaceComparefalse() {
	fmt.Println((interface{})(nil) == (*int)(nil)) // false
}

/////////////////////////////////////////////////////////////
type MyError struct {
	Name string
}

func (e *MyError) Error() string {
	return "a"
}

/*
在底层，interface作为两个成员实现：一个类型和一个值。
该值被称为接口的动态值， 它是一个任意的具体值，
而该接口的类型则为该值的类型。
只有在内部值和类型都未设置时(nil, nil)，一个接口的值才为 nil。
特别是，一个 nil 接口将总是拥有一个 nil 类型。
若我们在一个接口值中存储一个 *int 类型的指针，则内部类型将为 int，无论该指针的值是什么：(int, nil)。
因此，这样的接口值会是非 nil 的
*/
func interfaceTwoMember() {

	// nil只能赋值给指针、channel、func、interface、map或slice类型的变量 (非基础类型) 否则会引发 panic
	var a *MyError                          // 这里不是 interface 类型  => 而是 一个值为 nil 的指针变量 a
	fmt.Printf("%+v\n", reflect.TypeOf(a))  // *main.MyError
	fmt.Printf("%#v\n", reflect.ValueOf(a)) // a => (*main.MyError)(nil)
	fmt.Printf("%p %#v\n", &a, a)           // 0xc42000c028 (*main.MyError)(nil)
	i := reflect.ValueOf(a)
	fmt.Println(i.IsNil()) // true

	if a == nil {
		/*
			a 是个变量指针，（注意这里 a 并不是interface）
			该变量指针只是声明，但并没有指向任何地址，所以 a == nil
		*/
		fmt.Println("a == nil") //  a == nil
	} else {
		fmt.Println("a != nil")
	}

	fmt.Println("a Error():", a.Error())
	//a 为什么 a 为 nil 也能调用方法？（指针类型的值为nil 也可以调用方法！但不能进行赋值操作！如下：）
	//2. 指针类型的值为 nil ，也能调用方法，但不能进行赋值操作，否则就会引起 panic

	/*
	 a.Name = "1"           // panic: runtime error: invalid memory address or nil pointer dereference
	*/

	var b error = a
	/*
		 问:为什么 a 为 nil 给了 b 而 b != nil
		答: var b error = a,  意味着error 是 interface 类型, 所以这里的b也是一个interface,
			即应该要满足上面提到的 interface 与 nil 的关系，
			在底层，interface作为两个成员来实现，一个类型和一个值
			即 只有当它的 type 和 data 都为 nil 时才 = nil!
			(b 是有类型的 为 *main.MyError) 所以最后会有 b != nil
	*/
	fmt.Printf("%+v\n", reflect.TypeOf(b))  // type => *main.MyError
	fmt.Printf("%+v\n", reflect.ValueOf(b)) // data => a == nil
	if b == nil {
		fmt.Println("b == nil")
	} else {
		fmt.Println("b != nil")
	}
	fmt.Println("b Error():", b.Error()) // a
}

//////////////////////////////////////////////////////////////////
type interface2 interface {
	call1()
}

type interface2S struct {
}

func (*interface2S) call1() {
	fmt.Println("interface2S call1 called")
}

func interfaceTypeValue(){
	var a interface2
	var s *interface2S

	/*
		类型变量只定义，没有定义， 也可以调用这个类型变量里的方法
		类型变量的值是nil, 那他就是nil,
		接口的变量只有值和类型都是nil, 才是nil
	 */
	if s == nil {
		fmt.Println(" s is nil")
	}else{
		fmt.Println(" s is no nil")
	}
	fmt.Printf("*interface2S type=%+v\n", reflect.TypeOf(s))
	fmt.Printf("*interface2S value=%+v\n", reflect.ValueOf(s))
	/*
	 s is nil
	*interface2S type=*basic.interface2S
	*interface2S value=<nil>
	 */

	s.call1()
	/*
	   interface2S call1 called
	*/

	if a == nil {
		fmt.Println(" a is nil")
	}else{
		fmt.Println(" a is no nil")
	}
	fmt.Printf("interface2 type=%+v\n", reflect.TypeOf(a))
	fmt.Printf("interface2 value=%+v\n", reflect.ValueOf(a))
/*
    a is nil
   interface2 type=<nil>
   interface2 value=<invalid reflect.Value>
 */

	a = s
	if a == nil {
		fmt.Println(" a is nil")
	}else{
		fmt.Println(" a is no nil")
	}
	fmt.Printf("interface2 type=%+v\n", reflect.TypeOf(a))
	fmt.Printf("interface2 value=%+v\n", reflect.ValueOf(a))
	/*
	 a is no nil
	interface2 type=*basic.interface2S
	interface2 value=<nil>
	 */
}

func Nil1() {
	fmt.Println("<------------------------ Nil1 begin ----------------------------->")
	niladdress()
	nilNotKey()
	nilSize()
	typeCompare()
	interfaceComparefalse()
	interfaceTwoMember()
	interfaceTypeValue()
	interfaceNil()
	fmt.Println("<------------------------ Nil1 end ----------------------------->")
}

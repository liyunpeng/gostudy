package basic

import (
	"bytes"
	"fmt"
	"reflect"
)

type dataStruct struct {
	num   int
	key   *string
	items map[string]bool
}

/*
	结构体里方法通过func后面加结构体变量声明完成
	一般变量名命名为this
	如果需要改变结构里的成员的值， 需要声明为指针
*/
func (this *dataStruct) pmethod() {
	this.num = 7
}

func (this dataStruct) vmethod() {
	this.num = 8
	/*
		声明的不是指针， 又想改变成员的值， 就在前面加*
	*/
	*this.key = "v.key"
	this.items["vmethod"] = true
}
func struct1() {
	key := "key.1"
	d := dataStruct{1, &key, make(map[string]bool)}
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=1 key=key.1 items=map[]
	d.pmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=7 key=key.1 items=map[]
	d.vmethod()
	fmt.Printf("num=%v key=%v items=%v\n", d.num, *d.key, d.items)
	//prints num=7 key=v.key items=map[vmethod:true]
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
	struct1()

	compare()

	compare1()
	compare3()

	fmt.Println("<--------------------- Struct end ----------------- >")
}

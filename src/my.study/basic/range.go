package basic

import "fmt"

func RangeSliceMap() {

	a := []int{10, 20, 30}

	for n := range a {
		fmt.Println( n)
	}
	for k, v := range a {
		fmt.Println(k, v)
	}

	b := map[string]int{"top1": 1000, "top2": 500}
	for k, v := range b {
		fmt.Println(k, v)
	}
}

func RangeByte1() {
	data := "A\xfe\x02\xff\x04"
	for _, v := range data { //没有[]byte强制转换， 得不到正确字节输出
		fmt.Printf("%#x ", v)
	}
	//输出: 0x41 0xfffd 0x2 0xfffd 0x4
	fmt.Println()

	for _, v := range []byte(data) {  //有[]byte强制转换， 得到正确字节输出
		fmt.Printf("%#x ", v)
	}
	//输出：0x41 0xfe 0x2 0xff 0x4
}

func Range() {
	fmt.Println("<--------------------Rage begin------------------> ")
	RangeSliceMap()

	RangeByte1()
	fmt.Println("\n <--------------------Rage end------------------> ")
}

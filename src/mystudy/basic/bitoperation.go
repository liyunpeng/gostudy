package basic

import (
	"fmt"
)

func bitOperation() {
	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%08b [A]\n", a)
	fmt.Printf("%08b [B]\n", b)
	fmt.Printf("%08b (NOT B)\n", ^b)                                 // ^ 即取反
	/*
		异或运算， 两个不一样才为1
	 */
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n", b, 0xff, b^0xff)
	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n", a, b, a^b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n", a, b, a&b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n", a, b, a&^b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n", a, b, a&(^b))
}

func bitPriority() {
	fmt.Printf("0x2 & 0x2 + 0x4 -> %#x\n", 0x2&0x2+0x4)
	//prints: 0x2 & 0x2 + 0x4 -> 0x6
	//Go:    (0x2 & 0x2) + 0x4    go 位运算 优先于 加减运算
	//C++:    0x2 & (0x2 + 0x4) -> 0x2
	fmt.Printf("0x2 + 0x2 << 0x1 -> %#x\n", 0x2+0x2<<0x1)
	//prints: 0x2 + 0x2 << 0x1 -> 0x6
	//Go:     0x2 + (0x2 << 0x1)  go 位运算 优先于 加减运算
	//C++:   (0x2 + 0x2) << 0x1 -> 0x8
	fmt.Printf("0xf | 0x2 ^ 0x2 -> %#x\n", 0xf|0x2^0x2)
	//prints: 0xf | 0x2 ^ 0x2 -> 0xd
	//Go:    (0xf | 0x2) ^ 0x2   go 从左往右运算
	//C++:    0xf | (0x2 ^ 0x2) -> 0xf
}

func BitOperation() {
	bitOperation()

	bitPriority()
}

package bufio1

import (
	"bufio"
	"fmt"
	"strings"
)

func bufioReadline() {
	/*
		strings.NewReader 返回一个os.reader对象
	*/
	sr := strings.NewReader("123\n456")
	/*
		bufio.Reader不能直接使用，需要绑定到某个io.Reader上
	*/
	reader := bufio.NewReader(sr)

	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	/*
		ReadLine读取的结束标志是\n, 所以只输出123
	*/
	fmt.Println(string(line))
}

func bufioReaderSize() {
	sr := strings.NewReader("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	/*
		将 io.Reader对象 封装成一个带缓存的 bufio.Reader 对象，
		缓存大小由 size 指定（如果小于 16 则会被设置为 16）。
		这里写0，实际就是16个字节的缓冲区
	*/
	buf := bufio.NewReaderSize(sr, 0)
	fmt.Printf("buf.Size()=%d buf.Buffered()=%d \n", buf.Size(), buf.Buffered())

	/*
		从当前缓冲区取出2个字节, 但并不是读取， 读取的指针没有变化，
		因为peek操作的原因， 会引发一次IO, 所以缓冲区被填满， 所以buffer()返回16
	buf.Buffered() 表示 可以从当前缓冲区读出来的字节数
	*/
	s, _ := buf.Peek(3)
	fmt.Printf("buf.Peek(3), buf.Buffered()=%d  %q\n", buf.Buffered(), s) // 16   "abcDE"

	b := make([]byte, 5)
	buf.Read(b)
	fmt.Printf("buf.Read(5), buf.Buffered()=%d  buf.Read()=%q\n", buf.Buffered(), b)

	b = make([]byte, 3)
	buf.Read(b)
	fmt.Printf("buf.Read(3), buf.Buffered()=%d  buf.Read()=%q\n", buf.Buffered(), b)

	b = make([]byte, 2)
	buf.Read(b)
	fmt.Printf("buf.Read(2), buf.Buffered()=%d  buf.Read()=%q\n", buf.Buffered(), b)


	buf.Discard(1)
	fmt.Printf("buf.Read(1), buf.Buffered()=%d  %q\n", buf.Buffered(), s)

	//for n, err := 0, error(nil); err == nil; {
	//	n, err = buf.Read(b)
	//	fmt.Printf("%d   %q   %v\n", buf.Buffered(), b[:n], err)
	//}
	// 5   "bcDEFGHIJK"   <nil>
	// 0   "LMNOP"   <nil>
	// 6   "QRSTUVWXYZ"   <nil>
	// 0   "123456"   <nil>
	// 0   "7890"   <nil>
	// 0   ""   EOF
}

// 示例：ReadLine
func reader3() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ\n1234567890")
	buf := bufio.NewReaderSize(sr, 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFGHIJKLMNOP"   true   <nil>
	// "QRSTUVWXYZ"   false   <nil>
	// "1234567890"   false   <nil>
	// ""   false   EOF


	// 尾部有一个换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG\n"), 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFG"   false   <nil>
	// ""   false   EOF

	fmt.Println("----------")

	// 尾部没有换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG"), 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	// "ABCDEFG"   false   <nil>
	// ""   false   EOF
}

// 示例：ReadSlice
func reader4() {
	// 尾部有换行标记
	buf := bufio.NewReaderSize(strings.NewReader("ABCDEFG\n"), 0)

	for line, err := []byte{0}, error(nil); len(line) > 0 && err == nil; {
		line, err = buf.ReadSlice('\n')
		fmt.Printf("%q   %v\n", line, err)
	}
	// "ABCDEFG\n"   <nil>
	// ""   EOF

	fmt.Println("----------")

	// 尾部没有换行标记
	buf = bufio.NewReaderSize(strings.NewReader("ABCDEFG"), 0)

	for line, err := []byte{0}, error(nil); len(line) > 0 && err == nil; {
		line, err = buf.ReadSlice('\n')
		fmt.Printf("%q   %v\n", line, err)
	}
	// "ABCDEFG"   EOF
}

func BufioReader()  {
	bufioReadline()
	bufioReaderSize()
}
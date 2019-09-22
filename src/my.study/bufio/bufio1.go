package bufio1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	for {
		/*
			readline 返回的是字节切片， 即byte slice, 即byte[]
			readline 返回的不带\t, \n等字符
			所以，不用strings.TrimSpace把这些最前面，最后面的\t\n去掉， 因为本来返回的就不带这些东西
		*/
		line, _, err := buf.ReadLine()
		//line = []byte(strings.TrimSpace(line))
		/*
			将byte[] 转换为string, 即用string(byte[] s)
		*/
		handle(string(line))
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		return nil
	}
}

func NewReaderSizeLear() {
	// 用 strings.Reader 模拟一个文件IO对象
	strReader := strings.NewReader("12345678901234567890123456789012345678901234567890")

	// go 的缓冲区最小为 16 byte，我们用最小值比较容易演示
	bufReader := bufio.NewReaderSize(strReader, 16)

	// bn = 0 但 rn >= buf_size 缓冲区不启用 发生文件IO
	tmpStr := make([]byte, 16)
	n, _ := bufReader.Read(tmpStr)
	// bufReader buffered: 0, content: 1234567890123456
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

	// bn = 0 rn < buf_size 缓冲区启用
	// 缓冲区从文件读取 buf_size 字节 发生文件IO
	// 程序从缓冲区读取 rn 字节
	// 缓冲区剩余 bn = buf_size - rn 字节
	tmpStr = make([]byte, 15)
	n, _ = bufReader.Read(tmpStr)
	// bufReader buffered: 1, content: 789012345678901
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

	// bn = 1 rn > bn
	// 程序从缓冲区读取 bn 字节 缓冲区置空 不发生文件IO
	// 注意这里只能读到一个字节
	tmpStr = make([]byte, 10)
	n, _ = bufReader.Read(tmpStr)
	// bufReader buffered: 0, content: 2
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

	// bn = 0 rn < buf_size 启用缓冲读 发生文件IO
	// 缓冲区从文件读取 buf_size 字节
	// 程序从缓冲区读取 rn 字节
	// 缓冲区剩余 bn = buf_size - rn 字节
	tmpStr = make([]byte, 10)
	n, _ = bufReader.Read(tmpStr)
	// bufReader buffered: 6, content: 3456789012
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

	// bn = 6 rn <= bn
	// 则程序冲缓冲区读取 rn 字节 不发生文件IO
	tmpStr = make([]byte, 3)
	n, _ = bufReader.Read(tmpStr)
	// bufReader buffered: 3, content: 345
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])

	// bn = 3 rn <= bn
	// 则程序冲缓冲区读取 rn 字节 不发生文件IO
	tmpStr = make([]byte, 3)
	n, _ = bufReader.Read(tmpStr)
	// bufReader buffered: 0, content: 678
	fmt.Printf("bufReader buffered: %d, content: %s\n", bufReader.Buffered(), tmpStr[:n])
}

func reader1() {
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

func reader2() {
	fmt.Println("<----------------------- reader2 begin ------------------->")
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	/*
		将 io.Reader对象 封装成一个带缓存的 bufio.Reader 对象，
		缓存大小由 size 指定（如果小于 16 则会被设置为 16）。 这里写0， 实际就是16个字节的缓冲区
	*/
	buf := bufio.NewReaderSize(sr, 0)

	b := make([]byte, 10)

	/*
		buf.Buffered() 表示 可以从当前缓冲区读出来的字节数
	*/
	fmt.Println(buf.Buffered()) // 输出0

	/*
		从当前缓冲区取出2个字节, 但并不是读取， 读取的指针没有变化，
		因为peek操作的原因， 会引发一次IO, 所以缓冲区被填满， 所以buffer()返回16
	*/
	s, _ := buf.Peek(2)
	fmt.Printf("%d   %q\n", buf.Buffered(), s) // 16   "abcDE"

	buf.Discard(1)

	for n, err := 0, error(nil); err == nil; {
		n, err = buf.Read(b)
		fmt.Printf("%d   %q   %v\n", buf.Buffered(), b[:n], err)
	}
	// 5   "bcDEFGHIJK"   <nil>
	// 0   "LMNOP"   <nil>
	// 6   "QRSTUVWXYZ"   <nil>
	// 0   "123456"   <nil>
	// 0   "7890"   <nil>
	// 0   ""   EOF
	fmt.Println("<----------------------- reader2 end ------------------->")
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

	fmt.Println("----------")

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

func Reader() {
	fmt.Println("<---------------------------- bufio 1  begin ----------------------->")
	NewReaderSizeLear()
	reader1()
	reader2()
	reader3()
	reader4()

	handle1 := func(s string) {
		fmt.Println(s)
	}
	ReadFile("a.txt", handle1)
	fmt.Println("<---------------------------- bufio 1 end ----------------------->")
}

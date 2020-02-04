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
	strReader := strings.NewReader("1234567890abcdefghiklmnopqrstuvwxyz")

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


func Reader() {
	fmt.Println("<---------------------------- bufio 1  begin ----------------------->")
	NewReaderSizeLear()


	handle1 := func(s string) {
		fmt.Println(s)
	}
	ReadFile("a.txt", handle1)
	fmt.Println("<---------------------------- bufio 1 end ----------------------->")
}

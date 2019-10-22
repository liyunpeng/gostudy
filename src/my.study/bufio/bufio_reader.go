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
	/*
		buf.Buffered() 为缓冲区存放的字节个数
	 */
	fmt.Printf("buf.Size()=%d buf.Buffered()=%d \n", buf.Size(), buf.Buffered())
	/*
		运行结果:
		buf.Size()=16 buf.Buffered()=0
	 */

	/*
		从当前缓冲区取出2个字节, 但并不是读取， 读取的指针没有变化，
		因为peek操作的原因， 会引发一次IO, 所以缓冲区被填满， 所以buffer()返回16
	*/
	s, _ := buf.Peek(3)
	fmt.Printf("buf.Peek(3), buf.Buffered()=%d  %q\n", buf.Buffered(), s)
	/*
	运行结果:
	buf.Peek(3), buf.Buffered()=16  "123"
	 */

	/*
		要读取的个数小于缓冲里的个数, 就直接从缓冲里读
	 */
	b := make([]byte, 5)
	buf.Read(b)
	fmt.Printf("buf.Read(5), buf.Buffered()=%d  buf.Read()=%q\n", buf.Buffered(), b)
	/*
	运行结果:
	buf.Read(5), buf.Buffered()=11  buf.Read()="12345"
	 */

	b = make([]byte, 3)
	buf.Read(b)
	fmt.Printf("buf.Read(3), buf.Buffered()=%d  buf.Read()=%q \n", buf.Buffered(), b)
	/*
	运行结果:
	buf.Read(3), buf.Buffered()=8  buf.Read()="678"
	*/

	b = make([]byte, 2)
	buf.Read(b)
	fmt.Printf("buf.Read(2), buf.Buffered()=%d  buf.Read()=%q\n", buf.Buffered(), b)
	/*
	运行结果:
	buf.Read(2), buf.Buffered()=6  buf.Read()="90"
	 */

	buf.Discard(1)
	fmt.Printf("buf.Discard(1), buf.Buffered()=%d \n", buf.Buffered())
	/*
	运行结果:
	buf.Discard(1), buf.Buffered()=5
	*/

	/*
		TODO
		理解\x00\x00
	 */
	b = make([]byte, 7)
	buf.Read(b)
	fmt.Printf("buf.Read(7), buf.Buffered()=%d  buf.Read()=%q \n", buf.Buffered(), b)
	/*
	运行结果:
	buf.Read(6), buf.Buffered()=0  buf.Read()="BCDEF\x00\x00"
	*/
}

func bufReadline() {
	sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ\n1234567890")

	/*
		buffer size小于16, 按16计算, 决定了readline的最大长度
		这里buffer size设置为20 则一次readline读入20个字节
	 */
	buf := bufio.NewReaderSize(sr, 20)
	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	/*
	运行结果:
	"ABCDEFGHIJKLMNOPQRST"   true   <nil>
	"UVWXYZ"   false   <nil>
	"1234567890"   false   <nil>
	""   false   EOF
	 */

	/*
		readline会自动处理尾部的\n,
	 */
	buf = bufio.NewReaderSize(strings.NewReader("readline end with mark\n"), 0)

	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	/*
		运行结果:
	   "ABCDEFG line end"   true   <nil>
	   " with mark"   false   <nil>
	   ""   false   EOF
	 */

	buf = bufio.NewReaderSize(strings.NewReader("readline end without mark"), 0)
	for line, isPrefix, err := []byte{0}, false, error(nil); len(line) > 0 && err == nil; {
		line, isPrefix, err = buf.ReadLine()
		fmt.Printf("%q   %t   %v\n", line, isPrefix, err)
	}
	/*
	   运行结果:
	   "ABCDEFG line end"   true   <nil>
	   " without mark"   false   <nil>
	   ""   false   EOF
	 */
}

func bufReadSlice() {
	buf := bufio.NewReaderSize(strings.NewReader("readslice test string over 16 words"), 0)

	for line, err := []byte{0}, error(nil); len(line) > 0 && err == nil; {
		line, err = buf.ReadSlice('\n')
		fmt.Printf("%q   %v\n", line, err)
	}
	/*
	运行结果:
	"readslice test s"   bufio: buffer full
	 */

}

func BufioReader()  {
	bufioReadline()
	bufioReaderSize()
	bufReadline()
	bufReadSlice()
}
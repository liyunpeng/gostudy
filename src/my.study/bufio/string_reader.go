package bufio1

import (
	"fmt"
	"strings"
)

func readerLenSize() {
	r := strings.NewReader("1234567890")
	fmt.Println(r.Len(), r.Size())
	/*
	运行结果:
	10 10
	 */

	var buf []byte
	buf = make([]byte, 5)
	readLen, _ := r.Read(buf)
	fmt.Println("读取到的长度:", readLen)
	/*
		运行结果:
		读取到的长度: 5
	 */

	fmt.Println(buf)
	/*
	运行结果:
		[49 50 51 52 53]
	 */

	fmt.Println(string(buf))
	/*
	运行结果:
	12345
	 */

	fmt.Println(r.Len(), r.Size())
	/*
	运行结果:
	5 10
	 */
}

func StringReader(){
	readerLenSize()
}
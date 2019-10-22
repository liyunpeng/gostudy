package bufio1

import (
	"fmt"
	"strings"
)

func reader1() {
	r := strings.NewReader("abcdefghijklmn")
	fmt.Println(r.Len()) // 输出14  初始时，未读长度等于字符串长度
	var buf []byte
	buf = make([]byte, 5)
	readLen, err := r.Read(buf)
	fmt.Println("读取到的长度:", readLen) //读取到的长度5
	if err != nil {
		fmt.Println("错误:", err)
	}
	fmt.Println(buf)      //adcde
	fmt.Println(r.Len())  //9   读取到了5个 剩余未读是14-5
	fmt.Println(r.Size()) //14   字符串的长度
}

func ReaderMain(){
	reader1()
}
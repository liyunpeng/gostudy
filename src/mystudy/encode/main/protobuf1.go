package encodemain

import (
	"fmt"
	protobuf "github.com/golang/protobuf/proto"
	"io"
	"gostudy/src/mystudy/goprotobuf"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func createProfofile() {
	//初始化protobuf数据格式
	msg := &goprotobuf.HelloWorld{
		Id:   protobuf.Int32(17),
		Name: protobuf.String("BGbiao"),
		Opt:  protobuf.Int32(18),
	}

	filename := "./protobuf-test.txt"
	fmt.Printf("使用protobuf创建文件 %s\n", filename)
	fObj, _ := os.Create(filename)
	defer fObj.Close()
	buffer, _ := protobuf.Marshal(msg)
	fObj.Write(buffer)
}

func ReadPro() {
	filename := "protobuf-test.txt"
	file, fileErr := os.Open(filename)
	checkError(fileErr)

	defer file.Close()
	fs, fsErr := file.Stat()
	checkError(fsErr)
	buffer := make([]byte, fs.Size())
	//把file文件内容读取到buffer
	_, readErr := io.ReadFull(file, buffer)
	checkError(readErr)

	//初始化pb结构体对象并将buffer中的文件内容读取到pb结构体中
	msg := &goprotobuf.HelloWorld{}
	pbErr := protobuf.Unmarshal(buffer, msg)
	checkError(pbErr)
	fmt.Printf("读取文件:%s \r\nname:%s\nid:%d\nopt:%d\n", filename, msg.GetName(), msg.GetId(), msg.GetOpt())
}

func Protobuf()  {
	fmt.Println("<------------------------------Protobuf begin ------------------------->")
	createProfofile()
	ReadPro()
	fmt.Println("<------------------------------ Protobuf end------------------------->")
}

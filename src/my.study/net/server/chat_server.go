package server

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func ChatSever() {
	listener, _ := net.Listen("tcp", ":7890")

	go broadcast() //  需要加深理解

	for { //  需要加深理解
		conn, _ := listener.Accept()
		go handleConn1(conn)
	}
}

type ClientInfo struct {
	message chan string
	Name    string
}

var (
	messages    = make(chan string)
	clientChans = make(map[ClientInfo]bool)
)

func broadcast() {
	for { //  所有不能一次select就退出的都用for
		select {
		case msg := <-messages:
			for cli := range clientChans {
				debugLog(msg)
				cli.message <- msg
			}
		}
	}

	fmt.Println("broadcast routine exit ")

}

func InputTimeout(c net.Conn, timeout time.Duration, input func(chan struct{}) ()) {
	// 一个回车后写done通道， 一个超时后写done通道
	done := make(chan struct{})
	// 客户端每次输入前，写设个sig通道， 这样读sig通道的地方会重新计时， 达到每次输入等待时间限定指定时间内
	sig := make(chan struct{})
	go func() {
		timer := time.NewTimer(timeout)
		for {
			// 不想select一次退出，必须在前有for循环
			select { // select 下所有的case都是在等着读的通道

			// 问：是否有可能主程序已经退出了， routine还在这等待，
			// 答：不会，因为本routine有自超时的处理， 这是避免routine在出程序退出还存在等待的好办法
			case <-sig:
				timer.Reset(timeout)

			case <-timer.C:
				done <- struct{}{}  // 通道作为信号使用的写法， 空结构体
				return  // 不是退出for, 而是退出了整个go routine函数
			}
		}
	}() // 启动routine，都是在调用函数，必须有这个圆括号

	go func() {
		input(sig)
		done <- struct{}{}
	}()

	<-done  // 写操作的routine和超时通道的routine 都会写这个done通道， 一个写了，这里就解除阻塞
}

func debugLog(s string){
	//  fmt.Println("[Debug] : ", s)  统一关闭调试打印
}

func handleConn1(c net.Conn) {

	clientinfo := ClientInfo{make(chan string), ""}  //不能用new的范式， 不然clientChans[clientinfo]  找不到

	debugLog(c.RemoteAddr().String())

	clientinfo.message = make(chan string) //  只有make 才创建管道， 用var xxx chan stirng只是声明

	go ouputToConnection(c, clientinfo.message)

	clientinfo.message <- "input your name:"

	input := bufio.NewScanner(c)
	inputC := func(sig chan struct{}) {
		for i := 0; input.Scan(); i++ {
			sig <- struct{}{}
			if i == 0 {
				clientinfo.Name = input.Text()

				messages <- c.RemoteAddr().String() + " " + clientinfo.Name + " 11111111111 enter chat room"

				clientChans[clientinfo] = true // 作为索引的结构体不能用new的方式，因为new放回的是指针地址
			} else {
				s := input.Text()
				debugLog(s)
				messages <- clientinfo.Name + " : " + s
			}
		}
	}

	InputTimeout(c, 30*time.Second, inputC)

	close(clientinfo.message)

	delete(clientChans, clientinfo)

	messages <- c.RemoteAddr().String() + " leave chat room"

	c.Close()

}

func ouputToConnection(c net.Conn, data <-chan string) {
	for v := range data {
		fmt.Fprintf(c, " %s \r\n", v) //  需要加深理解
	}
}

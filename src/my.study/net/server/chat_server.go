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
	for { //  所有不能一次退出的都用for
		select {
		case msg := <-messages:
			for cli := range clientChans {
				userlog(msg)
				cli.message <- msg
			}
		}
	}

	fmt.Println("broadcast routine exit ")

}

func InputTimeout(c net.Conn, timeout time.Duration, input func(chan struct{}) ()) {
	done := make(chan struct{})
	sig := make(chan struct{})
	go func() {
		timer := time.NewTimer(timeout)
		for {
			select {
			case <-sig:
				timer.Reset(timeout)

			case <-timer.C:
				done <- struct{}{}
				return
			}
		}
	}()

	go func() {
		input(sig)
		done <- struct{}{}
	}()

	<-done
}

func  userlog(s string){
	fmt.Println("[Debug] : ", s)
}

func handleConn1(c net.Conn) {

	clientinfo := ClientInfo{make(chan string), ""}

	fmt.Println(c.RemoteAddr().String())

	clientinfo.message = make(chan string) //  只有make 才创建管道， 用var xxx chan stirng只是声明

	go ouputToConnection(c, clientinfo.message)

	clientinfo.message <- "input your name:"

	input := bufio.NewScanner(c)

	inputC := func(sig chan struct{}) {
		for i := 0; input.Scan(); i++ {
			sig <- struct{}{}
			if i == 0 {
				clientinfo.Name = input.Text()

				messages <- c.RemoteAddr().String() + " " + clientinfo.Name + " enter chat room"

				clientChans[clientinfo] = true // 作为索引的结构体不能用new的方式，因为new放回的是指针地址
			} else {
				s := input.Text()
				userlog(s)
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
	//	scanner1 := bufio.NewScanner(c)
	for v := range data {

		fmt.Fprintf(c, " %s \r\n", v) //  需要加深理解
	}
}

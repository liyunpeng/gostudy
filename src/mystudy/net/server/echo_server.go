package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func Echoserver() {
	/*
	 listener, err := net.Listen("tcp", "localhost:7890") 表示仅能有本机访问
	*/
	listener, err := net.Listen("tcp", "0.0.0.0:7890")

	if err != nil {
		log.SetFlags(log.Ldate|log.Ltime|log.LstdFlags)
		log.Println("err")
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例如，连接终止
			continue
		}
		go  handleConn(conn) // 一次处理一个连接
	}
}

func handleConn(c net.Conn) {
	fmt.Printf("handleConn net.Conn=%v \n ", c)
	input := bufio.NewScanner(c)

	io.WriteString(c, time.Now().Format("2006/01/02 15:04:05\r\n"))

	for input.Scan() {
		fmt.Fprintln(c, "this is words from server: ", strings.ToUpper(input.Text()))
	}
	c.Close()
}
package client

import (
	"io"
	"log"
	"net"
	"os"
	"fmt"
	"sync"
)

func ClientOnlySendReceive() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":7890")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	// receive data from network , then send data to standar output
	go recvNetwork(conn , &wg)
	// get data from stardar input, then send data to network
	sendtoNework(conn)

	fmt.Print("chan1 \n")

	conn.CloseWrite()
	wg.Wait()

	fmt.Print("ok \n")
}

func recvNetwork(c *net.TCPConn, wg *sync.WaitGroup){
	if _, err := io.Copy(os.Stdout, c); err != nil {
		fmt.Print("recvNetwork errr")
		log.Fatal(err)
	}
	fmt.Print("receive data from network connection finished \n")
	c.CloseRead()
	wg.Done()
}

func sendtoNework(c net.Conn){
	if _, err := io.Copy(c, os.Stdin); err != nil {
		fmt.Print("sendtoNework errr")
		log.Fatal(err)
	}
	fmt.Print("send data to network connection finished \n")
}
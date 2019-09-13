package main

import (
	"fmt"
	"os"
	"sync"

	"my.study/auto"
	"my.study/basic"
	"my.study/btcoin"
	"my.study/net/client"
	"my.study/net/server"
)

var help = func() {
	fmt.Println("help")
}

func main() {
	args := os.Args

	if len(args) < 1 || args == nil {
		help()
		return
	}

	switch args[1] {
	case "gen":
		auto.Genfile()
		break
	case "base":
		base()
		break
	case "chan":
		basic.Chan()
		break
	case "net1":
		client.Net1()
		break
	case "net2":
		client.Net2()
		break
	case "net3":
		client.Net3()
		break
	case "net4":
		client.Net4()
		break
	case "bit":
		btcoin.Transaction()
	case "sha256":
		btcoin.TestSha256()
		break;
	case "client":
		client.ClientOnlySendReceive()
		break;
	case "echoserver":
		server.Echoserver()
		break
	case "chatserver":
		server.ChatSever()
		break;
	default:
		fmt.Println("cmd err,exit")
		break
	}
}
func base() {
	basic.Interface()
	basic.AddA(5, 8)
	basic.Testrage()
	basic.Testfile()
	basic.Testreflect()

	m := []int{7, 8, 9}
	basic.TestMultiPara(5, m...)

	wg := sync.WaitGroup{}

	const synNum int = 5
	wg.Add(synNum)
	for i := 0; i < synNum; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}

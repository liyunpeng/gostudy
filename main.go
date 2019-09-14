package main

import (
	"fmt"
	"my.study/auto"
	"my.study/basic"
	"my.study/btcoin"
	"my.study/net/client"
	"my.study/net/server"
	"os"
)

var help = func() {
	fmt.Println("help")
}
// TODO ： 封装一个像log文件写Log的函数 log要带时间戳

func main() {
	args := os.Args

	// TODO: 创建保存log的文件
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
	basic.Base()

	basic.Interface()

	basic.Range()


	basic.Io()

	basic.Reflect()

	basic.Encode()

	basic.Chan()
}

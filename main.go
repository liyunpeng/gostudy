package main

import (
	"fmt"
	"log"
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

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LstdFlags)
}

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
		break
	case "clientTcp":
		client.ClientOnlySendReceiveTCPConn()
		break
	case "clientNonTcp":
		if len(args) <= 2 {
			client.ClientOnlySendReceiveNonTCPConn("locolhost")
		} else {
			client.ClientOnlySendReceiveNonTCPConn(args[2])
		}
		break
	case "echoserver":
		server.Echoserver()
		break
	case "chatserver":
		server.ChatSever()
		break
	default:
		fmt.Println("cmd err,exit")
		break
	}

	//global.LoggerFile.Close()
}
func base() {
	basic.Base()
	basic.Chan()
	basic.Encode()
	basic.Go()
	basic.Interface()
	basic.Io()
	basic.Map()
	basic.Range()
	basic.Reflect()
	basic.Runtime()
	basic.Select()
	basic.Slice()
	basic.String()
	basic.Struct()
	basic.Sync()
}

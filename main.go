package main

import (
	"fmt"
	"log"
	"my.study/auto"
	"my.study/basic"
	beego1 "my.study/beego"
	"my.study/btcoin"
	context1 "my.study/context"
	"my.study/distribute"
	"my.study/encode"
	gin1 "my.study/gin"
	httpserver "my.study/http"
	"my.study/io"
	locale1 "my.study/locale"
	log1 "my.study/log"
	"my.study/net/client"
	"my.study/net/server"
	"my.study/nosql"
	regexp1 "my.study/regexp"
	rpc1 "my.study/rpc"
	"my.study/sql"
	template1 "my.study/template"
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
	case "httpserver":
		httpserver.HttpServer()
		break
	case "sql":
		sql.Sql()
		break
	case "redis":
		nosql.Redis()
		break
	case "template":
		template1.Template()
		break

	case "rpcserver":
		rpc1.RpcServer()
		break
	case "rpcclient":
		rpc1.Rpcclient()
		break

	case "regexp":
		regexp1.Exp1()
		break

	case "gin":
		gin1.Gin1()
		break
	case "context":
		context1.Context()
		break

	case "encode":
		encode1.Encode()
		break

	case "locale":
		locale1.Locale()
		break
	case "log":
		log1.Log()
		break
	case "beego":
		beego1.Beego1()
		break
	case "distribute":
		distribute.Distribute()
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
	encode1.Encode()
	basic.Go()
	basic.Interface()
	io.Io()
	basic.Map()
	basic.Range()
	basic.Reflect()
	basic.Runtime()
	basic.Select()
	basic.Slice()
	basic.String()
	basic.Struct()
	basic.Sync()
	basic.Recover()
}

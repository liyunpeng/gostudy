package main

import (
	"fmt"
	"github.com/cihub/seelog"
	"log"
	"github.com/pkg/profile"

	"gostudy/src/mystudy/auto"
	"gostudy/src/mystudy/basic"
	beego1 "gostudy/src/mystudy/beego"
	//"gostudy/src/mystudy/btcoin"
	bufio1 "gostudy/src/mystudy/bufio"
	context1 "gostudy/src/mystudy/context"
	"gostudy/src/mystudy/distribute"
	//elasticesearch1 "gostudy/src/mystudy/elasticesearch"
	main2 "gostudy/src/mystudy/encode/main"
	gin1 "gostudy/src/mystudy/gin"
	httpserver "gostudy/src/mystudy/http"
	"gostudy/src/mystudy/io"
	//iris1 "gostudy/src/mystudy/iris"
	//"gostudy/src/mystudy/iris/casbin1"
	//iriscookie "gostudy/src/mystudy/iris/cookie"
	kafka1 "gostudy/src/mystudy/kafka"
	locale1 "gostudy/src/mystudy/locale"
	"gostudy/src/mystudy/net/client"
	"gostudy/src/mystudy/net/server"
	"gostudy/src/mystudy/nosql"
	regexp1 "gostudy/src/mystudy/regexp"
	rpc1 "gostudy/src/mystudy/rpc"
	"gostudy/src/mystudy/sql"
	template1 "gostudy/src/mystudy/template"
	//xorm1 "gostudy/src/mystudy/xorm"
	"os"
)

var help = func() {
	fmt.Println("help")
}

/*
一个包可以出线多个 init() 函数,一个源文件也可以包含多个 init() 函数；
init() 函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误;
一个包被引用多次，如 A import B,C import B,A import C，B 被引用多次，但 B 包只会初始化一次；
*/
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LstdFlags)
}

func main() {
	//defer profile.Start().Stop()
 	defer profile.Start(profile.MemProfile).Stop()
	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	/*
		决定了所有seelog 在main结束时输出，控制台的随州输出seelog
	*/
	defer seelog.Flush()
	if err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	seelog.ReplaceLogger(logger)

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
	//case "bit":
	//	btcoin.Transaction()
	//case "sha256":
	//	btcoin.TestSha256()
	//	break
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
		regexp1.RegexpMain()
		break
	case "gin":
		gin1.Gin1()
		break
	case "context":
		context1.Context()
		break
	case "encode":
		main2.Encode()
		break
	case "locale":
		locale1.Locale()
		break
	case "beego":
		beego1.Beego1()
		break
	case "distribute":
		distribute.Distribute()
		break
	case "bufio":
		bufio1.BufioMain()
		break
	//case "es":
	//	elasticesearch1.Es()
	//	break
	//case "es2":
	//	elasticesearch1.Es2()
	//	break
	case "proto":
		main2.Protobuf()
		break
	//case "xorm":
	//	xorm1.Xorm1()
	//	break
	case "kafkas":
		kafka1.Kafkaserver()
		break
	case "kafkac":
		kafka1.KafkaClient()
		break

	//case "gopool":
	//	gopoll.GopollMain()
	//	break
	//case "iris":
	//	iris1.IrisMain()
	//	break
	//case "iris1":
	//	iris1.Irismain1()
	//	break
	//case "irisc":
	//	iris1.CacheClient()
	//	break
	//case "irisa":
	//	iris1.Authmain()
	//	break
	//
	//case "iriss":
	//	iris1.CacheServer()
	//	break
	//case "iriscon":
	//	iris1.Configmain()
	//	break
	//case "template1":
	//	template1.MainSever()
	//	break
	//case "iriscookie":
	//	iriscookie.Cookiemain1()
	//	break
	//case "iriscookieen":
	//	iriscookie.CookieEncryptmain()
	//	break
	////case "casbin":
	////	casbin1.Casbinmain()
	////	break
	//case "aws":
	//	iris1.Awsmain()
	//	break
	//case "xors":
	//	iris1.Xors()
	//	break
	default:
		fmt.Println("cmd err,exit")
		break
	}
}

func base() {
	basic.BitOperation()
	basic.Chan()
	main2.Encode()
	basic.Go()
	basic.Interface()
	io.Io()
	basic.Map1()
	basic.Range()
	basic.Reflect()
	basic.Runtime()
	basic.Select()
	basic.Slice()
	basic.String()
	basic.Struct()
	basic.Sync()
	basic.Recover()
	basic.Set1()

	basic.ArrraySlice()
	basic.Defer()
	basic.Type1()

	basic.IfSwitch()
	basic.BitOperation()
	basic.VarConst()
	basic.Closer1()
	basic.StubMain()

	basic.SyncPool()

	basic.Receiver()
	basic.Err()
	basic.Nil1()
}

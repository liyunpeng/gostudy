package distribute

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//用于json和结构体对象的互转
type NodeInfo struct {
	NodeId     int    `json:"nodeId"`     //节点ID，通过随机数生成
	NodeIpAddr string `json:"nodeIpAddr"` //节点ip地址
	Port       string `json:"port"`       //节点端口号
}

//添加一个节点到集群的一个请求或者响应的标准格式
type AddToClusterMessage struct {
	Source  NodeInfo `json:"source"`
	Dest    NodeInfo `json:"dest"`
	Message string   `json:"message"`
}

//将节点信息格式化输出
func (node *NodeInfo) String() string {
	return "NodeInfo {nodeId:" + strconv.Itoa(node.NodeId) + ", " +
		"nodeIpAddr:" + node.NodeIpAddr + ", port:" + node.Port + "}"
}

//将添加节点信息格式化
func (req AddToClusterMessage) String() string {
	return "AddToClusterMessage:{\n  source:" + req.Source.String() +
		",\n  dest: " + req.Dest.String() + ",\n  message:" + req.Message + " }"
}

func distribute() {
	/*
		flag.数据类型("参数名“， 默认值) 表示对命令行参数的收集
		flag收集的命令参数都是以两个中线开头， 然后更参数名， 空格， 参数值
		如 --makeMasterOnError true
	*/
	makeMasterOnError := flag.Bool("makeMasterOnError", false,
		"make this node master if unable to connect to the cluster ip provided.")

	clusterip := flag.String("clusterip", "127.0.0.1:8001",
		"ip address of any node to connnect")

	myport := flag.String("myport", "8001",
		"ip address to run this node on. default is 8001.")

	/*
		flag.parse和flag.CommandLine.Parse(os.Args[2:])都表示对命令行的解析
		必须把flag的解析放在flag的收集之后
		flag.parse从第一个参数解析， 这里需要从第二个参数解析， 所以不用flag.Parse()， 而用
		flag.CommandLine.Parse(os.Args[2:])
	*/
	//flag.Parse()

	/*
		从命令行的第2个参数开始解析
		如创建主节点， 并开始监听从节点连接到主节点的请求的命令：
		./gostudy.exe distribute  --makeMasterOnError true

		创建一个节点，并把这个节点连接到代表集群的主节点：
		./gostudy.exe distribute  --myport 8004 --clusterip 127.0.0.1:8001
	*/

	flag.CommandLine.Parse(os.Args[2:])

	rand.Seed(time.Now().UTC().UnixNano()) //种子
	myid := rand.Intn(9999999)

	//获取本机ip地址
	myIp, _ := net.InterfaceAddrs()

	/*
		集群是很多服务节点组成的， 每个服务节点提供的服务都是相同的，
		为叙述方便， 服务节点简称节点
		节点分为主节点和从节点， 主节点代表着集群
		主节点做两件事情：
		1. 客户发额请求先到主节点， 主节点把这个请求发给处于空闲状态的从节点
		2. 主节点接收新节点的加入集群的请求， 从而把新节点纳入到集群里， 作为一个从节点
		流程上，
		第一步是创建主节点， 代表集群
		./gostudy.exe distribute  --makeMasterOnError true
		第二步创建新加点， 连接主节点， 加入集群
		./gostudy.exe distribute  --myport 8004 --clusterip 127.0.0.1:8001
	*/
	me := NodeInfo{
		NodeId:     myid,
		NodeIpAddr: myIp[13].String(),
		Port:       *myport}

	dest := NodeInfo{
		NodeId:     -1,
		NodeIpAddr: strings.Split(*clusterip, ":")[0],
		Port:       strings.Split(*clusterip, ":")[1]}

	fmt.Println("我的节点信息：", me.String())

	//尝试连接到集群，在已连接的情况下向集群发送请求
	ableToConnect := connectToCluster(me, dest)

	//如果dest节点不存在，则me节点为主节点启动，否则直接退出系统
	if ableToConnect || (!ableToConnect && *makeMasterOnError) {
		if *makeMasterOnError {
			fmt.Println("将启动me节点为主节点")
		}
		listenOnPort(me)
	} else {
		fmt.Println("正在退出系统，请设置me节点为主节点")
	}
}

//发送请求时格式化json包有用的工具
func getAddToClusterMessage(source NodeInfo, dest NodeInfo, message string) AddToClusterMessage {
	return AddToClusterMessage{
		Source: NodeInfo{
			NodeId:     source.NodeId,
			NodeIpAddr: source.NodeIpAddr,
			Port:       source.Port},
		Dest: NodeInfo{
			NodeId:     dest.NodeId,
			NodeIpAddr: dest.NodeIpAddr,
			Port:       dest.Port},
		Message: message,
	}
}
func connectToCluster(me NodeInfo, dest NodeInfo) bool {
	//连接到socket的相关细节信息
	connOut, err := net.DialTimeout("tcp", dest.NodeIpAddr+":"+dest.Port,
		time.Duration(10)*time.Second)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			fmt.Println("不能连接到集群", me.NodeId)
			return false
		}
	} else {
		fmt.Println("连接到集群")
		text := "Hi nody.. 请添加我到集群"

		requestMessage := getAddToClusterMessage(me, dest, text)
		json.NewEncoder(connOut).Encode(&requestMessage)

		/*
			将网络连接的数据进行编解码
			先从网络连接NewDecoder构造一个json解码对象
			然后decoder.Decode将从连接中解码出的数据读到引用里面
		 */
		decoder := json.NewDecoder(connOut)
		var responseMessage AddToClusterMessage
		decoder.Decode(&responseMessage)
		fmt.Println("得到数据响应:\n" + responseMessage.String())
		return true
	}
	return false
}

//me节点连接其它节点成功或者自身成为主节点之后开始监听别的节点在未来可能对它自身的连接
func listenOnPort(me NodeInfo) {
	//监听即将到来的信息
	ln, _ := net.Listen("tcp", fmt.Sprint(":"+me.Port))
	//接受连接
	for {
		connIn, err := ln.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println("Error received while listening.", me.NodeId)
			}
		} else {
			var requestMessage AddToClusterMessage
			json.NewDecoder(connIn).Decode(&requestMessage)
			fmt.Println("Got request:\n" + requestMessage.String())

			text := "已添加你到集群"
			responseMessage := getAddToClusterMessage(me, requestMessage.Source, text)
			json.NewEncoder(connIn).Encode(&responseMessage)
			connIn.Close()
		}
	}
}

/*
	Hi用说明：
	第一步是创建主节点， 代表集群
	./gostudy.exe distribute  --makeMasterOnError true
	第二步创建新加点， 连接主节点， 加入集群
	./gostudy.exe distribute  --myport 8004 --clusterip 127.0.0.1:8001
*/
func Distribute() {
	distribute()
}

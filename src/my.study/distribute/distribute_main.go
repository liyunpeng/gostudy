package distribute

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

////用于json和结构体对象的互转
//type NodeInfo struct {
//	NodeId     int    `json:"nodeId"`     //节点ID，通过随机数生成
//	NodeIpAddr string `json:"nodeIpAddr"` //节点ip地址
//	Port       string `json:"port"`       //节点端口号
//}
//
////添加一个节点到集群的一个请求或者响应的标准格式
//type AddToClusterMessage struct {
//	Source  NodeInfo `json:"source"`
//	Dest    NodeInfo `json:"dest"`
//	Message string   `json:"message"`
//}
//
////将节点信息格式化输出
//func (node *NodeInfo) String() string {
//	return "NodeInfo {nodeId:" + strconv.Itoa(node.NodeId) + ", nodeIpAddr:" + node.NodeIpAddr + ", port:" + node.Port + "}"
//}
//
////将添加节点信息格式化
//func (req AddToClusterMessage) String() string {
//	return "AddToClusterMessage:{\n  source:" + req.Source.String() + ",\n  dest: " + req.Dest.String() + ",\n  message:" + req.Message + " }"
//}

func DistributeMain() {

	// makeMasterOnError := flag.Bool("makeMasterOnError", false, "make this node master if unable to connect to the cluster ip provided.")
	// clusterip := flag.String("clusterip", "127.0.0.1:8001", "ip address of any node to connnect")
	// myport := flag.String("myport", "8001", "ip address to run this node on. default is 8001."

	rand.Seed(time.Now().UTC().UnixNano()) //种子
	myid := rand.Intn(9999999)
	fmt.Println(myid)

	//获取ip地址
	myIp, _ := net.InterfaceAddrs()
	fmt.Println(myIp[13])

	//创建nodeInfo结构体
	me := NodeInfo{NodeId: myid, NodeIpAddr: myIp[13].String(), Port: "8001"}
	fmt.Println(me.String())

}
package xorm1
//
//import (
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/xormplus/core"
//	"github.com/xormplus/xorm"
//	"os"
//	"time"
//)
//
//var engine *xorm.Engine
//
//func xorm2() {
//	var _ error
//	engine, _ = xorm.NewEngine("mysql",
//		"root:root@/test?charset=utf8")
//
//	f, err := os.Create("sql.log")
//	if err != nil {
//		println(err.Error())
//		return
//	}
//	engine.SetLogger(xorm.NewSimpleLogger(f))
//
//	//logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
//	//if err != nil {
//	//	log.Fatalf("Fail to create xorm system logger: %v\n", err)
//	//}
//	//
//	//logger := xorm.NewSimpleLogger(logWriter)
//	//logger.ShowSQL(true)
//	//engine.SetLogger(logger)
//
//}
//
//type User struct {
//	Id       int       `xorm:"not null pk autoincr INT(11)"`
//	Username string    `xorm:"not null VARCHAR(32)"`
//	Birthday time.Time `xorm:"DATE"`
//	Sex      string    `xorm:"CHAR(1)"`
//	Address  string    `xorm:"VARCHAR(256)"`
//}
//
//func Xorm1() {
//
//	//创建orm引擎
//	engine, err := xorm.NewEngine("mysql", "root:root@/test?charset=utf8")
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	/*
//		在增删改查之前， 先建立一个表， 在mysql终端输入：
//		CREATE TABLE `user` (
//		  `id` int(11) NOT NULL AUTO_INCREMENT,
//		  `username` varchar(32) NOT NULL COMMENT '用户名称',
//		  `birthday` date DEFAULT NULL COMMENT '生日',
//		  `sex` char(1) DEFAULT NULL COMMENT '性别',
//		  `address` varchar(256) DEFAULT NULL COMMENT '地址',
//		  PRIMARY KEY (`id`)
//		) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;
//	*/
//
//	/*
//		连接测试
//			没有ping通说明指定的数据库不存在， 如打印下面log:
//		[xorm] [info]  2019/09/25 21:53:19.116500 PING DATABASE mysql
//		Error 1049: Unknown database 'xorm'
//	*/
//	if err := engine.Ping(); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	//日志打印SQL
//	engine.ShowSQL(true)
//
//	//设置连接池的空闲数大小
//	engine.SetMaxIdleConns(5)
//
//	//设置最大打开连接数
//	engine.SetMaxOpenConns(5)
//
//	//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
//	engine.SetTableMapper(core.SnakeMapper{})
//
//	//增
//	user := new(User)
//	user.Username = "tyming"
//	affected, err := engine.Insert(user)
//	fmt.Println(affected)
//
//	//删
//	//user := new(User)
//	//user.Username="tyming"
//	//affected_delete,err := engine.Delete(user)
//	//fmt.Println(affected_delete)
//
//	//改
//	//user := new(User)
//	//user.Username="tyming"
//	//affected_update,err := engine.Id(1).Update(user)
//	//fmt.Println(affected_update)
//
//	//查
//	user = new(User)
//	//result,err := engine.Id(1).Get(user)
//	result, err := engine.Where("id=?", 1).Get(user)
//	fmt.Println(result)
//}

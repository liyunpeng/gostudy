package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Test struct {
	Id  int
	Num int
}

func createTables(db *gorm.DB) {
	db.CreateTable(&Test{})
}

func main() {
	db, err := gorm.Open("mysql", "用户名:密码@(主机地址:端口)/数据库名称?charset=utf8&parseTime=True&loc=Local")
	if err == nil {
		fmt.Println("open db sucess")
	} else {
		fmt.Println("open db error ", err)
		return
	}

	if !db.HasTable("tests") {
		createTables(db)
	}

	test := Test{Num: 123456}
	db.Create(&test);
	fmt.Println("test.id is ", test.Id)

	var tests []Test
	db.Find(&tests)
	fmt.Println(tests)

	for index, line := range tests {
		fmt.Println("index", index, " line ", line)
	}

	defer db.Close()
}
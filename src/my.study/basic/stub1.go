package basic

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/prashantv/gostub"
)

var counter = 100

func stubGlobalVariable() {
	stubs := gostub.Stub(&counter, 200)
	defer stubs.Reset()
	fmt.Println("Counter:", counter)
}

var configFile = "config.json"

func GetConfig() ([]byte, error) {
	return ioutil.ReadFile(configFile)
}

func stubGlobalVariable1() {
	stubs := gostub.Stub(&configFile, "/tmp/test.config")
	defer stubs.Reset()
	/// 返回tmp/test.config文件的内容
	data, err := GetConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

var timeNow = time.Now
var osHostname = os.Hostname

func getDate() int {
	return timeNow().Day()
}
func getHostName() (string, error) {
	return osHostname()
}

func StubTimeNowFunction() {
	stubs := gostub.Stub(&timeNow, func() time.Time {
		return time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
	})
	fmt.Println(getDate())
	defer stubs.Reset()
}

func StubHostNameFunction() {
	stubs := gostub.StubFunc(&osHostname, "LocalHost", nil)
	defer stubs.Reset()
	fmt.Println(getHostName())
}
func StubMain() {
	fmt.Println("<-----------------------StubMain begin -------------------------->")
	stubGlobalVariable()
	stubGlobalVariable1()

	StubTimeNowFunction()

	StubHostNameFunction()
	fmt.Println("<-----------------------StubMain end -------------------------->")
}

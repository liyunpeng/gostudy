package regexp1
import (
	"fmt"
	"regexp"
)

var hzRegexp = regexp.MustCompile("[\u4e00-\u9fa5，。：,.:\\t\\r\\n]")

func StrFilterNonChinese(src *string) {
	strn := ""
	for _, c := range *src {
		if hzRegexp.MatchString(string(c)) {
			strn += string(c)
		}
	}

	*src = strn
}

/*
	把汉字过滤出来
 */
func hanziFilter()  {
	str := `
通过 & 获取 a 的地址。同时，这里还定义了一个新的变量 p 
用于保存变量 a 的地址。p 的类型为 int 指针，也就是变量 p 中的内容是变量 a 的地址。
如下代码输出它们的地址：
var a = 1
var p = &a
fmt.Printf("%p\n", p)
fmt.Printf("%p\n", &p)
我这里的输出结果是，变量 a 和 p 的地址分别为 0xc000092000 和 0xc00008c010。此时的内存的分布如下：
`
	StrFilterNonChinese(&str)
	fmt.Println("hanziFilter=", str)
	/*
	运行结果:
	通过获取的地址。同时，这里还定义了一个新的变量
	用于保存变量的地址。的类型为指针，也就是变量中的内容是变量的地址。
	如下代码输出它们的地址：


	.,
	.,
	我这里的输出结果是，变量和的地址分别为和。此时的内存的分布如下：
	 */
}
func HanziHandle() {
	hanziFilter()
}



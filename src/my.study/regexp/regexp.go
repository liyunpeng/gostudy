package regexp1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func exp1() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}

	src := string(body)

	fmt.Println("src=", src)
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	fmt.Println("after src=", src)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
	fmt.Println("<---------------------  exp1 end ------------------->")
}

func exp2() {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	//查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串，不然就是返回指定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	//查找符合条件的index位置,开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	//查找符合条件的所有的index位置，n同上
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	//查找Submatch,返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的，第三个是第二个()里面的
	//下面的输出第一个元素是"am learning Go language"
	//第二个元素是" learning Go "，注意包含空格的输出
	//第三个元素是"uage"
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	//定义和上面的FindIndex一样
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	//FindAllSubmatch,查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	//FindAllSubmatchIndex,查找所有字匹配的index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)
}

func exp3() {
	fmt.Println("<------------------------------ exp3 begin------------------------------->")
	context1 := "3.14 123123 .68 haha 1.0 abc 6.66 123."

	//MustCompile解析并返回一个正则表达式。如果成功返回，该Regexp就可用于匹配文本。
	//解析失败时会产生panic
	// \d 匹配数字[0-9]，d+ 重复>=1次匹配d，越多越好（优先重复匹配d）
	exp1 := regexp.MustCompile(`\d+\.\d+`)

	//返回保管正则表达式所有不重叠的匹配结果的[]string切片。如果没有匹配到，会返回nil。
	//result1 := exp1.FindAllString(context1, -1) //[3.14 1.0 6.66]
	result1 := exp1.FindAllStringSubmatch(context1, -1) //[[3.14] [1.0] [6.66]]

	fmt.Printf("%v\n", result1)
	fmt.Printf("\n------------------------------------\n\n")

	context2 := `
        <title>标题</title>
        <div>你过来啊</div>
        <div>hello mike</div>
        <div>你大爷</div>
        <body>呵呵</body>
    `
	//(.*?)被括起来的表达式作为分组
	//匹配<div>xxx</div>模式的所有子串
	exp2 := regexp.MustCompile(`<div>.*?</div>`)
	result2 := exp2.FindAllStringSubmatch(context2, -1)

	//[[<div>你过来啊</div> 你过来啊] [<div>hello mike</div> hello mike] [<div>你大爷</div> 你大爷]]
	fmt.Printf("%v\n", result2)
	fmt.Printf("\n------------------------------------\n\n")

	context3 := `
        <title>标题</title>
        <div>你过来啊</div>
        <div>hello 
        mike
        go</div>
        <div>你大爷</div>
        <body>呵呵</body>
    `
	exp3 := regexp.MustCompile(`<div>.*?</div>`)
	result3 := exp3.FindAllStringSubmatch(context3, -1)

	//[[<div>你过来啊</div> 你过来啊] [<div>你大爷</div> 你大爷]]
	fmt.Printf("%v\n", result3)
	fmt.Printf("\n------------------------------------\n\n")

	context4 := `
        <title>标题</title>
        <div>你过来啊</div>
        <div>hello 
        mike
        go</div>
        <div>你大爷</div>
        <body>呵呵</body>
    `
	exp4 := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	result4 := exp4.FindAllStringSubmatch(context4, -1)

	/*
	   [[<div>你过来啊</div> 你过来啊] [<div>hello
	       mike
	       go</div> hello
	       mike
	       go] [<div>你大爷</div> 你大爷]]
	*/
	fmt.Printf("%v\n", result4)
	fmt.Printf("\n------------------------------------\n\n")

	for _, text := range result4 {
		fmt.Println(text[0]) //带有div
		fmt.Println(text[1]) //不带带有div
		fmt.Println("================\n")
	}
	fmt.Println("<------------------------------ exp3 end ------------------------------->")
}

func exp4() {
	str := "880218end"
	fmt.Println("<------------------------------ exp4 begin ----------------------------- >")
	//match, _ := regexp.MatchString("\\d{5}", str) //5位连续的数字
	match, _ := regexp.MatchString(`\d{5}`, str) //5位连续的数字
	fmt.Println("src=", str)
	fmt.Println(match) //输出true

	//reg := regexp.MustCompile("\\d{6}")
	reg := regexp.MustCompile(`\d{3}`)
	//返回str中第一个匹配reg的字符串
	data := reg.Find([]byte(str))
	fmt.Println(string(data)) //880218

	//go语言正则表达式判断是否为汉字
	matchChinese, _ := regexp.Match("[\u4e00-\u9fa5]", []byte("经度"))
	fmt.Println(matchChinese) //输出true
	//go语言正则表达式判断是否含有字符（大小写）
	matchChar, _ := regexp.Match("[a-zA-Z]", []byte("av132"))
	fmt.Println(matchChar) //输出false

	//go语言正则表达式判断是否含有以数字开头，不是为true
	matchDigit, _ := regexp.Match(`[^\d]`, []byte("as132"))
	fmt.Println(matchDigit) //输出true
	//go语言正则表达式判断是否含有为IP地址
	ip := "10.32.12.01"
	pattern := "[\\d]+\\.[\\d]+\\.[\\d]+\\.[\\d]+"
	matchIp, _ := regexp.MatchString(pattern, ip)
	fmt.Println(matchIp) //输出true
	//go语言正则表达式判断是否包含某些字段
	id := "id=123;dfg"
	reg = regexp.MustCompile("id=[\\d]+")
	MEId := reg.FindString(id)
	fmt.Println(MEId) //输出id=123
	// 判断是否是个单词
	pattern = "^[A-Z]+$|(^[A-Z]?[a-z]+)$"
	char := "mAfeng"
	match, _ = regexp.Match(pattern, []byte(char))
	fmt.Println(match)
}

func exp5() {
	src := `
<!DOCTYPE html>
<html>
   <head>
      <title>Bootstrap 模板</title>
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <!-- 引入 Bootstrap -->
      <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
 
      <!-- HTML5 Shiv 和 Respond.js 用于让 IE8 支持 HTML5元素和媒体查询 -->
      <!-- 注意： 如果通过 file://  引入 Respond.js 文件，则该文件无法起效果 -->
      <!--[if lt IE 9]>
         <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
         <script src="https://oss.maxcdn.com/libs/respond.js/1.3.0/respond.min.js"></script>
      <![endif]-->
   </head>
   <body>
      <h1>Hello, world!</h1>
 
      <!-- jQuery (Bootstrap 的 JavaScript 插件需要引入 jQuery) -->
      
      <script src="jquery/jquery-3.3.1.min.js"></script>
      <!-- 包括所有已编译的插件 -->
      <script src="js/bootstrap.min.js"></script>
   </body>
</html>
`
	re, _ := regexp.Compile("<script>[\\S\\s]+</script>")
	src = re.ReplaceAllString(src, "")

	fmt.Println("src=", src)

}
func Exp1() {
	exp1()
	exp2()
	exp3()
	exp4()
	exp5()

	/*
		 不接受单引号的字符串， 只接收两种形式的字符串，
		斜杠于反斜杠最后的判断记忆， 除法是斜杠，
		一个是双引号， 里面的反斜杠会被转义，  即反斜杠会被当做转义字符
		一个是反引号， 里面的反斜杠不会被转义
		正确的书写习惯， 因为正则表达式有很多反斜杆， 所以正则表达式字符串用反引号
	*/
	fmt.Println("a \n b")
	fmt.Println(`a \n b `) // 输出 a \n b
}

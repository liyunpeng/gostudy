package regexp1

import (
	"fmt"
	"regexp"
)



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

func mustCompileFindAllString() {
	fmt.Println("<------------------------------ exp3 begin------------------------------->")
	context1 := "3.14 123123 .68 haha 1.0 abc 6.66 123."

	/*
		\d匹配数字[0-9]，d+重复>=1次匹配d
	*/
	exp1 := regexp.MustCompile(`\d+\.\d+`)
	result1 := exp1.FindAllString(context1, -1)
	fmt.Printf("\\d+\\.\\d+ FindAllString=%v\n", result1)

	/*
		运行结果:
		\d+\.\d+ FindAllString=[3.14 1.0 6.66]
	*/

	result2 := exp1.FindAllStringSubmatch(context1, -1) //
	fmt.Printf("\\d+\\.\\d+ FindAllStringSubmatch=%v\n", result2)
	/*
		运行结果:
		\d+\.\d+ FindAllStringSubmatch=[[3.14] [1.0] [6.66]]
	*/

	context2 := `
        <title>标题</title>
        <div>第一个块</div>
        <div>第二个块</div>
        <div>第三个块</div>
        <body>相应正文</body>
    `
	exp2 := regexp.MustCompile(`<div>.*?</div>`)
	result3 := exp2.FindAllStringSubmatch(context2, -1)
	fmt.Printf("<div>.*?</div> FindAllStringSubmatch=%v\n", result3)
	/*
		<div>.*?</div> FindAllStringSubmatch=[[<div>第一个块</div>] [<div>第二个块</div>] [<div>第三个块</div>]]
	*/

	exp4 := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	result4 := exp4.FindAllStringSubmatch(context2, -1)
	fmt.Printf("<div>(?s:(.*?))</div> FindAllStringSubmatch=%v\n", result4)
	/*
	   <div>(?s:(.*?))</div> FindAllStringSubmatch=[[<div>第一个块</div> 第一个块] [<div>第二个块</div> 第二个块] [<div>第三个块</div> 第三个块]]
	*/

}

func matchStringMatch(){
	str := "12345678end"
	/*
		下面的表达式等价于:
		match, _ := regexp.MatchString("\\d{5}", str) //5位连续的数字
	*/
	match, _ := regexp.MatchString(`\d{5}`, str) //5位连续的数字
	fmt.Println(match) //输出true

	//判断是否为汉字
	matchChinese, _ := regexp.Match("[\u4e00-\u9fa5]", []byte("经度"))
	fmt.Println(matchChinese) //输出true

	//判断是否含有字母
	matchChar, _ := regexp.Match("[a-zA-Z]", []byte("av132"))
	fmt.Println(matchChar) //输出false

	//判断是否含有以数字开头，不是为true
	matchDigit, _ := regexp.Match(`[^\d]`, []byte("as132"))
	fmt.Println(matchDigit) //输出true

	//判断是否含有为IP地址
	ip := "172.16.2.231"
	pattern := "[\\d]+\\.[\\d]+\\.[\\d]+\\.[\\d]+"
	/*
		等价于:
		pattern := `[\d]+\.[\d]+\.[\d]+\.[\d]+`
	 */
	matchIp, _ := regexp.MatchString(pattern, ip)
	fmt.Println(matchIp) //输出true

	// 判断是否是个单词
	pattern = "^[A-Z]+$|(^[A-Z]?[a-z]+)$"
	char := "mAfeng"
	match, _ = regexp.Match(pattern, []byte(char))
	fmt.Println(match)
}

func MustCompileFind() {
	str := "12345678end"

	reg := regexp.MustCompile(`\d{3}`)
	data := reg.Find([]byte(str))
	fmt.Println(str, "MustCompile Find=", string(data))
	/*
		运行结果:

	 */

	/*
		过滤出字段值
	 */
	id := "id=123;dfg"
	reg = regexp.MustCompile("id=[\\d]+")
	MEId := reg.FindString(id)
	fmt.Println(MEId)
	/*
		运行结果:
		123
	 */
}

func ExpBasic() {
	exp2()
	matchStringMatch()
	mustCompileFindAllString()
	MustCompileFind()

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

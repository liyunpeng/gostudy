package encode1

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One int
	two string
}

func lowercaseNotEncode() {
	in := MyData{1, "two"}

	fmt.Printf("%#v\n", in) //prints main.MyData{One:1, two:"two"}
	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded))

	// 以小写字母开头的结构体将不会被（json、xml、gob等）编码, 所以这里的two不会被编码
	// 输出  {"One":1}
	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) //prints main.MyData{One:1, two:""}
}

func Encode() {
	fmt.Println("<---------------------- Encode begin ---------------------->")
	lowercaseNotEncode()
	base641()
	aes1()
	aes2()

	simple()
	salt()
	script()

	fmt.Println("<---------------------- Encode end ---------------------->")
}

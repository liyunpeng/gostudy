package encode1

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	/*
		tag就是标签，给结构体的每个字段打上一个标签，
		标签冒号前是类型，后面是标签名
		json编码序列化出来的键值对的键名就是这个标签名
	*/
	Name  string  `json:"name"`
	Price float64 `json:"price"`

	/*
		-表示不进行序列化，
		tag标签里没有反序列化的键的名字，所以反序列化字符串不会有这一项：
		{"name":"name0","price":5000,"number":"10000","is_on_sale":"true"}
		但是，如果有了备用类型：
		ProductID int64 `json:"-,string"`
		序列化时就会有这个东西：
		{"name":"name0","price":5000,"-":"199","number":"10000","is_on_sale":"true"}
	*/
	ProductID int64 `json:"-,string"`

	/*
		tag里面加上omitempy，可以在序列化的时候忽略0值或者空值
	*/
	ProductIDOmi int64 `json:"product_idomi,omitempty"`

	/*
		在反序列化的时候，
		可能结构体类型和需要的类型不一致，这个时候可以指定多个类型,
		比如下面的Number类型为int,
		如果tag标签是number，不带其他类型，表示仅支持int,
		即
		Number   int  `json:"number`
		反序列化的字符串中键值对：
		"number":"10000",
		这样写就会报error的错误：
		json: cannot unmarshal string into Go struct field Product.number of type int
		下面的写法
		Number   int  `json:"number,string"`
		就不会报这个错误
	*/
	Number   int  `json:"number,string"`
	IsOnSale bool `json:"is_on_sale,string"`
}

func jsonEncode() {
	p := &Product{}
	p.Name = "name0"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 5000.00
	p.ProductID = 199
	//p.ProductIDOmi = 20
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	/*
		ProductID 的tag为`json:"-"`， 所以在json编码时， 被忽略。
		所以结果：
		{"name":"name0","price":5000,"-":"199","number":"10000","is_on_sale":"true"}
	*/
}

func jsonUnencode() {
	/*
		需要被反序列化的字符串要用反引号括起来
	*/
	var data = `{
		"name":"name3",
		"product_idomi":200,
		"number":"10000",
		"price":"5000",
		"is_on_sale":"true"
	}`

	/*
		待被填充的结构体的常用写法，
		声明空结构体，然后加个引用符号，这样p就是指向这个结构体的指针了
		这样p传到其他函数，也可以改变结构体成员变量的值
	 */
	p := &Product{}

	/*
		接收字符串的参数，都是要把字符串转换为byte数组，这里的反序列化的Unmarshal也不例外
	 */
	err := json.Unmarshal([]byte(data), p)
	
	fmt.Println(err)
	fmt.Println(*p)
	/*
		结果
		json: cannot unmarshal string into Go struct field Product.price of type float64
		{name3 0 0 200 10000 true}
		结果分析：
		虽然有err错误， 但是有个输出， 若果不打印这个err,  很容易让人看不到这个错误
		所以一定要对错误判断， 不能忽略错误
	*/
}

func lowercaseNotEncode() {
	type MyData struct {
		One int
		two string
	}
	in := MyData{1, "two"}

	fmt.Printf("%#v\n", in) //prints main.MyData{One:1, two:"two"}
	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded))

	/*
		以小写字母开头的结构体将不会被（json、xml、gob等）编码, 所以这里的two不会被编码
		 输出  {"One":1}
	*/
	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out) //prints main.MyData{One:1, two:""}
}

func JsonEncode() {
	fmt.Println("<---------------------------JsonEncode begin----------------------------->")
	lowercaseNotEncode()

	jsonEncode()

	jsonUnencode()

	fmt.Println("<---------------------------JsonEncode end----------------------------->")
}

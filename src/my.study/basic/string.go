package basic

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func compareString() {

	var str string = "one"
	var in interface{} = "one"
	fmt.Println("str == in:", str == in, reflect.DeepEqual(str, in))
	//prints: str == in: true true
	v1 := []string{"one", "two"}
	v2 := []interface{}{"one", "two"}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2))
	//prints: v1 == v2: false (not ok)
	data := map[string]interface{}{
		"code":  200,
		"value": []string{"one", "two"},
	}
	encoded, _ := json.Marshal(data)
	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)
	fmt.Println("data == decoded:", reflect.DeepEqual(data, decoded))
	//prints: data == decoded: false (not ok)
}

func changeChar() {
	/*
		go中string是常量，只能用双引号来表示。
		a := "this is string"
		a[0] = 'c' (这个是错误的，会报错)
		如果要做改变某个字符的操作应该先将string转换为byte数组：
	*/
	a := "this is string"
	c := []byte(a)
	c[0] = 'c'
	d := string(c)
	fmt.Println(d)
}

func String() {
	compareString()

	changeChar()
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Net1() {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://127.0.0.1:8080/index", nil)

	response, _ := client.Do(request)

	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Print(string(body))
	} else {
		fmt.Print("err: %v", response.StatusCode)
	}
}

func Net2() {

	resp, err := http.Get("http://127.0.0.1:8080/index/a=a1&b=b1")
	if err != nil {
		fmt.Printf("%#v \n", err)
		return
	}

	//defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func Net3() {

	resp, _ := http.Get("http://www.baidu.com/")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

type User struct {
	Name string
	id   int
}

func Net4() {
	user := User{"name1", 12}
	marshal, err := json.Marshal(user);
	if err != nil {
		return
	}
	fmt.Printf("marchal return : %#v \n", marshal)

	req := bytes.NewBuffer([]byte(marshal))
	bodyType := "application/json;charset=utf-8"
	resp, err := http.Post("http://127.0.0.1:8080/index/", bodyType, req)
	if err != nil {
		fmt.Printf("err: %#v \n", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

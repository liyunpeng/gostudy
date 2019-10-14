package encodemain

import (
	"fmt"
	"my.study/encode/easyjson1"
	"time"
)

func JsonEasyEncode() {
	fmt.Println("<--------------------   JsonEasyEncode  begin ------------------------->")
	s := easyjson1.Student{
		Id:   11,
		Name: "qq",
		School: easyjson1.School{
			Name: "CUMT",
			Addr: "xz",
		},
		Birthday: time.Now(),
	}
	bt, err := s.MarshalJSON()

	fmt.Println(string(bt), err)
	json := `{
		"id":11,
		"s_name":"qq",
		"s_chool":
			{"name":"CUMT",
			"addr":"xz"},
		"birthday":"2017-08-04T20:58:07.9894603+08:00"}`
	ss := easyjson1.Student{}
	ss.UnmarshalJSON([]byte(json))
	fmt.Printf("%+v", ss)
	fmt.Println("<--------------------   JsonEasyEncode  end ------------------------->")
}

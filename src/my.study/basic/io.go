package basic

import (
	"fmt"
	"os"
	"log"
)

func Testrage(){
	fmt.Println("<---------------Testrage--------> ")
	a := []int{10, 20, 30}

	for n := range a {
		fmt.Println("k", n)
	}
	for k, v := range a {
		fmt.Println(k, v)
	}

	b := map[string]int{"top1":1000, "top2":500}
	for k, v := range b {
		fmt.Println(k, v)
	}
}

func TestMultiPara(i int, m...int){
	fmt.Println(i)

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func Testfile(){
	f, _ := os.Create("a.txt")
	
	f.Write([]byte("this is text from programe"))

	f.Seek(0, os.SEEK_SET)

	p := make([]byte, 5)

	if _, err := f.Read(p); err != nil {
		log.Fatal("[File]", err)
	}

	f.Close()	
}



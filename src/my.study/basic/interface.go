package basic

import "fmt"

type interface1 interface {
	call()
}

type implement1 struct {
}

func (im implement1 ) call(){
	fmt.Println("the implementaion method of interface is called")
}

func newCall()  {
	var p interface1
	p = new(implement1)
	p.call()
}

func Interface()  {
	fmt.Println("<--------------------------- Interface begin ------------------->")
	newCall()
	fmt.Println("<--------------------------- Interface begin ------------------->")
}
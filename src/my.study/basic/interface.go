package basic

import "fmt"

type interface1 interface {
	call()
}

type implement1 struct {
}

func (im implement1 ) call(){
	fmt.Println("call")
}

func Interface()  {
	var p interface1
	p = new(implement1)
	p.call()
}


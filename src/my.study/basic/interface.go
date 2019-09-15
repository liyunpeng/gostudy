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

type dataI struct {
	name string
}
func (p *dataI) print() {
	fmt.Println("name:",p.name)
}
type printer interface {
	print()
}
func interfaceStruct() {
	d1 := dataI{"one"}
	d1.print() //ok
	//var in printer = dataI{"two"} //error
	var in printer = new(dataI)
	in.print()

	//m := map[string]dataI {"x":dataI{"three"}}
	//m["x"].print() //error
}

func Interface()  {
	fmt.Println("<--------------------------- Interface begin ------------------->")
	newCall()
	interfaceStruct()
	fmt.Println("<--------------------------- Interface end ------------------->")
}
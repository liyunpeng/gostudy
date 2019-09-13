package basic

import "fmt"

// bibao
func func1() func() int {
	return func() int {
		return 10
	}
}

func bibao()  {
	//var f1 func
	f1 := func1()
	f1=func1()
	fmt.Println(f1())
}

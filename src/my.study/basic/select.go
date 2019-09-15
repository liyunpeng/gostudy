package basic

import "fmt"

func select1() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			break loop   // 因为for后面还有代码要执行， 只跳出for, 可以用break这种方式
		}
	}
	fmt.Println("out!")
}

func Select()  {
	select1()
}
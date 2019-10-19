package basic

import (
	"errors"
	"fmt"
	"log"
)

var errv = errors.New("errors.New 提示信息")

type TYPE1 struct {
	x, y int
}
/*
	定义了自己的errorr类型
 */
func (TYPE1) Error() string {
	return "TYPE1 ERROR"
}

func div(x int, y int) (int, error) {
	if y == 0 {
		/*
			对自己error类型的引用
		 */
		return 0, TYPE1{x, y}
	}

	return x / y, nil
}

func err1() {
	x, err := div(6, 0)

	if err == errv {
		log.Fatal(err)
	}
	println(x)

	if err != nil {
		switch e := err.(type) {
		case TYPE1:
			fmt.Println(e, e.x, e.y)
		default:
			fmt.Println(e)
		}
	}

	/*
		运行结果:
	TYPE1 ERROR 6 0
	 */
}

func Err()  {
	err1()
}
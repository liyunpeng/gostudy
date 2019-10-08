package basic

import (
	"fmt"
)

func go1() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}
	//time.Sleep(3 * time.Second)
	//goroutines print: three, three, three
}

func go2() {
	data := []string{"one", "two", "three"}
	for _, v := range data {
		vcopy := v //
		go func() {
			fmt.Println(vcopy)
		}()
	}
	//time.Sleep(1 * time.Second)
	//goroutines print: one, two, three
}

func defer1() {
	// defer执行的语句只有在defer声明的时候求值
	var i int = 1
	defer fmt.Println("result =>", func() int { return i * 2 }())
	i++
	//prints: result => 2 (not ok if you expected 4)
}

func go3() {
	done := false
	go func() {
		done = true
	}()
	for !done {
	}
	fmt.Println("go3 done!")
}
func Go() {
	fmt.Println("<-------------------------- Go begin -------------------->")
	defer1()
	go1()
	go2()
	go3()
	fmt.Println("<-------------------------- Go end -------------------->")
}
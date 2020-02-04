package basic

import "fmt"

type Person1 struct {
	name string
	age  int
}

/*
	函数定义中的前置的结构体变量就是receiver, 函数要改变作用调用者的结构体, receiver就用指针
*/
func (this *Person1) Growth() {
	this.age++
}

func (this *Person1) ChangeName(newname string) {
	this.name = newname
}

func receiver1() {
	p := Person1{"wangzy", 30}
	p.Growth()
	fmt.Printf("%d \n", p.age)
}

type MyStruct1 struct{ i1, i2, i3, i4 int }

func (h *MyStruct1) incI2() { h.i2++ }
func (h MyStruct1) incI3()  { h.i3++ }

func receiver2() {
	h1 := MyStruct1{0x1111, 0x2222, 0x3333, 0x4444}
	h2 := &h1

	h2.incI2()
	h2.incI3()

	fmt.Printf("i2=%x, i3=%x\n", h1.i2, h2.i3)
}

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func receiver3() {
	u := user{"Bill", "bill@email.com"}
	u.notify()
	/*
		使用指针接收者来实现一个接口，
		那么只有指向那个类型的指针才能够实现对应的接口。
		所以这里sendNotificatioin(u)是不合法的
	sendNotificatioin(&u)是合法的

			如果使用值接收者来实现一个接口，如:
		func (u *user) notify() {
			fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
		}
		sendNotificatioin(u) 就是合法的,
		sendNotificatioin(&u) 也是合法的

	方法集:
	①：func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
	}
	sendNotificatioin(&u)

	②：func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
	}
	sendNotificatioin(u)

	③：func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
	}
	sendNotificatioin(&u)
	*/
	sendNotificatioin(&u)
}

func sendNotificatioin(n notifier) {
	n.notify()
}

func Receiver() {
	fmt.Println("<-------------------------- Receiver begin ------------------------->")
	receiver1()
	receiver2()
	receiver3()
	fmt.Println("<-------------------------- Receiver end ------------------------->")
}

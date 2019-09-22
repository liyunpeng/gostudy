package basic

import (
	"fmt"
	"reflect"
)

type User struct {
	id int
	name string
}

func (u User) Memfunc(){
	fmt.Println("memfunc")
}

func (u User) MemfuncWitshargs(i int){
	fmt.Println("MemfuncWitshargs : ", i)
}

func reflect1(any interface{}){
	fmt.Printf("interface{}=%#v \n\n", any)

	type1 := reflect.TypeOf(any)
	fmt.Printf("reflect.TypeOf(any)=%v \n\n", type1)

	value := reflect.ValueOf(any)
	fmt.Printf("reflect.ValueOf(any)=%v \n\n", value)

	for i := 0; i< type1.NumField(); i++ {
		field := type1.Field(i)
		value1 := value.Field(i)
/*
		if i == 0  {
			newValue := value1.Elem()
			newValue.SetInt(100)
		}
*/
		fmt.Printf("for %d, field.Type=%v, field.Name=%v, value1=%v \n", i, field.Type, field.Name, value1)
	}

	fmt.Println("\n")

	for i := 0; i< type1.NumMethod(); i++ {
		method := type1.Method(i)
		fmt.Printf("method.Type=%v, method.Name=%v \n", method.Type, method.Name)

		//fmt.Println(type1.NumIn())

		if i == 0 {
			// 无参数反射方法的调用
			methodValue := value.MethodByName(method.Name)
			args := make([]reflect.Value, 0)
			/*
				不管反射的方法有没有参数， 用call调用， 必须有个切片作为参数
			 */
			methodValue.Call(args)
		}else if i == 1 {
			methodValue := value.MethodByName(method.Name)
			args := []reflect.Value{reflect.ValueOf(50)}
			methodValue.Call(args)
		}
	}
}


func Testreflect(){
	user := User{1, "abc"}
	fmt.Printf("user: %#v \n", user)
	reflect1(user)
	fmt.Printf("after reflect1 user: %#v \n", user)
}

func Reflect(){
	fmt.Println("<------------------------------Reflect begin ------------->")

	Testreflect()

	fmt.Println("\n<------------------------------Reflect end ------------->")

}
	




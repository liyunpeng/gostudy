package template1

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func template1() {

	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "Astaxie"}
	t.Execute(os.Stdout, p)
}

func Template()  {
	fmt.Println("<---------------------- template begin ------------------>")
	template1()
	fmt.Println("\n<---------------------- template end ------------------>")
}

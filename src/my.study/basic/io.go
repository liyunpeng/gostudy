package basic

import (
	"fmt"
	"os"
	"log"
)





func Testfile(){
	f, _ := os.Create("a.txt")
	
	f.Write([]byte("this is text from programe"))

	f.Seek(0, os.SEEK_SET)

	p := make([]byte, 5)

	if _, err := f.Read(p); err != nil {
		log.Fatal("[File]", err)
	}

	f.Close()	
}


func Io(){
	fmt.Println("<---------------------- Io begin --------------------->")

	Testfile()

	fmt.Println("<----------------------Io end ------------------------->")
}


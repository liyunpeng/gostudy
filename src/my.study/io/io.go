package io

import (
	"crypto/md5"
	"fmt"
	"io"
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
func writestring()  {
	h := md5.New()
	io.WriteString(h, "1234567890123456")
	fmt.Println("h.size=", h.Size())
	io.WriteString(h, "12345678901234561234567890123456")
	fmt.Println("h.size=", h.Size())

	io.WriteString(os.Stdout, "1234567890123456")
	io.WriteString(os.Stdout, "12345678901234561234567890123456")



}

func Io(){
	fmt.Println("<---------------------- Io begin --------------------->")

	Testfile()

	writestring()
	fmt.Println("<----------------------Io end ------------------------->")
}


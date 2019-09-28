package basic

import (
	"fmt"
	"log"
	"path/filepath"
	_ "reflect"
	"runtime"
	"strings"
)

func run1() {
	/*
		GOMAXPROCS表示了CPU的数量，Go将使用这个数量来运行goroutine。
		而runtime.GOMAXPROCS()函数的文档让人更加的迷茫
	*/
	fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 1
	fmt.Println(runtime.NumCPU())       //prints: 1 (on play.golang.org)
	runtime.GOMAXPROCS(20)
	fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 20
	runtime.GOMAXPROCS(300)
	fmt.Println(runtime.GOMAXPROCS(-1)) //prints: 256
}

type MyStruct struct {
}

func (m *MyStruct) foo(p string) {
	ENTRY("")
	ENTRY("Param p=%s", p)
	DEBUG("Test %s %s", "Hello", "World")
}

func DEBUG(formating string, args ...interface{}) {
	LOG("DEBUG", formating, args...)
}

func ENTRY(formating string, args ...interface{}) {
	LOG("ENTRY", formating, args...)
}

func LOG(level string, formating string, args ...interface{}) {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(2)
	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo

		filename = filepath.Base(filename) // /full/path/basename.go => basename.go
	}

	log.Printf("%s:%d:%s: %s: %s\n", filename, line, funcname, level, fmt.Sprintf(formating, args...))
}

func run2() {
	ss := MyStruct{}
	ss.foo("helloworld")
}

func Runtime() {
	fmt.Println("<-------------------------------- Runtime ----------------------->")
	run1()
	run2()
	fmt.Println("<-------------------------------- Runtime ----------------------->")
}